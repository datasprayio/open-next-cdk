// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/datasprayio/open-next-cdk/opennextcdk/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/datasprayio/open-next-cdk/opennextcdk/internal"
)

// Create a CloudFront distribution to serve a Next.js application.
type NextjsDistribution interface {
	constructs.Construct
	// The AWS Certificate Manager certificate for the custom domain.
	Certificate() awscertificatemanager.ICertificate
	SetCertificate(val awscertificatemanager.ICertificate)
	CustomDomainName() *string
	// If the custom domain is enabled, this is the URL of the website with the custom domain.
	CustomDomainUrl() *string
	// The internally created CloudFront `Distribution` instance.
	Distribution() awscloudfront.Distribution
	SetDistribution(val awscloudfront.Distribution)
	// The domain name of the internally created CloudFront Distribution.
	DistributionDomain() *string
	// The ID of the internally created CloudFront Distribution.
	DistributionId() *string
	// The Route 53 hosted zone for the custom domain.
	HostedZone() awsroute53.IHostedZone
	SetHostedZone(val awsroute53.IHostedZone)
	// The tree node.
	Node() constructs.Node
	Props() *NextjsDistributionProps
	SetProps(val *NextjsDistributionProps)
	TempBuildDir() *string
	SetTempBuildDir(val *string)
	// The CloudFront URL of the website.
	Url() *string
	LookupHostedZone() awsroute53.IHostedZone
	// Returns a string representation of this construct.
	ToString() *string
	ValidateCustomDomainSettings()
}

// The jsii proxy struct for NextjsDistribution
type jsiiProxy_NextjsDistribution struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NextjsDistribution) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) CustomDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) CustomDomainUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) Distribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) DistributionDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) HostedZone() awsroute53.IHostedZone {
	var returns awsroute53.IHostedZone
	_jsii_.Get(
		j,
		"hostedZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) Props() *NextjsDistributionProps {
	var returns *NextjsDistributionProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) TempBuildDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBuildDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsDistribution) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewNextjsDistribution(scope constructs.Construct, id *string, props *NextjsDistributionProps) NextjsDistribution {
	_init_.Initialize()

	if err := validateNewNextjsDistributionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NextjsDistribution{}

	_jsii_.Create(
		"open-next-cdk.NextjsDistribution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNextjsDistribution_Override(n NextjsDistribution, scope constructs.Construct, id *string, props *NextjsDistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"open-next-cdk.NextjsDistribution",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NextjsDistribution)SetCertificate(val awscertificatemanager.ICertificate) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_NextjsDistribution)SetDistribution(val awscloudfront.Distribution) {
	if err := j.validateSetDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distribution",
		val,
	)
}

func (j *jsiiProxy_NextjsDistribution)SetHostedZone(val awsroute53.IHostedZone) {
	_jsii_.Set(
		j,
		"hostedZone",
		val,
	)
}

func (j *jsiiProxy_NextjsDistribution)SetProps(val *NextjsDistributionProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_NextjsDistribution)SetTempBuildDir(val *string) {
	if err := j.validateSetTempBuildDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempBuildDir",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NextjsDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextjsDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.NextjsDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NextjsDistribution_FallbackOriginRequestPolicyProps() *awscloudfront.OriginRequestPolicyProps {
	_init_.Initialize()
	var returns *awscloudfront.OriginRequestPolicyProps
	_jsii_.StaticGet(
		"open-next-cdk.NextjsDistribution",
		"fallbackOriginRequestPolicyProps",
		&returns,
	)
	return returns
}

func NextjsDistribution_SetFallbackOriginRequestPolicyProps(val *awscloudfront.OriginRequestPolicyProps) {
	_init_.Initialize()
	if err := validateNextjsDistribution_SetFallbackOriginRequestPolicyPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"open-next-cdk.NextjsDistribution",
		"fallbackOriginRequestPolicyProps",
		val,
	)
}

