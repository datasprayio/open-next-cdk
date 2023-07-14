//go:build no_runtime_type_checking

// Deploy a NextJS app using OpenNEXT packaging to serverless AWS using CDK
package opennextcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNextJsAssetsDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NextJsAssetsDeployment) validateSetBucketParameters(val awss3.IBucket) error {
	return nil
}

func (j *jsiiProxy_NextJsAssetsDeployment) validateSetDeploymentsParameters(val *[]awss3deployment.BucketDeployment) error {
	return nil
}

func (j *jsiiProxy_NextJsAssetsDeployment) validateSetPropsParameters(val *NextjsAssetsDeploymentProps) error {
	return nil
}

func (j *jsiiProxy_NextJsAssetsDeployment) validateSetStaticTempDirParameters(val *string) error {
	return nil
}

func validateNewNextJsAssetsDeploymentParameters(scope constructs.Construct, id *string, props *NextjsAssetsDeploymentProps) error {
	return nil
}
