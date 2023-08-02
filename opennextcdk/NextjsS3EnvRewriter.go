package opennextcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/datasprayio/open-next-cdk/opennextcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/datasprayio/open-next-cdk/opennextcdk/internal"
)

// Rewrites variables in S3 objects after a deployment happens to replace CloudFormation tokens with their values.
//
// These values are not resolved at build time because they are
// only known at deploy time.
type NextjsS3EnvRewriter interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	RewriteNode() constructs.Construct
	SetRewriteNode(val constructs.Construct)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NextjsS3EnvRewriter
type jsiiProxy_NextjsS3EnvRewriter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NextjsS3EnvRewriter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsS3EnvRewriter) RewriteNode() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"rewriteNode",
		&returns,
	)
	return returns
}


func NewNextjsS3EnvRewriter(scope constructs.Construct, id *string, props *NextjsS3EnvRewriterProps) NextjsS3EnvRewriter {
	_init_.Initialize()

	if err := validateNewNextjsS3EnvRewriterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NextjsS3EnvRewriter{}

	_jsii_.Create(
		"open-next-cdk.NextjsS3EnvRewriter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNextjsS3EnvRewriter_Override(n NextjsS3EnvRewriter, scope constructs.Construct, id *string, props *NextjsS3EnvRewriterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"open-next-cdk.NextjsS3EnvRewriter",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NextjsS3EnvRewriter)SetRewriteNode(val constructs.Construct) {
	_jsii_.Set(
		j,
		"rewriteNode",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NextjsS3EnvRewriter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextjsS3EnvRewriter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.NextjsS3EnvRewriter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextjsS3EnvRewriter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

