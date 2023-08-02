package opennextcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/datasprayio/open-next-cdk/opennextcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/datasprayio/open-next-cdk/opennextcdk/internal"
)

// Represents a built NextJS application.
//
// This construct runs `npm build` in standalone output mode inside your `nextjsPath`.
// This construct can be used by higher level constructs or used directly.
type NextjsBuild interface {
	constructs.Construct
	// Contains function for processessing image requests.
	//
	// Should be arm64.
	NextImageFnDir() *string
	SetNextImageFnDir(val *string)
	// Contains code for middleware.
	//
	// Not currently used.
	NextMiddlewareFnDir() *string
	SetNextMiddlewareFnDir(val *string)
	// Contains server code and dependencies.
	NextServerFnDir() *string
	SetNextServerFnDir(val *string)
	// Static files containing client-side code.
	NextStaticDir() *string
	SetNextStaticDir(val *string)
	// The tree node.
	Node() constructs.Node
	OpenNextPath() *string
	Props() *NextjsBuildProps
	SetProps(val *NextjsBuildProps)
	ReadPublicFileList() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NextjsBuild
type jsiiProxy_NextjsBuild struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NextjsBuild) NextImageFnDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextImageFnDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsBuild) NextMiddlewareFnDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMiddlewareFnDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsBuild) NextServerFnDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextServerFnDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsBuild) NextStaticDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextStaticDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsBuild) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsBuild) OpenNextPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openNextPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextjsBuild) Props() *NextjsBuildProps {
	var returns *NextjsBuildProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewNextjsBuild(scope constructs.Construct, id *string, props *NextjsBuildProps) NextjsBuild {
	_init_.Initialize()

	if err := validateNewNextjsBuildParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NextjsBuild{}

	_jsii_.Create(
		"open-next-cdk.NextjsBuild",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNextjsBuild_Override(n NextjsBuild, scope constructs.Construct, id *string, props *NextjsBuildProps) {
	_init_.Initialize()

	_jsii_.Create(
		"open-next-cdk.NextjsBuild",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NextjsBuild)SetNextImageFnDir(val *string) {
	if err := j.validateSetNextImageFnDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextImageFnDir",
		val,
	)
}

func (j *jsiiProxy_NextjsBuild)SetNextMiddlewareFnDir(val *string) {
	_jsii_.Set(
		j,
		"nextMiddlewareFnDir",
		val,
	)
}

func (j *jsiiProxy_NextjsBuild)SetNextServerFnDir(val *string) {
	if err := j.validateSetNextServerFnDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextServerFnDir",
		val,
	)
}

func (j *jsiiProxy_NextjsBuild)SetNextStaticDir(val *string) {
	if err := j.validateSetNextStaticDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextStaticDir",
		val,
	)
}

func (j *jsiiProxy_NextjsBuild)SetProps(val *NextjsBuildProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NextjsBuild_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextjsBuild_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.NextjsBuild",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextjsBuild) ReadPublicFileList() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"readPublicFileList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextjsBuild) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

