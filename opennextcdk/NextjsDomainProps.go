// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

type NextjsDomainProps struct {
	// The domain to be assigned to the website URL (ie. domain.com).
	//
	// Supports domains that are hosted either on [Route 53](https://aws.amazon.com/route53/) or externally.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Specify additional names that should route to the Cloudfront Distribution.
	AlternateNames *[]*string `field:"optional" json:"alternateNames" yaml:"alternateNames"`
	// Import the certificate for the domain.
	//
	// By default, SST will create a certificate with the domain name. The certificate will be created in the `us-east-1`(N. Virginia) region as required by AWS CloudFront.
	//
	// Set this option if you have an existing certificate in the `us-east-1` region in AWS Certificate Manager you want to use.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// An alternative domain to be assigned to the website URL.
	//
	// Visitors to the alias will be redirected to the main domain. (ie. `www.domain.com`).
	//
	// Use this to create a `www.` version of your domain and redirect visitors to the root domain.
	DomainAlias *string `field:"optional" json:"domainAlias" yaml:"domainAlias"`
	// Import the underlying Route 53 hosted zone.
	HostedZone awsroute53.IHostedZone `field:"optional" json:"hostedZone" yaml:"hostedZone"`
	// Set this option if the domain is not hosted on Amazon Route 53.
	IsExternalDomain *bool `field:"optional" json:"isExternalDomain" yaml:"isExternalDomain"`
}

