package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type RewriterParams struct {
	ReplacementConfig *RewriteReplacementsConfig `field:"required" json:"replacementConfig" yaml:"replacementConfig"`
	S3Bucket awss3.IBucket `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	S3keys *[]*string `field:"required" json:"s3keys" yaml:"s3keys"`
	CloudfrontDistributionId *string `field:"optional" json:"cloudfrontDistributionId" yaml:"cloudfrontDistributionId"`
	Debug *bool `field:"optional" json:"debug" yaml:"debug"`
}

