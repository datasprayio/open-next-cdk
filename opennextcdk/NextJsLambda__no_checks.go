//go:build no_runtime_type_checking

package opennextcdk

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NextJsLambda) validateCreateConfigBucketParameters(replacementParams *map[string]*string) error {
	return nil
}

func validateNextJsLambda_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NextJsLambda) validateSetLambdaFunctionParameters(val awslambda.Function) error {
	return nil
}

func validateNewNextJsLambdaParameters(scope constructs.Construct, id *string, props *NextjsLambdaProps) error {
	return nil
}

