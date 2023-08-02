package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type NextjsLambdaProps struct {
	// Optional value used to install NextJS node dependencies.
	//
	// It defaults to 'npx --yes open-next@1 build'.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The directory to execute `npm run build` from.
	//
	// By default, it is `nextjsPath`.
	// Can be overridden, particularly useful for monorepos where `build` is expected to run
	// at the root of the project.
	BuildPath *string `field:"optional" json:"buildPath" yaml:"buildPath"`
	// 0 - no compression, fastest 9 - maximum compression, slowest.
	// Default: 1.
	//
	CompressionLevel *float64 `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// Custom environment variables to pass to the NextJS build and runtime.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Used in conjunction with nextJsPath to skip building NextJS app and assume .open-next folder already exists. Useful when using `next dev` for local development.
	// Deprecated: use `openNextPath` instead.
	IsPlaceholder *bool `field:"optional" json:"isPlaceholder" yaml:"isPlaceholder"`
	// Relative path to the directory where the NextJS project is located.
	//
	// Can be the root of your project (`.`) or a subdirectory (`packages/web`).
	//
	// One of `openNextPath`, `nextJsPath` or `nextjsPath` must be supplied.
	// Deprecated: use `nextJsPath` instead.
	NextjsPath *string `field:"optional" json:"nextjsPath" yaml:"nextjsPath"`
	// Relative path to the directory where the NextJS project is located.
	//
	// Can be the root of your project (`.`) or a subdirectory (`packages/web`).
	//
	// One of `openNextPath` or `nextJsPath` must be supplied.
	NextJsPath *string `field:"optional" json:"nextJsPath" yaml:"nextJsPath"`
	// Optional value for NODE_ENV during build and runtime.
	NodeEnv *string `field:"optional" json:"nodeEnv" yaml:"nodeEnv"`
	// Relative path to the OpenNext package named `.open-next` by default.
	//
	// One of `openNextPath` or `nextJsPath` must be supplied.
	OpenNextPath *string `field:"optional" json:"openNextPath" yaml:"openNextPath"`
	// Less build output.
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
	// Optional arn for the sharp lambda layer.
	//
	// If omitted, the layer will be created.
	SharpLayerArn *string `field:"optional" json:"sharpLayerArn" yaml:"sharpLayerArn"`
	// Directory to store temporary build files in.
	//
	// Defaults to os.tmpdir().
	TempBuildDir *string `field:"optional" json:"tempBuildDir" yaml:"tempBuildDir"`
	// Built nextJS application.
	NextBuild NextjsBuild `field:"required" json:"nextBuild" yaml:"nextBuild"`
	// Override function properties.
	Lambda *awslambda.FunctionOptions `field:"optional" json:"lambda" yaml:"lambda"`
}

