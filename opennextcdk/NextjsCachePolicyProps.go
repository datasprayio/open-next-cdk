// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

type NextjsCachePolicyProps struct {
	ImageCachePolicy awscloudfront.ICachePolicy `field:"optional" json:"imageCachePolicy" yaml:"imageCachePolicy"`
	LambdaCachePolicy awscloudfront.ICachePolicy `field:"optional" json:"lambdaCachePolicy" yaml:"lambdaCachePolicy"`
	StaticCachePolicy awscloudfront.ICachePolicy `field:"optional" json:"staticCachePolicy" yaml:"staticCachePolicy"`
	// Cache-control max-age default for static assets (/_next/*).
	//
	// Default: 30 days.
	StaticClientMaxAgeDefault awscdk.Duration `field:"optional" json:"staticClientMaxAgeDefault" yaml:"staticClientMaxAgeDefault"`
}

