// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type NextjsAssetsCachePolicyProps struct {
	// Cache-control max-age default for S3 static assets.
	//
	// Default: 30 days.
	StaticMaxAgeDefault awscdk.Duration `field:"optional" json:"staticMaxAgeDefault" yaml:"staticMaxAgeDefault"`
	// Cache-control stale-while-revalidate default for S3 static assets.
	//
	// Default: 1 day.
	StaticStaleWhileRevalidateDefault awscdk.Duration `field:"optional" json:"staticStaleWhileRevalidateDefault" yaml:"staticStaleWhileRevalidateDefault"`
}

