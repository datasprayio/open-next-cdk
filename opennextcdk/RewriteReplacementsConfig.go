// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type RewriteReplacementsConfig struct {
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	JsonS3Bucket awss3.IBucket `field:"optional" json:"jsonS3Bucket" yaml:"jsonS3Bucket"`
	JsonS3Key *string `field:"optional" json:"jsonS3Key" yaml:"jsonS3Key"`
}

