// Copyright 2018-2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { ICertificate, Certificate, CertificateValidation } from 'aws-cdk-lib/aws-certificatemanager';
import {
  CloudFrontWebDistribution,
  OriginProtocolPolicy,
  PriceClass,
  ViewerCertificate,
  ViewerProtocolPolicy,
} from 'aws-cdk-lib/aws-cloudfront';
import { ARecord, AaaaRecord, IHostedZone, RecordTarget } from 'aws-cdk-lib/aws-route53';
import { CloudFrontTarget } from 'aws-cdk-lib/aws-route53-targets';
import { BlockPublicAccess, Bucket, RedirectProtocol } from 'aws-cdk-lib/aws-s3';
import { ArnFormat, RemovalPolicy, Stack, Token } from 'aws-cdk-lib/core';
import { md5hash } from 'aws-cdk-lib/core/lib/helpers-internal';
import { ROUTE53_PATTERNS_USE_CERTIFICATE } from 'aws-cdk-lib/cx-api';
import { Construct } from 'constructs';
import { domainAddTrailingDot } from './utils';

/**
 * Properties to configure an HTTPS Redirect
 *
 * Patched version of Taken from https://github.com/aws/aws-cdk/blob/main/packages/aws-cdk-lib/aws-route53-patterns/lib/website-redirect.ts
 * - Workaround for https://github.com/aws/aws-cdk/issues/26572
 * - Remove deprecated DnsValidatedCertificate
 */
export interface HttpsRedirectProps {
  /**
   * Hosted zone of the domain which will be used to create alias record(s) from
   * domain names in the hosted zone to the target domain. The hosted zone must
   * contain entries for the domain name(s) supplied through `recordNames` that
   * will redirect to the target domain.
   *
   * Domain names in the hosted zone can include a specific domain (example.com)
   * and its subdomains (acme.example.com, zenith.example.com).
   *
   */
  readonly zone: IHostedZone;

  /**
   * The redirect target fully qualified domain name (FQDN). An alias record
   * will be created that points to your CloudFront distribution. Root domain
   * or sub-domain can be supplied.
   */
  readonly targetDomain: string;

  /**
   * The domain names that will redirect to `targetDomain`
   *
   * @default - the domain name of the hosted zone
   */
  readonly recordNames?: string[];

  /**
   * The AWS Certificate Manager (ACM) certificate that will be associated with
   * the CloudFront distribution that will be created. If provided, the certificate must be
   * stored in us-east-1 (N. Virginia)
   *
   * @default - A new certificate is created in us-east-1 (N. Virginia)
   */
  readonly certificate?: ICertificate;
}

/**
 * Allows creating a domainA -> domainB redirect using CloudFront and S3.
 * You can specify multiple domains to be redirected.
 */
export class HttpsRedirectPatched extends Construct {
  constructor(scope: Construct, id: string, props: HttpsRedirectProps) {
    super(scope, id);

    const domainNames = props.recordNames ?? [props.zone.zoneName];

    if (props.certificate) {
      const certificateRegion = Stack.of(this).splitArn(
        props.certificate.certificateArn,
        ArnFormat.SLASH_RESOURCE_NAME
      ).region;
      if (!Token.isUnresolved(certificateRegion) && certificateRegion !== 'us-east-1') {
        throw new Error(
          `The certificate must be in the us-east-1 region and the certificate you provided is in ${certificateRegion}.`
        );
      }
    }
    const redirectCert = props.certificate ?? this.createCertificate(domainNames, props.zone);

    const redirectBucket = new Bucket(this, 'RedirectBucket', {
      websiteRedirect: {
        hostName: props.targetDomain,
        protocol: RedirectProtocol.HTTPS,
      },
      removalPolicy: RemovalPolicy.DESTROY,
      blockPublicAccess: BlockPublicAccess.BLOCK_ALL,
    });
    const redirectDist = new CloudFrontWebDistribution(this, 'RedirectDistribution', {
      defaultRootObject: '',
      originConfigs: [
        {
          behaviors: [{ isDefaultBehavior: true }],
          customOriginSource: {
            domainName: redirectBucket.bucketWebsiteDomainName,
            originProtocolPolicy: OriginProtocolPolicy.HTTP_ONLY,
          },
        },
      ],
      viewerCertificate: ViewerCertificate.fromAcmCertificate(redirectCert, {
        aliases: domainNames,
      }),
      comment: `Redirect to ${props.targetDomain} from ${domainNames.join(', ')}`,
      priceClass: PriceClass.PRICE_CLASS_ALL,
      viewerProtocolPolicy: ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
    });

    domainNames.forEach((domainName) => {
      const hash = md5hash(domainName).slice(0, 6);
      domainName = domainAddTrailingDot(domainName);
      const aliasProps = {
        recordName: domainName,
        zone: props.zone,
        target: RecordTarget.fromAlias(new CloudFrontTarget(redirectDist)),
      };
      new ARecord(this, `RedirectAliasRecord${hash}`, aliasProps);
      new AaaaRecord(this, `RedirectAliasRecordSix${hash}`, aliasProps);
    });
  }

  /**
   * Gets the stack to use for creating the Certificate
   * If the current stack is not in `us-east-1` then this
   * will create a new `us-east-1` stack.
   *
   * CloudFront is a global resource which you can create (via CloudFormation) from
   * _any_ region. So I could create a CloudFront distribution in `us-east-2` if I wanted
   * to (maybe the rest of my application lives there). The problem is that some supporting resources
   * that CloudFront uses (i.e. ACM Certificates) are required to exist in `us-east-1`. This means
   * that if I want to create a CloudFront distribution in `us-east-2` I still need to create a ACM certificate in
   * `us-east-1`.
   *
   * In order to do this correctly we need to know which region the CloudFront distribution is being created in.
   * We have two options, either require the user to specify the region or make an assumption if they do not.
   * This implementation requires the user specify the region.
   */
  private certificateScope(): Construct {
    const stack = Stack.of(this);
    const parent = stack.node.scope;
    if (!parent) {
      throw new Error(`Stack ${stack.stackId} must be created in the scope of an App or Stage`);
    }
    if (Token.isUnresolved(stack.region)) {
      throw new Error(`When ${ROUTE53_PATTERNS_USE_CERTIFICATE} is enabled, a region must be defined on the Stack`);
    }
    if (stack.region !== 'us-east-1') {
      const stackId = `certificate-redirect-stack-${stack.node.addr}`;
      const certStack = parent.node.tryFindChild(stackId) as Stack;
      return (
        certStack ??
        new Stack(parent, stackId, {
          env: { region: 'us-east-1', account: stack.account },
        })
      );
    }
    return this;
  }

  /**
   * Creates a certificate.
   */
  private createCertificate(domainNames: string[], zone: IHostedZone): ICertificate {
    // this preserves backwards compatibility. Previously the certificate was always created in `this` scope
    // so we need to keep the name the same
    const id = this.certificateScope() === this ? 'RedirectCertificate' : 'RedirectCertificate' + this.node.addr;
    return new Certificate(this.certificateScope(), id, {
      domainName: domainNames[0],
      subjectAlternativeNames: domainNames,
      validation: CertificateValidation.fromDns(zone),
    });
  }
}
