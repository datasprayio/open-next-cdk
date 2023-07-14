//go:build no_runtime_type_checking

// Deploy a NextJS app using OpenNEXT packaging to serverless AWS using CDK
package opennextcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNextjsBuild_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NextjsBuild) validateSetNextImageFnDirParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NextjsBuild) validateSetNextServerFnDirParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NextjsBuild) validateSetNextStaticDirParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NextjsBuild) validateSetProjectRootParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NextjsBuild) validateSetPropsParameters(val *NextjsBuildProps) error {
	return nil
}

func validateNewNextjsBuildParameters(scope constructs.Construct, id *string, props *NextjsBuildProps) error {
	return nil
}

