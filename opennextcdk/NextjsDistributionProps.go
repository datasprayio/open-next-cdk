// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type NextjsDistributionProps struct {
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
	// Lambda function to optimize images.
	//
	// Must be provided if you want to serve dynamic requests.
	ImageOptFunction awslambda.IFunction `field:"required" json:"imageOptFunction" yaml:"imageOptFunction"`
	// Built NextJS app.
	NextBuild NextjsBuild `field:"required" json:"nextBuild" yaml:"nextBuild"`
	// Lambda function to route all non-static requests to.
	//
	// Must be provided if you want to serve dynamic requests.
	ServerFunction awslambda.IFunction `field:"required" json:"serverFunction" yaml:"serverFunction"`
	// Bucket containing static assets.
	//
	// Must be provided if you want to serve static files.
	StaticAssetsBucket awss3.IBucket `field:"required" json:"staticAssetsBucket" yaml:"staticAssetsBucket"`
	// Override the default CloudFront cache policies created internally.
	CachePolicies *NextjsCachePolicyProps `field:"optional" json:"cachePolicies" yaml:"cachePolicies"`
	// Overrides for created CDK resources.
	Cdk *NextjsDistributionCdkProps `field:"optional" json:"cdk" yaml:"cdk"`
	// The customDomain for this website. Supports domains that are hosted either on [Route 53](https://aws.amazon.com/route53/) or externally.
	//
	// Note that you can also migrate externally hosted domains to Route 53 by
	// [following this guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html).
	//
	// Example:
	//   new NextjsDistribution(this, "Dist", {
	//     customDomain: "domain.com",
	//   });
	//
	//   new NextjsDistribution(this, "Dist", {
	//     customDomain: {
	//       domainName: "domain.com",
	//       domainAlias: "www.domain.com",
	//       hostedZone: "domain.com"
	//     },
	//   });
	//
	CustomDomain interface{} `field:"optional" json:"customDomain" yaml:"customDomain"`
	// Override lambda function url auth type.
	FunctionUrlAuthType awslambda.FunctionUrlAuthType `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// Override the default CloudFront origin request policies created internally.
	OriginRequestPolicies *NextjsOriginRequestPolicyProps `field:"optional" json:"originRequestPolicies" yaml:"originRequestPolicies"`
	// Optional value to prefix the edge function stack It defaults to "Nextjs".
	StackPrefix *string `field:"optional" json:"stackPrefix" yaml:"stackPrefix"`
	// Include the name of your deployment stage if present.
	//
	// Used to name the edge functions stack.
	// Required if using SST.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

