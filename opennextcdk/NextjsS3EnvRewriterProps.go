// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type NextjsS3EnvRewriterProps struct {
	// Relative path to the directory where the NextJS project is located.
	//
	// Can be the root of your project (`.`) or a subdirectory (`packages/web`).
	NextjsPath *string `field:"required" json:"nextjsPath" yaml:"nextjsPath"`
	// Optional value used to install NextJS node dependencies.
	//
	// It defaults to 'npx --yes open-next@latest build'.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The directory to execute `npm run build` from.
	//
	// By default, it is `nextjsPath`.
	// Can be overridden, particularly useful for monorepos where `build` is expected to run
	// at the root of the project.
	BuildPath *string `field:"optional" json:"buildPath" yaml:"buildPath"`
	// 0 - no compression, fastest 9 - maximum compression, slowest.
	CompressionLevel *float64 `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// Custom environment variables to pass to the NextJS build and runtime.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Skip building app and deploy a placeholder.
	//
	// Useful when using `next dev` for local development.
	IsPlaceholder *bool `field:"optional" json:"isPlaceholder" yaml:"isPlaceholder"`
	// Optional value for NODE_ENV during build and runtime.
	NodeEnv *string `field:"optional" json:"nodeEnv" yaml:"nodeEnv"`
	// Root of your project, if different from `nextjsPath`.
	//
	// Defaults to current working directory.
	ProjectRoot *string `field:"optional" json:"projectRoot" yaml:"projectRoot"`
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
	ReplacementConfig *RewriteReplacementsConfig `field:"required" json:"replacementConfig" yaml:"replacementConfig"`
	S3Bucket awss3.IBucket `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	S3keys *[]*string `field:"required" json:"s3keys" yaml:"s3keys"`
	CloudfrontDistributionId *string `field:"optional" json:"cloudfrontDistributionId" yaml:"cloudfrontDistributionId"`
	Debug *bool `field:"optional" json:"debug" yaml:"debug"`
}

