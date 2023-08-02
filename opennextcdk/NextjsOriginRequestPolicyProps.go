package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

type NextjsOriginRequestPolicyProps struct {
	FallbackOriginRequestPolicy awscloudfront.IOriginRequestPolicy `field:"optional" json:"fallbackOriginRequestPolicy" yaml:"fallbackOriginRequestPolicy"`
	ImageOptimizationOriginRequestPolicy awscloudfront.IOriginRequestPolicy `field:"optional" json:"imageOptimizationOriginRequestPolicy" yaml:"imageOptimizationOriginRequestPolicy"`
	LambdaOriginRequestPolicy awscloudfront.IOriginRequestPolicy `field:"optional" json:"lambdaOriginRequestPolicy" yaml:"lambdaOriginRequestPolicy"`
}

