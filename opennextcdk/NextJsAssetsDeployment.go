// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/datasprayio/open-next-cdk/opennextcdk/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/datasprayio/open-next-cdk/opennextcdk/internal"
)

// Uploads NextJS-built static and public files to S3.
//
// Will rewrite CloudFormation references with their resolved values after uploading.
type NextJsAssetsDeployment interface {
	constructs.Construct
	// Bucket containing assets.
	Bucket() awss3.IBucket
	SetBucket(val awss3.IBucket)
	// Asset deployments to S3.
	Deployments() *[]awss3deployment.BucketDeployment
	SetDeployments(val *[]awss3deployment.BucketDeployment)
	// The tree node.
	Node() constructs.Node
	Props() *NextjsAssetsDeploymentProps
	SetProps(val *NextjsAssetsDeploymentProps)
	Rewriter() NextjsS3EnvRewriter
	SetRewriter(val NextjsS3EnvRewriter)
	StaticTempDir() *string
	SetStaticTempDir(val *string)
	PrepareArchiveDirectory() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NextJsAssetsDeployment
type jsiiProxy_NextJsAssetsDeployment struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NextJsAssetsDeployment) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsAssetsDeployment) Deployments() *[]awss3deployment.BucketDeployment {
	var returns *[]awss3deployment.BucketDeployment
	_jsii_.Get(
		j,
		"deployments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsAssetsDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsAssetsDeployment) Props() *NextjsAssetsDeploymentProps {
	var returns *NextjsAssetsDeploymentProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsAssetsDeployment) Rewriter() NextjsS3EnvRewriter {
	var returns NextjsS3EnvRewriter
	_jsii_.Get(
		j,
		"rewriter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsAssetsDeployment) StaticTempDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticTempDir",
		&returns,
	)
	return returns
}


func NewNextJsAssetsDeployment(scope constructs.Construct, id *string, props *NextjsAssetsDeploymentProps) NextJsAssetsDeployment {
	_init_.Initialize()

	if err := validateNewNextJsAssetsDeploymentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NextJsAssetsDeployment{}

	_jsii_.Create(
		"open-next-cdk.NextJsAssetsDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNextJsAssetsDeployment_Override(n NextJsAssetsDeployment, scope constructs.Construct, id *string, props *NextjsAssetsDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"open-next-cdk.NextJsAssetsDeployment",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NextJsAssetsDeployment)SetBucket(val awss3.IBucket) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_NextJsAssetsDeployment)SetDeployments(val *[]awss3deployment.BucketDeployment) {
	if err := j.validateSetDeploymentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployments",
		val,
	)
}

func (j *jsiiProxy_NextJsAssetsDeployment)SetProps(val *NextjsAssetsDeploymentProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_NextJsAssetsDeployment)SetRewriter(val NextjsS3EnvRewriter) {
	_jsii_.Set(
		j,
		"rewriter",
		val,
	)
}

func (j *jsiiProxy_NextJsAssetsDeployment)SetStaticTempDir(val *string) {
	if err := j.validateSetStaticTempDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticTempDir",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NextJsAssetsDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextJsAssetsDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.NextJsAssetsDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextJsAssetsDeployment) PrepareArchiveDirectory() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"prepareArchiveDirectory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextJsAssetsDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