func NextjsDistribution_ImageCachePolicyProps() *awscloudfront.CachePolicyProps {
	_init_.Initialize()
	var returns *awscloudfront.CachePolicyProps
	_jsii_.StaticGet(
		"open-next-cdk.NextjsDistribution",
		"imageCachePolicyProps",
		&returns,
	)
	return returns
}

func NextjsDistribution_SetImageCachePolicyProps(val *awscloudfront.CachePolicyProps) {
	_init_.Initialize()
	if err := validateNextjsDistribution_SetImageCachePolicyPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"open-next-cdk.NextjsDistribution",
		"imageCachePolicyProps",
		val,
	)
}

func NextjsDistribution_ImageOptimizationOriginRequestPolicyProps() *awscloudfront.OriginRequestPolicyProps {
	_init_.Initialize()
	var returns *awscloudfront.OriginRequestPolicyProps
	_jsii_.StaticGet(
		"open-next-cdk.NextjsDistribution",
		"imageOptimizationOriginRequestPolicyProps",
		&returns,
	)
	return returns
}

func NextjsDistribution_SetImageOptimizationOriginRequestPolicyProps(val *awscloudfront.OriginRequestPolicyProps) {
	_init_.Initialize()
	if err := validateNextjsDistribution_SetImageOptimizationOriginRequestPolicyPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"open-next-cdk.NextjsDistribution",
		"imageOptimizationOriginRequestPolicyProps",
		val,
	)
}

func NextjsDistribution_LambdaCachePolicyProps() *awscloudfront.CachePolicyProps {
	_init_.Initialize()
	var returns *awscloudfront.CachePolicyProps
	_jsii_.StaticGet(
		"open-next-cdk.NextjsDistribution",
		"lambdaCachePolicyProps",
		&returns,
	)
	return returns
}

func NextjsDistribution_SetLambdaCachePolicyProps(val *awscloudfront.CachePolicyProps) {
	_init_.Initialize()
	if err := validateNextjsDistribution_SetLambdaCachePolicyPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"open-next-cdk.NextjsDistribution",
		"lambdaCachePolicyProps",
		val,
	)
}

func NextjsDistribution_LambdaOriginRequestPolicyProps() *awscloudfront.OriginRequestPolicyProps {
	_init_.Initialize()
	var returns *awscloudfront.OriginRequestPolicyProps
	_jsii_.StaticGet(
		"open-next-cdk.NextjsDistribution",
		"lambdaOriginRequestPolicyProps",
		&returns,
	)
	return returns
}

func NextjsDistribution_SetLambdaOriginRequestPolicyProps(val *awscloudfront.OriginRequestPolicyProps) {
	_init_.Initialize()
	if err := validateNextjsDistribution_SetLambdaOriginRequestPolicyPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"open-next-cdk.NextjsDistribution",
		"lambdaOriginRequestPolicyProps",
		val,
	)
}

func NextjsDistribution_StaticCachePolicyProps() *awscloudfront.CachePolicyProps {
	_init_.Initialize()
	var returns *awscloudfront.CachePolicyProps
	_jsii_.StaticGet(
		"open-next-cdk.NextjsDistribution",
		"staticCachePolicyProps",
		&returns,
	)
	return returns
}

func NextjsDistribution_SetStaticCachePolicyProps(val *awscloudfront.CachePolicyProps) {
	_init_.Initialize()
	if err := validateNextjsDistribution_SetStaticCachePolicyPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"open-next-cdk.NextjsDistribution",
		"staticCachePolicyProps",
		val,
	)
}

func (n *jsiiProxy_NextjsDistribution) LookupHostedZone() awsroute53.IHostedZone {
	var returns awsroute53.IHostedZone

	_jsii_.Invoke(
		n,
		"lookupHostedZone",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextjsDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextjsDistribution) ValidateCustomDomainSettings() {
	_jsii_.InvokeVoid(
		n,
		"validateCustomDomainSettings",
		nil, // no parameters
	)
}

