package opennextcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/datasprayio/open-next-cdk/opennextcdk/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/datasprayio/open-next-cdk/opennextcdk/internal"
)

// The `Nextjs` construct is a higher level construct that makes it easy to create a NextJS app.
//
// Your standalone server application will be bundled using o(utput tracing and will be deployed to a Lambda function.
// Static assets will be deployed to an S3 bucket and served via CloudFront.
// You must use Next.js 10.3.0 or newer.
//
// Please provide a `nextjsPath` to the Next.js app inside your project.
//
// Example:
//   new Nextjs(this, "Web", {
//     nextjsPath: path.resolve("packages/web"),
//   })
//
type Nextjs interface {
	constructs.Construct
	// Asset deployment to S3.
	AssetsDeployment() NextJsAssetsDeployment
	SetAssetsDeployment(val NextJsAssetsDeployment)
	Bucket() awss3.IBucket
	ConfigBucket() awss3.Bucket
	SetConfigBucket(val awss3.Bucket)
	// CloudFront distribution.
	Distribution() NextjsDistribution
	SetDistribution(val NextjsDistribution)
	// The image optimization handler lambda function.
	ImageOptimizationFunction() ImageOptimizationLambda
	SetImageOptimizationFunction(val ImageOptimizationLambda)
	ImageOptimizationLambdaFunctionUrl() awslambda.FunctionUrl
	SetImageOptimizationLambdaFunctionUrl(val awslambda.FunctionUrl)
	LambdaFunctionUrl() awslambda.FunctionUrl
	SetLambdaFunctionUrl(val awslambda.FunctionUrl)
	// Built NextJS project output.
	NextBuild() NextjsBuild
	SetNextBuild(val NextjsBuild)
	// The tree node.
	Node() constructs.Node
	Props() *NextjsProps
	SetProps(val *NextjsProps)
	// The main NextJS server handler lambda function.
	ServerFunction() NextJsLambda
	SetServerFunction(val NextJsLambda)
	StaticAssetBucket() awss3.IBucket
	SetStaticAssetBucket(val awss3.IBucket)
	// Where build-time assets for deployment are stored.
	TempBuildDir() *string
	SetTempBuildDir(val *string)
	Url() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Nextjs
type jsiiProxy_Nextjs struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Nextjs) AssetsDeployment() NextJsAssetsDeployment {
	var returns NextJsAssetsDeployment
	_jsii_.Get(
		j,
		"assetsDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) ConfigBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"configBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) Distribution() NextjsDistribution {
	var returns NextjsDistribution
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) ImageOptimizationFunction() ImageOptimizationLambda {
	var returns ImageOptimizationLambda
	_jsii_.Get(
		j,
		"imageOptimizationFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) ImageOptimizationLambdaFunctionUrl() awslambda.FunctionUrl {
	var returns awslambda.FunctionUrl
	_jsii_.Get(
		j,
		"imageOptimizationLambdaFunctionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) LambdaFunctionUrl() awslambda.FunctionUrl {
	var returns awslambda.FunctionUrl
	_jsii_.Get(
		j,
		"lambdaFunctionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) NextBuild() NextjsBuild {
	var returns NextjsBuild
	_jsii_.Get(
		j,
		"nextBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) Props() *NextjsProps {
	var returns *NextjsProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) ServerFunction() NextJsLambda {
	var returns NextJsLambda
	_jsii_.Get(
		j,
		"serverFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) StaticAssetBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"staticAssetBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) TempBuildDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBuildDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nextjs) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewNextjs(scope constructs.Construct, id *string, props *NextjsProps) Nextjs {
	_init_.Initialize()

	if err := validateNewNextjsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Nextjs{}

	_jsii_.Create(
		"open-next-cdk.Nextjs",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNextjs_Override(n Nextjs, scope constructs.Construct, id *string, props *NextjsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"open-next-cdk.Nextjs",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_Nextjs)SetAssetsDeployment(val NextJsAssetsDeployment) {
	if err := j.validateSetAssetsDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetsDeployment",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetConfigBucket(val awss3.Bucket) {
	_jsii_.Set(
		j,
		"configBucket",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetDistribution(val NextjsDistribution) {
	if err := j.validateSetDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distribution",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetImageOptimizationFunction(val ImageOptimizationLambda) {
	if err := j.validateSetImageOptimizationFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageOptimizationFunction",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetImageOptimizationLambdaFunctionUrl(val awslambda.FunctionUrl) {
	if err := j.validateSetImageOptimizationLambdaFunctionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageOptimizationLambdaFunctionUrl",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetLambdaFunctionUrl(val awslambda.FunctionUrl) {
	if err := j.validateSetLambdaFunctionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFunctionUrl",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetNextBuild(val NextjsBuild) {
	if err := j.validateSetNextBuildParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextBuild",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetProps(val *NextjsProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetServerFunction(val NextJsLambda) {
	if err := j.validateSetServerFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverFunction",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetStaticAssetBucket(val awss3.IBucket) {
	if err := j.validateSetStaticAssetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticAssetBucket",
		val,
	)
}

func (j *jsiiProxy_Nextjs)SetTempBuildDir(val *string) {
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
func Nextjs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextjs_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.Nextjs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Nextjs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

