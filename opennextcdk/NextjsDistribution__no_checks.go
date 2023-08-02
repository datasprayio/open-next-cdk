//go:build no_runtime_type_checking

package opennextcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNextjsDistribution_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NextjsDistribution) validateSetDistributionParameters(val awscloudfront.Distribution) error {
	return nil
}

func (j *jsiiProxy_NextjsDistribution) validateSetPropsParameters(val *NextjsDistributionProps) error {
	return nil
}

func (j *jsiiProxy_NextjsDistribution) validateSetTempBuildDirParameters(val *string) error {
	return nil
}

func validateNextjsDistribution_SetFallbackOriginRequestPolicyPropsParameters(val *awscloudfront.OriginRequestPolicyProps) error {
	return nil
}

func validateNextjsDistribution_SetImageCachePolicyPropsParameters(val *awscloudfront.CachePolicyProps) error {
	return nil
}

func validateNextjsDistribution_SetImageOptimizationOriginRequestPolicyPropsParameters(val *awscloudfront.OriginRequestPolicyProps) error {
	return nil
}

func validateNextjsDistribution_SetLambdaCachePolicyPropsParameters(val *awscloudfront.CachePolicyProps) error {
	return nil
}

func validateNextjsDistribution_SetLambdaOriginRequestPolicyPropsParameters(val *awscloudfront.OriginRequestPolicyProps) error {
	return nil
}

func validateNextjsDistribution_SetStaticCachePolicyPropsParameters(val *awscloudfront.CachePolicyProps) error {
	return nil
}

func validateNewNextjsDistributionParameters(scope constructs.Construct, id *string, props *NextjsDistributionProps) error {
	return nil
}

