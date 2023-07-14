package opennextcdk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"open-next-cdk.BaseSiteDomainProps",
		reflect.TypeOf((*BaseSiteDomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.BaseSiteEnvironmentOutputsInfo",
		reflect.TypeOf((*BaseSiteEnvironmentOutputsInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.BaseSiteReplaceProps",
		reflect.TypeOf((*BaseSiteReplaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.CreateArchiveArgs",
		reflect.TypeOf((*CreateArchiveArgs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"open-next-cdk.ImageOptimizationLambda",
		reflect.TypeOf((*ImageOptimizationLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSource", GoMethod: "AddEventSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSourceMapping", GoMethod: "AddEventSourceMapping"},
			_jsii_.MemberMethod{JsiiMethod: "addFunctionUrl", GoMethod: "AddFunctionUrl"},
			_jsii_.MemberMethod{JsiiMethod: "addLayers", GoMethod: "AddLayers"},
			_jsii_.MemberMethod{JsiiMethod: "addPermission", GoMethod: "AddPermission"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "canCreatePermissions", GoGetter: "CanCreatePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "configureAsyncInvoke", GoMethod: "ConfigureAsyncInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "considerWarningOnInvokeFunctionPermissions", GoMethod: "ConsiderWarningOnInvokeFunctionPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "currentVersion", GoGetter: "CurrentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterTopic", GoGetter: "DeadLetterTopic"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeUrl", GoMethod: "GrantInvokeUrl"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "invalidateVersionBasedOn", GoMethod: "InvalidateVersionBasedOn"},
			_jsii_.MemberProperty{JsiiProperty: "isBoundToVpc", GoGetter: "IsBoundToVpc"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersion", GoGetter: "LatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrors", GoMethod: "MetricErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsNode", GoGetter: "PermissionsNode"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArnsForGrantInvoke", GoGetter: "ResourceArnsForGrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "warnInvokeFunctionPermissions", GoMethod: "WarnInvokeFunctionPermissions"},
		},
		func() interface{} {
			j := jsiiProxy_ImageOptimizationLambda{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaFunction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.ImageOptimizationProps",
		reflect.TypeOf((*ImageOptimizationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"open-next-cdk.NextJsAssetsDeployment",
		reflect.TypeOf((*NextJsAssetsDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "deployments", GoGetter: "Deployments"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "prepareArchiveDirectory", GoMethod: "PrepareArchiveDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "rewriter", GoGetter: "Rewriter"},
			_jsii_.MemberProperty{JsiiProperty: "staticTempDir", GoGetter: "StaticTempDir"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NextJsAssetsDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"open-next-cdk.NextJsLambda",
		reflect.TypeOf((*NextJsLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configBucket", GoGetter: "ConfigBucket"},
			_jsii_.MemberMethod{JsiiMethod: "createConfigBucket", GoMethod: "CreateConfigBucket"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NextJsLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"open-next-cdk.Nextjs",
		reflect.TypeOf((*Nextjs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assetsDeployment", GoGetter: "AssetsDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "configBucket", GoGetter: "ConfigBucket"},
			_jsii_.MemberProperty{JsiiProperty: "distribution", GoGetter: "Distribution"},
			_jsii_.MemberProperty{JsiiProperty: "imageOptimizationFunction", GoGetter: "ImageOptimizationFunction"},
			_jsii_.MemberProperty{JsiiProperty: "imageOptimizationLambdaFunctionUrl", GoGetter: "ImageOptimizationLambdaFunctionUrl"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionUrl", GoGetter: "LambdaFunctionUrl"},
			_jsii_.MemberProperty{JsiiProperty: "nextBuild", GoGetter: "NextBuild"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "serverFunction", GoGetter: "ServerFunction"},
			_jsii_.MemberProperty{JsiiProperty: "staticAssetBucket", GoGetter: "StaticAssetBucket"},
			_jsii_.MemberProperty{JsiiProperty: "tempBuildDir", GoGetter: "TempBuildDir"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_Nextjs{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsAssetsCachePolicyProps",
		reflect.TypeOf((*NextjsAssetsCachePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsAssetsDeploymentProps",
		reflect.TypeOf((*NextjsAssetsDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsBaseProps",
		reflect.TypeOf((*NextjsBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"open-next-cdk.NextjsBuild",
		reflect.TypeOf((*NextjsBuild)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "nextImageFnDir", GoGetter: "NextImageFnDir"},
			_jsii_.MemberProperty{JsiiProperty: "nextMiddlewareFnDir", GoGetter: "NextMiddlewareFnDir"},
			_jsii_.MemberProperty{JsiiProperty: "nextServerFnDir", GoGetter: "NextServerFnDir"},
			_jsii_.MemberProperty{JsiiProperty: "nextStaticDir", GoGetter: "NextStaticDir"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "projectRoot", GoGetter: "ProjectRoot"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberMethod{JsiiMethod: "readPublicFileList", GoMethod: "ReadPublicFileList"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NextjsBuild{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsBuildProps",
		reflect.TypeOf((*NextjsBuildProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsCachePolicyProps",
		reflect.TypeOf((*NextjsCachePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsDefaultsProps",
		reflect.TypeOf((*NextjsDefaultsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"open-next-cdk.NextjsDistribution",
		reflect.TypeOf((*NextjsDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "customDomainName", GoGetter: "CustomDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "customDomainUrl", GoGetter: "CustomDomainUrl"},
			_jsii_.MemberProperty{JsiiProperty: "distribution", GoGetter: "Distribution"},
			_jsii_.MemberProperty{JsiiProperty: "distributionDomain", GoGetter: "DistributionDomain"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "hostedZone", GoGetter: "HostedZone"},
			_jsii_.MemberMethod{JsiiMethod: "lookupHostedZone", GoMethod: "LookupHostedZone"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "tempBuildDir", GoGetter: "TempBuildDir"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberMethod{JsiiMethod: "validateCustomDomainSettings", GoMethod: "ValidateCustomDomainSettings"},
		},
		func() interface{} {
			j := jsiiProxy_NextjsDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsDistributionCdkProps",
		reflect.TypeOf((*NextjsDistributionCdkProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsDistributionProps",
		reflect.TypeOf((*NextjsDistributionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsDomainProps",
		reflect.TypeOf((*NextjsDomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsLambdaProps",
		reflect.TypeOf((*NextjsLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"open-next-cdk.NextjsLayer",
		reflect.TypeOf((*NextjsLayer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPermission", GoMethod: "AddPermission"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compatibleRuntimes", GoGetter: "CompatibleRuntimes"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "layerVersionArn", GoGetter: "LayerVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NextjsLayer{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaLayerVersion)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsLayerProps",
		reflect.TypeOf((*NextjsLayerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsOriginRequestPolicyProps",
		reflect.TypeOf((*NextjsOriginRequestPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsProps",
		reflect.TypeOf((*NextjsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"open-next-cdk.NextjsS3EnvRewriter",
		reflect.TypeOf((*NextjsS3EnvRewriter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "rewriteNode", GoGetter: "RewriteNode"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NextjsS3EnvRewriter{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.NextjsS3EnvRewriterProps",
		reflect.TypeOf((*NextjsS3EnvRewriterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.RewriteReplacementsConfig",
		reflect.TypeOf((*RewriteReplacementsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"open-next-cdk.RewriterParams",
		reflect.TypeOf((*RewriterParams)(nil)).Elem(),
	)
}
