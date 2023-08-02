package opennextcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/datasprayio/open-next-cdk/opennextcdk/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/datasprayio/open-next-cdk/opennextcdk/internal"
)

// Build a lambda function from a NextJS application to handle server-side rendering, API routes, and image optimization.
type NextJsLambda interface {
	constructs.Construct
	ConfigBucket() awss3.Bucket
	SetConfigBucket(val awss3.Bucket)
	LambdaFunction() awslambda.Function
	SetLambdaFunction(val awslambda.Function)
	// The tree node.
	Node() constructs.Node
	CreateConfigBucket(replacementParams *map[string]*string) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NextJsLambda
type jsiiProxy_NextJsLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NextJsLambda) ConfigBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"configBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewNextJsLambda(scope constructs.Construct, id *string, props *NextjsLambdaProps) NextJsLambda {
	_init_.Initialize()

	if err := validateNewNextJsLambdaParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NextJsLambda{}

	_jsii_.Create(
		"open-next-cdk.NextJsLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNextJsLambda_Override(n NextJsLambda, scope constructs.Construct, id *string, props *NextjsLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"open-next-cdk.NextJsLambda",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NextJsLambda)SetConfigBucket(val awss3.Bucket) {
	_jsii_.Set(
		j,
		"configBucket",
		val,
	)
}

func (j *jsiiProxy_NextJsLambda)SetLambdaFunction(val awslambda.Function) {
	if err := j.validateSetLambdaFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFunction",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NextJsLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextJsLambda_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.NextJsLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextJsLambda) CreateConfigBucket(replacementParams *map[string]*string) *map[string]interface{} {
	if err := n.validateCreateConfigBucketParameters(replacementParams); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"createConfigBucket",
		[]interface{}{replacementParams},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextJsLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

