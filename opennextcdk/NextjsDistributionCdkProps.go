// Deploy a NextJS app using OpenNEXT packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

type NextjsDistributionCdkProps struct {
	// Pass in a value to override the default settings this construct uses to create the CloudFront `Distribution` internally.
	Distribution *awscloudfront.DistributionProps `field:"optional" json:"distribution" yaml:"distribution"`
}

