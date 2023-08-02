//go:build no_runtime_type_checking

package opennextcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNextjs_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetAssetsDeploymentParameters(val NextJsAssetsDeployment) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetDistributionParameters(val NextjsDistribution) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetImageOptimizationFunctionParameters(val ImageOptimizationLambda) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetImageOptimizationLambdaFunctionUrlParameters(val awslambda.FunctionUrl) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetLambdaFunctionUrlParameters(val awslambda.FunctionUrl) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetNextBuildParameters(val NextjsBuild) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetPropsParameters(val *NextjsProps) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetServerFunctionParameters(val NextJsLambda) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetStaticAssetBucketParameters(val awss3.IBucket) error {
	return nil
}

func (j *jsiiProxy_Nextjs) validateSetTempBuildDirParameters(val *string) error {
	return nil
}

func validateNewNextjsParameters(scope constructs.Construct, id *string, props *NextjsProps) error {
	return nil
}

