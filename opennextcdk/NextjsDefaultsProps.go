// Deploy a NextJS app using OpenNEXT packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Defaults for created resources.
//
// Why `any`? see https://github.com/aws/jsii/issues/2901
type NextjsDefaultsProps struct {
	// Override static file deployment settings.
	AssetDeployment interface{} `field:"optional" json:"assetDeployment" yaml:"assetDeployment"`
	// Override CloudFront distribution settings.
	//
	// These properties should all be optional but cannot be due to a limitation in jsii.
	Distribution interface{} `field:"optional" json:"distribution" yaml:"distribution"`
	// Override server lambda function settings.
	Lambda *awslambda.FunctionOptions `field:"optional" json:"lambda" yaml:"lambda"`
}

