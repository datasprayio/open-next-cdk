//go:build no_runtime_type_checking

// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NextjsLayer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (n *jsiiProxy_NextjsLayer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NextjsLayer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NextjsLayer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNextjsLayer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateNextjsLayer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateNextjsLayer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNextjsLayer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNextjsLayer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNextjsLayerParameters(scope constructs.Construct, id *string, props *NextjsLayerProps) error {
	return nil
}

