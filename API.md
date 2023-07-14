# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### ImageOptimizationLambda <a name="ImageOptimizationLambda" id="open-next-cdk.ImageOptimizationLambda"></a>

This lambda handles image optimization.

#### Initializers <a name="Initializers" id="open-next-cdk.ImageOptimizationLambda.Initializer"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

new ImageOptimizationLambda(scope: Construct, id: string, props: ImageOptimizationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.ImageOptimizationProps">ImageOptimizationProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.ImageOptimizationLambda.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.ImageOptimizationLambda.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.ImageOptimizationProps">ImageOptimizationProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addEventSource">addEventSource</a></code> | Adds an event source to this function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addEventSourceMapping">addEventSourceMapping</a></code> | Adds an event source that maps to this AWS Lambda function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addFunctionUrl">addFunctionUrl</a></code> | Adds a url to this lambda function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addPermission">addPermission</a></code> | Adds a permission to the Lambda resource policy. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addToRolePolicy">addToRolePolicy</a></code> | Adds a statement to the IAM role assumed by the instance. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.configureAsyncInvoke">configureAsyncInvoke</a></code> | Configures options for asynchronous invocation. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.considerWarningOnInvokeFunctionPermissions">considerWarningOnInvokeFunctionPermissions</a></code> | A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.grantInvoke">grantInvoke</a></code> | Grant the given identity permissions to invoke this Lambda. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.grantInvokeUrl">grantInvokeUrl</a></code> | Grant the given identity permissions to invoke this Lambda Function URL. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metric">metric</a></code> | Return the given named metric for this Function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricDuration">metricDuration</a></code> | How long execution of this Lambda takes. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricErrors">metricErrors</a></code> | How many invocations of this Lambda fail. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricInvocations">metricInvocations</a></code> | How often this Lambda is invoked. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricThrottles">metricThrottles</a></code> | How often this Lambda is throttled. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addAlias">addAlias</a></code> | Defines an alias for this function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addEnvironment">addEnvironment</a></code> | Adds an environment variable to this Lambda function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.addLayers">addLayers</a></code> | Adds one or more Lambda Layers to this Lambda function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.invalidateVersionBasedOn">invalidateVersionBasedOn</a></code> | Mix additional information into the hash of the Version object. |

---

##### `toString` <a name="toString" id="open-next-cdk.ImageOptimizationLambda.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="open-next-cdk.ImageOptimizationLambda.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy: RemovalPolicy): void
```

Apply the given removal policy to this resource.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).

###### `policy`<sup>Required</sup> <a name="policy" id="open-next-cdk.ImageOptimizationLambda.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

##### `addEventSource` <a name="addEventSource" id="open-next-cdk.ImageOptimizationLambda.addEventSource"></a>

```typescript
public addEventSource(source: IEventSource): void
```

Adds an event source to this function.

Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.

The following example adds an SQS Queue as an event source:
```
import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
myFunction.addEventSource(new SqsEventSource(myQueue));
```

###### `source`<sup>Required</sup> <a name="source" id="open-next-cdk.ImageOptimizationLambda.addEventSource.parameter.source"></a>

- *Type:* aws-cdk-lib.aws_lambda.IEventSource

---

##### `addEventSourceMapping` <a name="addEventSourceMapping" id="open-next-cdk.ImageOptimizationLambda.addEventSourceMapping"></a>

```typescript
public addEventSourceMapping(id: string, options: EventSourceMappingOptions): EventSourceMapping
```

Adds an event source that maps to this AWS Lambda function.

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.ImageOptimizationLambda.addEventSourceMapping.parameter.id"></a>

- *Type:* string

---

###### `options`<sup>Required</sup> <a name="options" id="open-next-cdk.ImageOptimizationLambda.addEventSourceMapping.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.EventSourceMappingOptions

---

##### `addFunctionUrl` <a name="addFunctionUrl" id="open-next-cdk.ImageOptimizationLambda.addFunctionUrl"></a>

```typescript
public addFunctionUrl(options?: FunctionUrlOptions): FunctionUrl
```

Adds a url to this lambda function.

###### `options`<sup>Optional</sup> <a name="options" id="open-next-cdk.ImageOptimizationLambda.addFunctionUrl.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.FunctionUrlOptions

---

##### `addPermission` <a name="addPermission" id="open-next-cdk.ImageOptimizationLambda.addPermission"></a>

```typescript
public addPermission(id: string, permission: Permission): void
```

Adds a permission to the Lambda resource policy.

> [Permission for details.](Permission for details.)

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.ImageOptimizationLambda.addPermission.parameter.id"></a>

- *Type:* string

The id for the permission construct.

---

###### `permission`<sup>Required</sup> <a name="permission" id="open-next-cdk.ImageOptimizationLambda.addPermission.parameter.permission"></a>

- *Type:* aws-cdk-lib.aws_lambda.Permission

The permission to grant to this Lambda function.

---

##### `addToRolePolicy` <a name="addToRolePolicy" id="open-next-cdk.ImageOptimizationLambda.addToRolePolicy"></a>

```typescript
public addToRolePolicy(statement: PolicyStatement): void
```

Adds a statement to the IAM role assumed by the instance.

###### `statement`<sup>Required</sup> <a name="statement" id="open-next-cdk.ImageOptimizationLambda.addToRolePolicy.parameter.statement"></a>

- *Type:* aws-cdk-lib.aws_iam.PolicyStatement

---

##### `configureAsyncInvoke` <a name="configureAsyncInvoke" id="open-next-cdk.ImageOptimizationLambda.configureAsyncInvoke"></a>

```typescript
public configureAsyncInvoke(options: EventInvokeConfigOptions): void
```

Configures options for asynchronous invocation.

###### `options`<sup>Required</sup> <a name="options" id="open-next-cdk.ImageOptimizationLambda.configureAsyncInvoke.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.EventInvokeConfigOptions

---

##### `considerWarningOnInvokeFunctionPermissions` <a name="considerWarningOnInvokeFunctionPermissions" id="open-next-cdk.ImageOptimizationLambda.considerWarningOnInvokeFunctionPermissions"></a>

```typescript
public considerWarningOnInvokeFunctionPermissions(scope: Construct, action: string): void
```

A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function.

function.currentVersion is invoked before or after the permission is created.

This applies only to permissions on Lambda functions, not versions or aliases.
This function is overridden as a noOp for QualifiedFunctionBase.

###### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.ImageOptimizationLambda.considerWarningOnInvokeFunctionPermissions.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `action`<sup>Required</sup> <a name="action" id="open-next-cdk.ImageOptimizationLambda.considerWarningOnInvokeFunctionPermissions.parameter.action"></a>

- *Type:* string

---

##### `grantInvoke` <a name="grantInvoke" id="open-next-cdk.ImageOptimizationLambda.grantInvoke"></a>

```typescript
public grantInvoke(grantee: IGrantable): Grant
```

Grant the given identity permissions to invoke this Lambda.

###### `grantee`<sup>Required</sup> <a name="grantee" id="open-next-cdk.ImageOptimizationLambda.grantInvoke.parameter.grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

##### `grantInvokeUrl` <a name="grantInvokeUrl" id="open-next-cdk.ImageOptimizationLambda.grantInvokeUrl"></a>

```typescript
public grantInvokeUrl(grantee: IGrantable): Grant
```

Grant the given identity permissions to invoke this Lambda Function URL.

###### `grantee`<sup>Required</sup> <a name="grantee" id="open-next-cdk.ImageOptimizationLambda.grantInvokeUrl.parameter.grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

##### `metric` <a name="metric" id="open-next-cdk.ImageOptimizationLambda.metric"></a>

```typescript
public metric(metricName: string, props?: MetricOptions): Metric
```

Return the given named metric for this Function.

###### `metricName`<sup>Required</sup> <a name="metricName" id="open-next-cdk.ImageOptimizationLambda.metric.parameter.metricName"></a>

- *Type:* string

---

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metric.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricDuration` <a name="metricDuration" id="open-next-cdk.ImageOptimizationLambda.metricDuration"></a>

```typescript
public metricDuration(props?: MetricOptions): Metric
```

How long execution of this Lambda takes.

Average over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricDuration.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricErrors` <a name="metricErrors" id="open-next-cdk.ImageOptimizationLambda.metricErrors"></a>

```typescript
public metricErrors(props?: MetricOptions): Metric
```

How many invocations of this Lambda fail.

Sum over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricErrors.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricInvocations` <a name="metricInvocations" id="open-next-cdk.ImageOptimizationLambda.metricInvocations"></a>

```typescript
public metricInvocations(props?: MetricOptions): Metric
```

How often this Lambda is invoked.

Sum over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricInvocations.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricThrottles` <a name="metricThrottles" id="open-next-cdk.ImageOptimizationLambda.metricThrottles"></a>

```typescript
public metricThrottles(props?: MetricOptions): Metric
```

How often this Lambda is throttled.

Sum over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricThrottles.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `addAlias` <a name="addAlias" id="open-next-cdk.ImageOptimizationLambda.addAlias"></a>

```typescript
public addAlias(aliasName: string, options?: AliasOptions): Alias
```

Defines an alias for this function.

The alias will automatically be updated to point to the latest version of
the function as it is being updated during a deployment.

```ts
declare const fn: lambda.Function;

fn.addAlias('Live');

// Is equivalent to

new lambda.Alias(this, 'AliasLive', {
  aliasName: 'Live',
  version: fn.currentVersion,
});
```

###### `aliasName`<sup>Required</sup> <a name="aliasName" id="open-next-cdk.ImageOptimizationLambda.addAlias.parameter.aliasName"></a>

- *Type:* string

The name of the alias.

---

###### `options`<sup>Optional</sup> <a name="options" id="open-next-cdk.ImageOptimizationLambda.addAlias.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.AliasOptions

Alias options.

---

##### `addEnvironment` <a name="addEnvironment" id="open-next-cdk.ImageOptimizationLambda.addEnvironment"></a>

```typescript
public addEnvironment(key: string, value: string, options?: EnvironmentOptions): Function
```

Adds an environment variable to this Lambda function.

If this is a ref to a Lambda function, this operation results in a no-op.

###### `key`<sup>Required</sup> <a name="key" id="open-next-cdk.ImageOptimizationLambda.addEnvironment.parameter.key"></a>

- *Type:* string

The environment variable key.

---

###### `value`<sup>Required</sup> <a name="value" id="open-next-cdk.ImageOptimizationLambda.addEnvironment.parameter.value"></a>

- *Type:* string

The environment variable's value.

---

###### `options`<sup>Optional</sup> <a name="options" id="open-next-cdk.ImageOptimizationLambda.addEnvironment.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.EnvironmentOptions

Environment variable options.

---

##### `addLayers` <a name="addLayers" id="open-next-cdk.ImageOptimizationLambda.addLayers"></a>

```typescript
public addLayers(layers: ILayerVersion): void
```

Adds one or more Lambda Layers to this Lambda function.

###### `layers`<sup>Required</sup> <a name="layers" id="open-next-cdk.ImageOptimizationLambda.addLayers.parameter.layers"></a>

- *Type:* aws-cdk-lib.aws_lambda.ILayerVersion

the layers to be added.

---

##### `invalidateVersionBasedOn` <a name="invalidateVersionBasedOn" id="open-next-cdk.ImageOptimizationLambda.invalidateVersionBasedOn"></a>

```typescript
public invalidateVersionBasedOn(x: string): void
```

Mix additional information into the hash of the Version object.

The Lambda Function construct does its best to automatically create a new
Version when anything about the Function changes (its code, its layers,
any of the other properties).

However, you can sometimes source information from places that the CDK cannot
look into, like the deploy-time values of SSM parameters. In those cases,
the CDK would not force the creation of a new Version object when it actually
should.

This method can be used to invalidate the current Version object. Pass in
any string into this method, and make sure the string changes when you know
a new Version needs to be created.

This method may be called more than once.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.ImageOptimizationLambda.invalidateVersionBasedOn.parameter.x"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.isOwnedResource">isOwnedResource</a></code> | Returns true if the construct was created by CDK, and false otherwise. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.isResource">isResource</a></code> | Check whether the given construct is a Resource. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.classifyVersionProperty">classifyVersionProperty</a></code> | Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.fromFunctionArn">fromFunctionArn</a></code> | Import a lambda function into the CDK using its ARN. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.fromFunctionAttributes">fromFunctionAttributes</a></code> | Creates a Lambda function object which represents a function not defined within this stack. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.fromFunctionName">fromFunctionName</a></code> | Import a lambda function into the CDK using its name. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricAll">metricAll</a></code> | Return the given named metric for this Lambda. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricAllConcurrentExecutions">metricAllConcurrentExecutions</a></code> | Metric for the number of concurrent executions across all Lambdas. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricAllDuration">metricAllDuration</a></code> | Metric for the Duration executing all Lambdas. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricAllErrors">metricAllErrors</a></code> | Metric for the number of Errors executing all Lambdas. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricAllInvocations">metricAllInvocations</a></code> | Metric for the number of invocations of all Lambdas. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricAllThrottles">metricAllThrottles</a></code> | Metric for the number of throttled invocations of all Lambdas. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.metricAllUnreservedConcurrentExecutions">metricAllUnreservedConcurrentExecutions</a></code> | Metric for the number of unreserved concurrent executions across all Lambdas. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.ImageOptimizationLambda.isConstruct"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.ImageOptimizationLambda.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isOwnedResource` <a name="isOwnedResource" id="open-next-cdk.ImageOptimizationLambda.isOwnedResource"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.isOwnedResource(construct: IConstruct)
```

Returns true if the construct was created by CDK, and false otherwise.

###### `construct`<sup>Required</sup> <a name="construct" id="open-next-cdk.ImageOptimizationLambda.isOwnedResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

##### `isResource` <a name="isResource" id="open-next-cdk.ImageOptimizationLambda.isResource"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.isResource(construct: IConstruct)
```

Check whether the given construct is a Resource.

###### `construct`<sup>Required</sup> <a name="construct" id="open-next-cdk.ImageOptimizationLambda.isResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

##### `classifyVersionProperty` <a name="classifyVersionProperty" id="open-next-cdk.ImageOptimizationLambda.classifyVersionProperty"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.classifyVersionProperty(propertyName: string, locked: boolean)
```

Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.

See 'currentVersion' section in the module README for more details.

###### `propertyName`<sup>Required</sup> <a name="propertyName" id="open-next-cdk.ImageOptimizationLambda.classifyVersionProperty.parameter.propertyName"></a>

- *Type:* string

The property to classify.

---

###### `locked`<sup>Required</sup> <a name="locked" id="open-next-cdk.ImageOptimizationLambda.classifyVersionProperty.parameter.locked"></a>

- *Type:* boolean

whether the property should be associated to the version or not.

---

##### `fromFunctionArn` <a name="fromFunctionArn" id="open-next-cdk.ImageOptimizationLambda.fromFunctionArn"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.fromFunctionArn(scope: Construct, id: string, functionArn: string)
```

Import a lambda function into the CDK using its ARN.

###### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.ImageOptimizationLambda.fromFunctionArn.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.ImageOptimizationLambda.fromFunctionArn.parameter.id"></a>

- *Type:* string

---

###### `functionArn`<sup>Required</sup> <a name="functionArn" id="open-next-cdk.ImageOptimizationLambda.fromFunctionArn.parameter.functionArn"></a>

- *Type:* string

---

##### `fromFunctionAttributes` <a name="fromFunctionAttributes" id="open-next-cdk.ImageOptimizationLambda.fromFunctionAttributes"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.fromFunctionAttributes(scope: Construct, id: string, attrs: FunctionAttributes)
```

Creates a Lambda function object which represents a function not defined within this stack.

###### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.ImageOptimizationLambda.fromFunctionAttributes.parameter.scope"></a>

- *Type:* constructs.Construct

The parent construct.

---

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.ImageOptimizationLambda.fromFunctionAttributes.parameter.id"></a>

- *Type:* string

The name of the lambda construct.

---

###### `attrs`<sup>Required</sup> <a name="attrs" id="open-next-cdk.ImageOptimizationLambda.fromFunctionAttributes.parameter.attrs"></a>

- *Type:* aws-cdk-lib.aws_lambda.FunctionAttributes

the attributes of the function to import.

---

##### `fromFunctionName` <a name="fromFunctionName" id="open-next-cdk.ImageOptimizationLambda.fromFunctionName"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.fromFunctionName(scope: Construct, id: string, functionName: string)
```

Import a lambda function into the CDK using its name.

###### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.ImageOptimizationLambda.fromFunctionName.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.ImageOptimizationLambda.fromFunctionName.parameter.id"></a>

- *Type:* string

---

###### `functionName`<sup>Required</sup> <a name="functionName" id="open-next-cdk.ImageOptimizationLambda.fromFunctionName.parameter.functionName"></a>

- *Type:* string

---

##### `metricAll` <a name="metricAll" id="open-next-cdk.ImageOptimizationLambda.metricAll"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.metricAll(metricName: string, props?: MetricOptions)
```

Return the given named metric for this Lambda.

###### `metricName`<sup>Required</sup> <a name="metricName" id="open-next-cdk.ImageOptimizationLambda.metricAll.parameter.metricName"></a>

- *Type:* string

---

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricAll.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricAllConcurrentExecutions` <a name="metricAllConcurrentExecutions" id="open-next-cdk.ImageOptimizationLambda.metricAllConcurrentExecutions"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.metricAllConcurrentExecutions(props?: MetricOptions)
```

Metric for the number of concurrent executions across all Lambdas.

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricAllConcurrentExecutions.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricAllDuration` <a name="metricAllDuration" id="open-next-cdk.ImageOptimizationLambda.metricAllDuration"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.metricAllDuration(props?: MetricOptions)
```

Metric for the Duration executing all Lambdas.

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricAllDuration.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricAllErrors` <a name="metricAllErrors" id="open-next-cdk.ImageOptimizationLambda.metricAllErrors"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.metricAllErrors(props?: MetricOptions)
```

Metric for the number of Errors executing all Lambdas.

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricAllErrors.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricAllInvocations` <a name="metricAllInvocations" id="open-next-cdk.ImageOptimizationLambda.metricAllInvocations"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.metricAllInvocations(props?: MetricOptions)
```

Metric for the number of invocations of all Lambdas.

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricAllInvocations.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricAllThrottles` <a name="metricAllThrottles" id="open-next-cdk.ImageOptimizationLambda.metricAllThrottles"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.metricAllThrottles(props?: MetricOptions)
```

Metric for the number of throttled invocations of all Lambdas.

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricAllThrottles.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricAllUnreservedConcurrentExecutions` <a name="metricAllUnreservedConcurrentExecutions" id="open-next-cdk.ImageOptimizationLambda.metricAllUnreservedConcurrentExecutions"></a>

```typescript
import { ImageOptimizationLambda } from 'open-next-cdk'

ImageOptimizationLambda.metricAllUnreservedConcurrentExecutions(props?: MetricOptions)
```

Metric for the number of unreserved concurrent executions across all Lambdas.

###### `props`<sup>Optional</sup> <a name="props" id="open-next-cdk.ImageOptimizationLambda.metricAllUnreservedConcurrentExecutions.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this resource is defined. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.architecture">architecture</a></code> | <code>aws-cdk-lib.aws_lambda.Architecture</code> | The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64). |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.connections">connections</a></code> | <code>aws-cdk-lib.aws_ec2.Connections</code> | Access the Connections object. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.functionArn">functionArn</a></code> | <code>string</code> | ARN of this function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.functionName">functionName</a></code> | <code>string</code> | Name of this function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.grantPrincipal">grantPrincipal</a></code> | <code>aws-cdk-lib.aws_iam.IPrincipal</code> | The principal this Lambda Function is running as. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.isBoundToVpc">isBoundToVpc</a></code> | <code>boolean</code> | Whether or not this Lambda function was bound to a VPC. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.latestVersion">latestVersion</a></code> | <code>aws-cdk-lib.aws_lambda.IVersion</code> | The `$LATEST` version of this function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.permissionsNode">permissionsNode</a></code> | <code>constructs.Node</code> | The construct node where permissions are attached. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.resourceArnsForGrantInvoke">resourceArnsForGrantInvoke</a></code> | <code>string[]</code> | The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke(). |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.role">role</a></code> | <code>aws-cdk-lib.aws_iam.IRole</code> | Execution role associated with this function. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.currentVersion">currentVersion</a></code> | <code>aws-cdk-lib.aws_lambda.Version</code> | Returns a `lambda.Version` which represents the current version of this Lambda function. A new version will be created every time the function's configuration changes. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.logGroup">logGroup</a></code> | <code>aws-cdk-lib.aws_logs.ILogGroup</code> | The LogGroup where the Lambda function's logs are made available. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.runtime">runtime</a></code> | <code>aws-cdk-lib.aws_lambda.Runtime</code> | The runtime configured for this lambda. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.deadLetterQueue">deadLetterQueue</a></code> | <code>aws-cdk-lib.aws_sqs.IQueue</code> | The DLQ (as queue) associated with this Lambda Function (this is an optional attribute). |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.deadLetterTopic">deadLetterTopic</a></code> | <code>aws-cdk-lib.aws_sns.ITopic</code> | The DLQ (as topic) associated with this Lambda Function (this is an optional attribute). |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.timeout">timeout</a></code> | <code>aws-cdk-lib.Duration</code> | The timeout configured for this lambda. |
| <code><a href="#open-next-cdk.ImageOptimizationLambda.property.bucket">bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.ImageOptimizationLambda.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `env`<sup>Required</sup> <a name="env" id="open-next-cdk.ImageOptimizationLambda.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* aws-cdk-lib.ResourceEnvironment

The environment this resource belongs to.

For resources that are created and managed by the CDK
(generally, those created by creating new class instances like Role, Bucket, etc.),
this is always the same as the environment of the stack they belong to;
however, for imported resources
(those obtained from static methods like fromRoleArn, fromBucketName, etc.),
that might be different than the stack they were imported into.

---

##### `stack`<sup>Required</sup> <a name="stack" id="open-next-cdk.ImageOptimizationLambda.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this resource is defined.

---

##### `architecture`<sup>Required</sup> <a name="architecture" id="open-next-cdk.ImageOptimizationLambda.property.architecture"></a>

```typescript
public readonly architecture: Architecture;
```

- *Type:* aws-cdk-lib.aws_lambda.Architecture

The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64).

---

##### `connections`<sup>Required</sup> <a name="connections" id="open-next-cdk.ImageOptimizationLambda.property.connections"></a>

```typescript
public readonly connections: Connections;
```

- *Type:* aws-cdk-lib.aws_ec2.Connections

Access the Connections object.

Will fail if not a VPC-enabled Lambda Function

---

##### `functionArn`<sup>Required</sup> <a name="functionArn" id="open-next-cdk.ImageOptimizationLambda.property.functionArn"></a>

```typescript
public readonly functionArn: string;
```

- *Type:* string

ARN of this function.

---

##### `functionName`<sup>Required</sup> <a name="functionName" id="open-next-cdk.ImageOptimizationLambda.property.functionName"></a>

```typescript
public readonly functionName: string;
```

- *Type:* string

Name of this function.

---

##### `grantPrincipal`<sup>Required</sup> <a name="grantPrincipal" id="open-next-cdk.ImageOptimizationLambda.property.grantPrincipal"></a>

```typescript
public readonly grantPrincipal: IPrincipal;
```

- *Type:* aws-cdk-lib.aws_iam.IPrincipal

The principal this Lambda Function is running as.

---

##### `isBoundToVpc`<sup>Required</sup> <a name="isBoundToVpc" id="open-next-cdk.ImageOptimizationLambda.property.isBoundToVpc"></a>

```typescript
public readonly isBoundToVpc: boolean;
```

- *Type:* boolean

Whether or not this Lambda function was bound to a VPC.

If this is is `false`, trying to access the `connections` object will fail.

---

##### `latestVersion`<sup>Required</sup> <a name="latestVersion" id="open-next-cdk.ImageOptimizationLambda.property.latestVersion"></a>

```typescript
public readonly latestVersion: IVersion;
```

- *Type:* aws-cdk-lib.aws_lambda.IVersion

The `$LATEST` version of this function.

Note that this is reference to a non-specific AWS Lambda version, which
means the function this version refers to can return different results in
different invocations.

To obtain a reference to an explicit version which references the current
function configuration, use `lambdaFunction.currentVersion` instead.

---

##### `permissionsNode`<sup>Required</sup> <a name="permissionsNode" id="open-next-cdk.ImageOptimizationLambda.property.permissionsNode"></a>

```typescript
public readonly permissionsNode: Node;
```

- *Type:* constructs.Node

The construct node where permissions are attached.

---

##### `resourceArnsForGrantInvoke`<sup>Required</sup> <a name="resourceArnsForGrantInvoke" id="open-next-cdk.ImageOptimizationLambda.property.resourceArnsForGrantInvoke"></a>

```typescript
public readonly resourceArnsForGrantInvoke: string[];
```

- *Type:* string[]

The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().

---

##### `role`<sup>Optional</sup> <a name="role" id="open-next-cdk.ImageOptimizationLambda.property.role"></a>

```typescript
public readonly role: IRole;
```

- *Type:* aws-cdk-lib.aws_iam.IRole

Execution role associated with this function.

---

##### `currentVersion`<sup>Required</sup> <a name="currentVersion" id="open-next-cdk.ImageOptimizationLambda.property.currentVersion"></a>

```typescript
public readonly currentVersion: Version;
```

- *Type:* aws-cdk-lib.aws_lambda.Version

Returns a `lambda.Version` which represents the current version of this Lambda function. A new version will be created every time the function's configuration changes.

You can specify options for this version using the `currentVersionOptions`
prop when initializing the `lambda.Function`.

---

##### `logGroup`<sup>Required</sup> <a name="logGroup" id="open-next-cdk.ImageOptimizationLambda.property.logGroup"></a>

```typescript
public readonly logGroup: ILogGroup;
```

- *Type:* aws-cdk-lib.aws_logs.ILogGroup

The LogGroup where the Lambda function's logs are made available.

If either `logRetention` is set or this property is called, a CloudFormation custom resource is added to the stack that
pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the correct log retention
period (never expire, by default).

Further, if the log group already exists and the `logRetention` is not set, the custom resource will reset the log retention
to never expire even if it was configured with a different value.

---

##### `runtime`<sup>Required</sup> <a name="runtime" id="open-next-cdk.ImageOptimizationLambda.property.runtime"></a>

```typescript
public readonly runtime: Runtime;
```

- *Type:* aws-cdk-lib.aws_lambda.Runtime

The runtime configured for this lambda.

---

##### `deadLetterQueue`<sup>Optional</sup> <a name="deadLetterQueue" id="open-next-cdk.ImageOptimizationLambda.property.deadLetterQueue"></a>

```typescript
public readonly deadLetterQueue: IQueue;
```

- *Type:* aws-cdk-lib.aws_sqs.IQueue

The DLQ (as queue) associated with this Lambda Function (this is an optional attribute).

---

##### `deadLetterTopic`<sup>Optional</sup> <a name="deadLetterTopic" id="open-next-cdk.ImageOptimizationLambda.property.deadLetterTopic"></a>

```typescript
public readonly deadLetterTopic: ITopic;
```

- *Type:* aws-cdk-lib.aws_sns.ITopic

The DLQ (as topic) associated with this Lambda Function (this is an optional attribute).

---

##### `timeout`<sup>Optional</sup> <a name="timeout" id="open-next-cdk.ImageOptimizationLambda.property.timeout"></a>

```typescript
public readonly timeout: Duration;
```

- *Type:* aws-cdk-lib.Duration

The timeout configured for this lambda.

---

##### `bucket`<sup>Required</sup> <a name="bucket" id="open-next-cdk.ImageOptimizationLambda.property.bucket"></a>

```typescript
public readonly bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

---


### Nextjs <a name="Nextjs" id="open-next-cdk.Nextjs"></a>

The `Nextjs` construct is a higher level construct that makes it easy to create a NextJS app.

Your standalone server application will be bundled using o(utput tracing and will be deployed to a Lambda function.
Static assets will be deployed to an S3 bucket and served via CloudFront.
You must use Next.js 10.3.0 or newer.

Please provide a `nextjsPath` to the Next.js app inside your project.

*Example*

```typescript
new Nextjs(this, "Web", {
  nextjsPath: path.resolve("packages/web"),
})
```


#### Initializers <a name="Initializers" id="open-next-cdk.Nextjs.Initializer"></a>

```typescript
import { Nextjs } from 'open-next-cdk'

new Nextjs(scope: Construct, id: string, props: NextjsProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.Nextjs.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.Nextjs.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.Nextjs.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.NextjsProps">NextjsProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.Nextjs.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.Nextjs.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.Nextjs.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.NextjsProps">NextjsProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.Nextjs.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="open-next-cdk.Nextjs.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.Nextjs.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.Nextjs.isConstruct"></a>

```typescript
import { Nextjs } from 'open-next-cdk'

Nextjs.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.Nextjs.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.Nextjs.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.Nextjs.property.bucket">bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | *No description.* |
| <code><a href="#open-next-cdk.Nextjs.property.url">url</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.Nextjs.property.assetsDeployment">assetsDeployment</a></code> | <code><a href="#open-next-cdk.NextJsAssetsDeployment">NextJsAssetsDeployment</a></code> | Asset deployment to S3. |
| <code><a href="#open-next-cdk.Nextjs.property.distribution">distribution</a></code> | <code><a href="#open-next-cdk.NextjsDistribution">NextjsDistribution</a></code> | CloudFront distribution. |
| <code><a href="#open-next-cdk.Nextjs.property.imageOptimizationFunction">imageOptimizationFunction</a></code> | <code><a href="#open-next-cdk.ImageOptimizationLambda">ImageOptimizationLambda</a></code> | The image optimization handler lambda function. |
| <code><a href="#open-next-cdk.Nextjs.property.imageOptimizationLambdaFunctionUrl">imageOptimizationLambdaFunctionUrl</a></code> | <code>aws-cdk-lib.aws_lambda.FunctionUrl</code> | *No description.* |
| <code><a href="#open-next-cdk.Nextjs.property.lambdaFunctionUrl">lambdaFunctionUrl</a></code> | <code>aws-cdk-lib.aws_lambda.FunctionUrl</code> | *No description.* |
| <code><a href="#open-next-cdk.Nextjs.property.nextBuild">nextBuild</a></code> | <code><a href="#open-next-cdk.NextjsBuild">NextjsBuild</a></code> | Built NextJS project output. |
| <code><a href="#open-next-cdk.Nextjs.property.serverFunction">serverFunction</a></code> | <code><a href="#open-next-cdk.NextJsLambda">NextJsLambda</a></code> | The main NextJS server handler lambda function. |
| <code><a href="#open-next-cdk.Nextjs.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Where build-time assets for deployment are stored. |
| <code><a href="#open-next-cdk.Nextjs.property.configBucket">configBucket</a></code> | <code>aws-cdk-lib.aws_s3.Bucket</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.Nextjs.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `bucket`<sup>Required</sup> <a name="bucket" id="open-next-cdk.Nextjs.property.bucket"></a>

```typescript
public readonly bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

---

##### `url`<sup>Required</sup> <a name="url" id="open-next-cdk.Nextjs.property.url"></a>

```typescript
public readonly url: string;
```

- *Type:* string

---

##### `assetsDeployment`<sup>Required</sup> <a name="assetsDeployment" id="open-next-cdk.Nextjs.property.assetsDeployment"></a>

```typescript
public readonly assetsDeployment: NextJsAssetsDeployment;
```

- *Type:* <a href="#open-next-cdk.NextJsAssetsDeployment">NextJsAssetsDeployment</a>

Asset deployment to S3.

---

##### `distribution`<sup>Required</sup> <a name="distribution" id="open-next-cdk.Nextjs.property.distribution"></a>

```typescript
public readonly distribution: NextjsDistribution;
```

- *Type:* <a href="#open-next-cdk.NextjsDistribution">NextjsDistribution</a>

CloudFront distribution.

---

##### `imageOptimizationFunction`<sup>Required</sup> <a name="imageOptimizationFunction" id="open-next-cdk.Nextjs.property.imageOptimizationFunction"></a>

```typescript
public readonly imageOptimizationFunction: ImageOptimizationLambda;
```

- *Type:* <a href="#open-next-cdk.ImageOptimizationLambda">ImageOptimizationLambda</a>

The image optimization handler lambda function.

---

##### `imageOptimizationLambdaFunctionUrl`<sup>Required</sup> <a name="imageOptimizationLambdaFunctionUrl" id="open-next-cdk.Nextjs.property.imageOptimizationLambdaFunctionUrl"></a>

```typescript
public readonly imageOptimizationLambdaFunctionUrl: FunctionUrl;
```

- *Type:* aws-cdk-lib.aws_lambda.FunctionUrl

---

##### `lambdaFunctionUrl`<sup>Required</sup> <a name="lambdaFunctionUrl" id="open-next-cdk.Nextjs.property.lambdaFunctionUrl"></a>

```typescript
public readonly lambdaFunctionUrl: FunctionUrl;
```

- *Type:* aws-cdk-lib.aws_lambda.FunctionUrl

---

##### `nextBuild`<sup>Required</sup> <a name="nextBuild" id="open-next-cdk.Nextjs.property.nextBuild"></a>

```typescript
public readonly nextBuild: NextjsBuild;
```

- *Type:* <a href="#open-next-cdk.NextjsBuild">NextjsBuild</a>

Built NextJS project output.

---

##### `serverFunction`<sup>Required</sup> <a name="serverFunction" id="open-next-cdk.Nextjs.property.serverFunction"></a>

```typescript
public readonly serverFunction: NextJsLambda;
```

- *Type:* <a href="#open-next-cdk.NextJsLambda">NextJsLambda</a>

The main NextJS server handler lambda function.

---

##### `tempBuildDir`<sup>Required</sup> <a name="tempBuildDir" id="open-next-cdk.Nextjs.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Where build-time assets for deployment are stored.

---

##### `configBucket`<sup>Optional</sup> <a name="configBucket" id="open-next-cdk.Nextjs.property.configBucket"></a>

```typescript
public readonly configBucket: Bucket;
```

- *Type:* aws-cdk-lib.aws_s3.Bucket

---


### NextJsAssetsDeployment <a name="NextJsAssetsDeployment" id="open-next-cdk.NextJsAssetsDeployment"></a>

Uploads NextJS-built static and public files to S3.

Will rewrite CloudFormation references with their resolved values after uploading.

#### Initializers <a name="Initializers" id="open-next-cdk.NextJsAssetsDeployment.Initializer"></a>

```typescript
import { NextJsAssetsDeployment } from 'open-next-cdk'

new NextJsAssetsDeployment(scope: Construct, id: string, props: NextjsAssetsDeploymentProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps">NextjsAssetsDeploymentProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextJsAssetsDeployment.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextJsAssetsDeployment.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.NextJsAssetsDeployment.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.NextjsAssetsDeploymentProps">NextjsAssetsDeploymentProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="open-next-cdk.NextJsAssetsDeployment.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.NextJsAssetsDeployment.isConstruct"></a>

```typescript
import { NextJsAssetsDeployment } from 'open-next-cdk'

NextJsAssetsDeployment.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.NextJsAssetsDeployment.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.property.bucket">bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | Bucket containing assets. |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.property.deployments">deployments</a></code> | <code>aws-cdk-lib.aws_s3_deployment.BucketDeployment[]</code> | Asset deployments to S3. |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.property.staticTempDir">staticTempDir</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextJsAssetsDeployment.property.rewriter">rewriter</a></code> | <code><a href="#open-next-cdk.NextjsS3EnvRewriter">NextjsS3EnvRewriter</a></code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.NextJsAssetsDeployment.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `bucket`<sup>Required</sup> <a name="bucket" id="open-next-cdk.NextJsAssetsDeployment.property.bucket"></a>

```typescript
public readonly bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

Bucket containing assets.

---

##### `deployments`<sup>Required</sup> <a name="deployments" id="open-next-cdk.NextJsAssetsDeployment.property.deployments"></a>

```typescript
public readonly deployments: BucketDeployment[];
```

- *Type:* aws-cdk-lib.aws_s3_deployment.BucketDeployment[]

Asset deployments to S3.

---

##### `staticTempDir`<sup>Required</sup> <a name="staticTempDir" id="open-next-cdk.NextJsAssetsDeployment.property.staticTempDir"></a>

```typescript
public readonly staticTempDir: string;
```

- *Type:* string

---

##### `rewriter`<sup>Optional</sup> <a name="rewriter" id="open-next-cdk.NextJsAssetsDeployment.property.rewriter"></a>

```typescript
public readonly rewriter: NextjsS3EnvRewriter;
```

- *Type:* <a href="#open-next-cdk.NextjsS3EnvRewriter">NextjsS3EnvRewriter</a>

---


### NextjsBuild <a name="NextjsBuild" id="open-next-cdk.NextjsBuild"></a>

Represents a built NextJS application.

This construct runs `npm build` in standalone output mode inside your `nextjsPath`.
This construct can be used by higher level constructs or used directly.

#### Initializers <a name="Initializers" id="open-next-cdk.NextjsBuild.Initializer"></a>

```typescript
import { NextjsBuild } from 'open-next-cdk'

new NextjsBuild(scope: Construct, id: string, props: NextjsBuildProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsBuild.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsBuild.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsBuild.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.NextjsBuildProps">NextjsBuildProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextjsBuild.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextjsBuild.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.NextjsBuild.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.NextjsBuildProps">NextjsBuildProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsBuild.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#open-next-cdk.NextjsBuild.readPublicFileList">readPublicFileList</a></code> | *No description.* |

---

##### `toString` <a name="toString" id="open-next-cdk.NextjsBuild.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `readPublicFileList` <a name="readPublicFileList" id="open-next-cdk.NextjsBuild.readPublicFileList"></a>

```typescript
public readPublicFileList(): string[]
```

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsBuild.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.NextjsBuild.isConstruct"></a>

```typescript
import { NextjsBuild } from 'open-next-cdk'

NextjsBuild.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.NextjsBuild.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsBuild.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.NextjsBuild.property.nextImageFnDir">nextImageFnDir</a></code> | <code>string</code> | Contains function for processessing image requests. |
| <code><a href="#open-next-cdk.NextjsBuild.property.nextServerFnDir">nextServerFnDir</a></code> | <code>string</code> | Contains server code and dependencies. |
| <code><a href="#open-next-cdk.NextjsBuild.property.nextStaticDir">nextStaticDir</a></code> | <code>string</code> | Static files containing client-side code. |
| <code><a href="#open-next-cdk.NextjsBuild.property.projectRoot">projectRoot</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsBuild.property.props">props</a></code> | <code><a href="#open-next-cdk.NextjsBuildProps">NextjsBuildProps</a></code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsBuild.property.nextMiddlewareFnDir">nextMiddlewareFnDir</a></code> | <code>string</code> | Contains code for middleware. |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.NextjsBuild.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `nextImageFnDir`<sup>Required</sup> <a name="nextImageFnDir" id="open-next-cdk.NextjsBuild.property.nextImageFnDir"></a>

```typescript
public readonly nextImageFnDir: string;
```

- *Type:* string

Contains function for processessing image requests.

Should be arm64.

---

##### `nextServerFnDir`<sup>Required</sup> <a name="nextServerFnDir" id="open-next-cdk.NextjsBuild.property.nextServerFnDir"></a>

```typescript
public readonly nextServerFnDir: string;
```

- *Type:* string

Contains server code and dependencies.

---

##### `nextStaticDir`<sup>Required</sup> <a name="nextStaticDir" id="open-next-cdk.NextjsBuild.property.nextStaticDir"></a>

```typescript
public readonly nextStaticDir: string;
```

- *Type:* string

Static files containing client-side code.

---

##### `projectRoot`<sup>Required</sup> <a name="projectRoot" id="open-next-cdk.NextjsBuild.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.NextjsBuild.property.props"></a>

```typescript
public readonly props: NextjsBuildProps;
```

- *Type:* <a href="#open-next-cdk.NextjsBuildProps">NextjsBuildProps</a>

---

##### `nextMiddlewareFnDir`<sup>Optional</sup> <a name="nextMiddlewareFnDir" id="open-next-cdk.NextjsBuild.property.nextMiddlewareFnDir"></a>

```typescript
public readonly nextMiddlewareFnDir: string;
```

- *Type:* string

Contains code for middleware.

Not currently used.

---


### NextjsDistribution <a name="NextjsDistribution" id="open-next-cdk.NextjsDistribution"></a>

Create a CloudFront distribution to serve a Next.js application.

#### Initializers <a name="Initializers" id="open-next-cdk.NextjsDistribution.Initializer"></a>

```typescript
import { NextjsDistribution } from 'open-next-cdk'

new NextjsDistribution(scope: Construct, id: string, props: NextjsDistributionProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsDistribution.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsDistribution.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsDistribution.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.NextjsDistributionProps">NextjsDistributionProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextjsDistribution.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextjsDistribution.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.NextjsDistribution.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.NextjsDistributionProps">NextjsDistributionProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsDistribution.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="open-next-cdk.NextjsDistribution.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsDistribution.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.NextjsDistribution.isConstruct"></a>

```typescript
import { NextjsDistribution } from 'open-next-cdk'

NextjsDistribution.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.NextjsDistribution.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsDistribution.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.fallbackOriginRequestPolicyProps">fallbackOriginRequestPolicyProps</a></code> | <code>aws-cdk-lib.aws_cloudfront.OriginRequestPolicyProps</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsDistribution.property.imageCachePolicyProps">imageCachePolicyProps</a></code> | <code>aws-cdk-lib.aws_cloudfront.CachePolicyProps</code> | The default CloudFront cache policy properties for images. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.imageOptimizationOriginRequestPolicyProps">imageOptimizationOriginRequestPolicyProps</a></code> | <code>aws-cdk-lib.aws_cloudfront.OriginRequestPolicyProps</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsDistribution.property.lambdaCachePolicyProps">lambdaCachePolicyProps</a></code> | <code>aws-cdk-lib.aws_cloudfront.CachePolicyProps</code> | The default CloudFront cache policy properties for the Lambda server handler. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.lambdaOriginRequestPolicyProps">lambdaOriginRequestPolicyProps</a></code> | <code>aws-cdk-lib.aws_cloudfront.OriginRequestPolicyProps</code> | The default CloudFront lambda origin request policy. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.staticCachePolicyProps">staticCachePolicyProps</a></code> | <code>aws-cdk-lib.aws_cloudfront.CachePolicyProps</code> | The default CloudFront cache policy properties for static pages. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.distributionDomain">distributionDomain</a></code> | <code>string</code> | The domain name of the internally created CloudFront Distribution. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.distributionId">distributionId</a></code> | <code>string</code> | The ID of the internally created CloudFront Distribution. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.url">url</a></code> | <code>string</code> | The CloudFront URL of the website. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.customDomainName">customDomainName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsDistribution.property.customDomainUrl">customDomainUrl</a></code> | <code>string</code> | If the custom domain is enabled, this is the URL of the website with the custom domain. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.distribution">distribution</a></code> | <code>aws-cdk-lib.aws_cloudfront.Distribution</code> | The internally created CloudFront `Distribution` instance. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsDistribution.property.certificate">certificate</a></code> | <code>aws-cdk-lib.aws_certificatemanager.ICertificate</code> | The AWS Certificate Manager certificate for the custom domain. |
| <code><a href="#open-next-cdk.NextjsDistribution.property.hostedZone">hostedZone</a></code> | <code>aws-cdk-lib.aws_route53.IHostedZone</code> | The Route 53 hosted zone for the custom domain. |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.NextjsDistribution.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `fallbackOriginRequestPolicyProps`<sup>Required</sup> <a name="fallbackOriginRequestPolicyProps" id="open-next-cdk.NextjsDistribution.property.fallbackOriginRequestPolicyProps"></a>

```typescript
public readonly fallbackOriginRequestPolicyProps: OriginRequestPolicyProps;
```

- *Type:* aws-cdk-lib.aws_cloudfront.OriginRequestPolicyProps

---

##### `imageCachePolicyProps`<sup>Required</sup> <a name="imageCachePolicyProps" id="open-next-cdk.NextjsDistribution.property.imageCachePolicyProps"></a>

```typescript
public readonly imageCachePolicyProps: CachePolicyProps;
```

- *Type:* aws-cdk-lib.aws_cloudfront.CachePolicyProps

The default CloudFront cache policy properties for images.

---

##### `imageOptimizationOriginRequestPolicyProps`<sup>Required</sup> <a name="imageOptimizationOriginRequestPolicyProps" id="open-next-cdk.NextjsDistribution.property.imageOptimizationOriginRequestPolicyProps"></a>

```typescript
public readonly imageOptimizationOriginRequestPolicyProps: OriginRequestPolicyProps;
```

- *Type:* aws-cdk-lib.aws_cloudfront.OriginRequestPolicyProps

---

##### `lambdaCachePolicyProps`<sup>Required</sup> <a name="lambdaCachePolicyProps" id="open-next-cdk.NextjsDistribution.property.lambdaCachePolicyProps"></a>

```typescript
public readonly lambdaCachePolicyProps: CachePolicyProps;
```

- *Type:* aws-cdk-lib.aws_cloudfront.CachePolicyProps

The default CloudFront cache policy properties for the Lambda server handler.

---

##### `lambdaOriginRequestPolicyProps`<sup>Required</sup> <a name="lambdaOriginRequestPolicyProps" id="open-next-cdk.NextjsDistribution.property.lambdaOriginRequestPolicyProps"></a>

```typescript
public readonly lambdaOriginRequestPolicyProps: OriginRequestPolicyProps;
```

- *Type:* aws-cdk-lib.aws_cloudfront.OriginRequestPolicyProps

The default CloudFront lambda origin request policy.

---

##### `staticCachePolicyProps`<sup>Required</sup> <a name="staticCachePolicyProps" id="open-next-cdk.NextjsDistribution.property.staticCachePolicyProps"></a>

```typescript
public readonly staticCachePolicyProps: CachePolicyProps;
```

- *Type:* aws-cdk-lib.aws_cloudfront.CachePolicyProps

The default CloudFront cache policy properties for static pages.

---

##### `distributionDomain`<sup>Required</sup> <a name="distributionDomain" id="open-next-cdk.NextjsDistribution.property.distributionDomain"></a>

```typescript
public readonly distributionDomain: string;
```

- *Type:* string

The domain name of the internally created CloudFront Distribution.

---

##### `distributionId`<sup>Required</sup> <a name="distributionId" id="open-next-cdk.NextjsDistribution.property.distributionId"></a>

```typescript
public readonly distributionId: string;
```

- *Type:* string

The ID of the internally created CloudFront Distribution.

---

##### `url`<sup>Required</sup> <a name="url" id="open-next-cdk.NextjsDistribution.property.url"></a>

```typescript
public readonly url: string;
```

- *Type:* string

The CloudFront URL of the website.

---

##### `customDomainName`<sup>Optional</sup> <a name="customDomainName" id="open-next-cdk.NextjsDistribution.property.customDomainName"></a>

```typescript
public readonly customDomainName: string;
```

- *Type:* string

---

##### `customDomainUrl`<sup>Optional</sup> <a name="customDomainUrl" id="open-next-cdk.NextjsDistribution.property.customDomainUrl"></a>

```typescript
public readonly customDomainUrl: string;
```

- *Type:* string

If the custom domain is enabled, this is the URL of the website with the custom domain.

---

##### `distribution`<sup>Required</sup> <a name="distribution" id="open-next-cdk.NextjsDistribution.property.distribution"></a>

```typescript
public readonly distribution: Distribution;
```

- *Type:* aws-cdk-lib.aws_cloudfront.Distribution

The internally created CloudFront `Distribution` instance.

---

##### `tempBuildDir`<sup>Required</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsDistribution.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

---

##### `certificate`<sup>Optional</sup> <a name="certificate" id="open-next-cdk.NextjsDistribution.property.certificate"></a>

```typescript
public readonly certificate: ICertificate;
```

- *Type:* aws-cdk-lib.aws_certificatemanager.ICertificate

The AWS Certificate Manager certificate for the custom domain.

---

##### `hostedZone`<sup>Optional</sup> <a name="hostedZone" id="open-next-cdk.NextjsDistribution.property.hostedZone"></a>

```typescript
public readonly hostedZone: IHostedZone;
```

- *Type:* aws-cdk-lib.aws_route53.IHostedZone

The Route 53 hosted zone for the custom domain.

---


### NextJsLambda <a name="NextJsLambda" id="open-next-cdk.NextJsLambda"></a>

Build a lambda function from a NextJS application to handle server-side rendering, API routes, and image optimization.

#### Initializers <a name="Initializers" id="open-next-cdk.NextJsLambda.Initializer"></a>

```typescript
import { NextJsLambda } from 'open-next-cdk'

new NextJsLambda(scope: Construct, id: string, props: NextjsLambdaProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextJsLambda.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.NextJsLambda.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextJsLambda.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.NextjsLambdaProps">NextjsLambdaProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextJsLambda.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextJsLambda.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.NextJsLambda.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.NextjsLambdaProps">NextjsLambdaProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextJsLambda.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="open-next-cdk.NextJsLambda.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextJsLambda.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.NextJsLambda.isConstruct"></a>

```typescript
import { NextJsLambda } from 'open-next-cdk'

NextJsLambda.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.NextJsLambda.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextJsLambda.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.NextJsLambda.property.lambdaFunction">lambdaFunction</a></code> | <code>aws-cdk-lib.aws_lambda.Function</code> | *No description.* |
| <code><a href="#open-next-cdk.NextJsLambda.property.configBucket">configBucket</a></code> | <code>aws-cdk-lib.aws_s3.Bucket</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.NextJsLambda.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `lambdaFunction`<sup>Required</sup> <a name="lambdaFunction" id="open-next-cdk.NextJsLambda.property.lambdaFunction"></a>

```typescript
public readonly lambdaFunction: Function;
```

- *Type:* aws-cdk-lib.aws_lambda.Function

---

##### `configBucket`<sup>Optional</sup> <a name="configBucket" id="open-next-cdk.NextJsLambda.property.configBucket"></a>

```typescript
public readonly configBucket: Bucket;
```

- *Type:* aws-cdk-lib.aws_s3.Bucket

---


### NextjsLayer <a name="NextjsLayer" id="open-next-cdk.NextjsLayer"></a>

Lambda layer for Next.js. Contains Sharp 0.30.0.

#### Initializers <a name="Initializers" id="open-next-cdk.NextjsLayer.Initializer"></a>

```typescript
import { NextjsLayer } from 'open-next-cdk'

new NextjsLayer(scope: Construct, id: string, props: NextjsLayerProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsLayer.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsLayer.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsLayer.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.NextjsLayerProps">NextjsLayerProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextjsLayer.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextjsLayer.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.NextjsLayer.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.NextjsLayerProps">NextjsLayerProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsLayer.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#open-next-cdk.NextjsLayer.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |
| <code><a href="#open-next-cdk.NextjsLayer.addPermission">addPermission</a></code> | Add permission for this layer version to specific entities. |

---

##### `toString` <a name="toString" id="open-next-cdk.NextjsLayer.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="open-next-cdk.NextjsLayer.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy: RemovalPolicy): void
```

Apply the given removal policy to this resource.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).

###### `policy`<sup>Required</sup> <a name="policy" id="open-next-cdk.NextjsLayer.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

##### `addPermission` <a name="addPermission" id="open-next-cdk.NextjsLayer.addPermission"></a>

```typescript
public addPermission(id: string, permission: LayerVersionPermission): void
```

Add permission for this layer version to specific entities.

Usage within
the same account where the layer is defined is always allowed and does not
require calling this method. Note that the principal that creates the
Lambda function using the layer (for example, a CloudFormation changeset
execution role) also needs to have the ``lambda:GetLayerVersion``
permission on the layer version.

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextjsLayer.addPermission.parameter.id"></a>

- *Type:* string

---

###### `permission`<sup>Required</sup> <a name="permission" id="open-next-cdk.NextjsLayer.addPermission.parameter.permission"></a>

- *Type:* aws-cdk-lib.aws_lambda.LayerVersionPermission

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsLayer.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#open-next-cdk.NextjsLayer.isOwnedResource">isOwnedResource</a></code> | Returns true if the construct was created by CDK, and false otherwise. |
| <code><a href="#open-next-cdk.NextjsLayer.isResource">isResource</a></code> | Check whether the given construct is a Resource. |
| <code><a href="#open-next-cdk.NextjsLayer.fromLayerVersionArn">fromLayerVersionArn</a></code> | Imports a layer version by ARN. |
| <code><a href="#open-next-cdk.NextjsLayer.fromLayerVersionAttributes">fromLayerVersionAttributes</a></code> | Imports a Layer that has been defined externally. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.NextjsLayer.isConstruct"></a>

```typescript
import { NextjsLayer } from 'open-next-cdk'

NextjsLayer.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.NextjsLayer.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isOwnedResource` <a name="isOwnedResource" id="open-next-cdk.NextjsLayer.isOwnedResource"></a>

```typescript
import { NextjsLayer } from 'open-next-cdk'

NextjsLayer.isOwnedResource(construct: IConstruct)
```

Returns true if the construct was created by CDK, and false otherwise.

###### `construct`<sup>Required</sup> <a name="construct" id="open-next-cdk.NextjsLayer.isOwnedResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

##### `isResource` <a name="isResource" id="open-next-cdk.NextjsLayer.isResource"></a>

```typescript
import { NextjsLayer } from 'open-next-cdk'

NextjsLayer.isResource(construct: IConstruct)
```

Check whether the given construct is a Resource.

###### `construct`<sup>Required</sup> <a name="construct" id="open-next-cdk.NextjsLayer.isResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

##### `fromLayerVersionArn` <a name="fromLayerVersionArn" id="open-next-cdk.NextjsLayer.fromLayerVersionArn"></a>

```typescript
import { NextjsLayer } from 'open-next-cdk'

NextjsLayer.fromLayerVersionArn(scope: Construct, id: string, layerVersionArn: string)
```

Imports a layer version by ARN.

Assumes it is compatible with all Lambda runtimes.

###### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextjsLayer.fromLayerVersionArn.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextjsLayer.fromLayerVersionArn.parameter.id"></a>

- *Type:* string

---

###### `layerVersionArn`<sup>Required</sup> <a name="layerVersionArn" id="open-next-cdk.NextjsLayer.fromLayerVersionArn.parameter.layerVersionArn"></a>

- *Type:* string

---

##### `fromLayerVersionAttributes` <a name="fromLayerVersionAttributes" id="open-next-cdk.NextjsLayer.fromLayerVersionAttributes"></a>

```typescript
import { NextjsLayer } from 'open-next-cdk'

NextjsLayer.fromLayerVersionAttributes(scope: Construct, id: string, attrs: LayerVersionAttributes)
```

Imports a Layer that has been defined externally.

###### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextjsLayer.fromLayerVersionAttributes.parameter.scope"></a>

- *Type:* constructs.Construct

the parent Construct that will use the imported layer.

---

###### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextjsLayer.fromLayerVersionAttributes.parameter.id"></a>

- *Type:* string

the id of the imported layer in the construct tree.

---

###### `attrs`<sup>Required</sup> <a name="attrs" id="open-next-cdk.NextjsLayer.fromLayerVersionAttributes.parameter.attrs"></a>

- *Type:* aws-cdk-lib.aws_lambda.LayerVersionAttributes

the properties of the imported layer.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsLayer.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.NextjsLayer.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#open-next-cdk.NextjsLayer.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this resource is defined. |
| <code><a href="#open-next-cdk.NextjsLayer.property.layerVersionArn">layerVersionArn</a></code> | <code>string</code> | The ARN of the Lambda Layer version that this Layer defines. |
| <code><a href="#open-next-cdk.NextjsLayer.property.compatibleRuntimes">compatibleRuntimes</a></code> | <code>aws-cdk-lib.aws_lambda.Runtime[]</code> | The runtimes compatible with this Layer. |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.NextjsLayer.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `env`<sup>Required</sup> <a name="env" id="open-next-cdk.NextjsLayer.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* aws-cdk-lib.ResourceEnvironment

The environment this resource belongs to.

For resources that are created and managed by the CDK
(generally, those created by creating new class instances like Role, Bucket, etc.),
this is always the same as the environment of the stack they belong to;
however, for imported resources
(those obtained from static methods like fromRoleArn, fromBucketName, etc.),
that might be different than the stack they were imported into.

---

##### `stack`<sup>Required</sup> <a name="stack" id="open-next-cdk.NextjsLayer.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this resource is defined.

---

##### `layerVersionArn`<sup>Required</sup> <a name="layerVersionArn" id="open-next-cdk.NextjsLayer.property.layerVersionArn"></a>

```typescript
public readonly layerVersionArn: string;
```

- *Type:* string

The ARN of the Lambda Layer version that this Layer defines.

---

##### `compatibleRuntimes`<sup>Optional</sup> <a name="compatibleRuntimes" id="open-next-cdk.NextjsLayer.property.compatibleRuntimes"></a>

```typescript
public readonly compatibleRuntimes: Runtime[];
```

- *Type:* aws-cdk-lib.aws_lambda.Runtime[]

The runtimes compatible with this Layer.

---


### NextjsS3EnvRewriter <a name="NextjsS3EnvRewriter" id="open-next-cdk.NextjsS3EnvRewriter"></a>

Rewrites variables in S3 objects after a deployment happens to replace CloudFormation tokens with their values.

These values are not resolved at build time because they are
only known at deploy time.

#### Initializers <a name="Initializers" id="open-next-cdk.NextjsS3EnvRewriter.Initializer"></a>

```typescript
import { NextjsS3EnvRewriter } from 'open-next-cdk'

new NextjsS3EnvRewriter(scope: Construct, id: string, props: NextjsS3EnvRewriterProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriter.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriter.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriter.Initializer.parameter.props">props</a></code> | <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps">NextjsS3EnvRewriterProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="open-next-cdk.NextjsS3EnvRewriter.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="open-next-cdk.NextjsS3EnvRewriter.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="open-next-cdk.NextjsS3EnvRewriter.Initializer.parameter.props"></a>

- *Type:* <a href="#open-next-cdk.NextjsS3EnvRewriterProps">NextjsS3EnvRewriterProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriter.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="open-next-cdk.NextjsS3EnvRewriter.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriter.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="open-next-cdk.NextjsS3EnvRewriter.isConstruct"></a>

```typescript
import { NextjsS3EnvRewriter } from 'open-next-cdk'

NextjsS3EnvRewriter.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="open-next-cdk.NextjsS3EnvRewriter.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriter.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriter.property.rewriteNode">rewriteNode</a></code> | <code>constructs.Construct</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="open-next-cdk.NextjsS3EnvRewriter.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `rewriteNode`<sup>Optional</sup> <a name="rewriteNode" id="open-next-cdk.NextjsS3EnvRewriter.property.rewriteNode"></a>

```typescript
public readonly rewriteNode: Construct;
```

- *Type:* constructs.Construct

---


## Structs <a name="Structs" id="Structs"></a>

### BaseSiteDomainProps <a name="BaseSiteDomainProps" id="open-next-cdk.BaseSiteDomainProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.BaseSiteDomainProps.Initializer"></a>

```typescript
import { BaseSiteDomainProps } from 'open-next-cdk'

const baseSiteDomainProps: BaseSiteDomainProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.BaseSiteDomainProps.property.domainName">domainName</a></code> | <code>string</code> | The domain to be assigned to the website URL (ie. domain.com). |
| <code><a href="#open-next-cdk.BaseSiteDomainProps.property.alternateNames">alternateNames</a></code> | <code>string[]</code> | Specify additional names that should route to the Cloudfront Distribution. |
| <code><a href="#open-next-cdk.BaseSiteDomainProps.property.certificate">certificate</a></code> | <code>aws-cdk-lib.aws_certificatemanager.ICertificate</code> | Import the certificate for the domain. |
| <code><a href="#open-next-cdk.BaseSiteDomainProps.property.domainAlias">domainAlias</a></code> | <code>string</code> | An alternative domain to be assigned to the website URL. |
| <code><a href="#open-next-cdk.BaseSiteDomainProps.property.hostedZone">hostedZone</a></code> | <code>aws-cdk-lib.aws_route53.IHostedZone</code> | Import the underlying Route 53 hosted zone. |
| <code><a href="#open-next-cdk.BaseSiteDomainProps.property.isExternalDomain">isExternalDomain</a></code> | <code>boolean</code> | Set this option if the domain is not hosted on Amazon Route 53. |

---

##### `domainName`<sup>Required</sup> <a name="domainName" id="open-next-cdk.BaseSiteDomainProps.property.domainName"></a>

```typescript
public readonly domainName: string;
```

- *Type:* string

The domain to be assigned to the website URL (ie. domain.com).

Supports domains that are hosted either on [Route 53](https://aws.amazon.com/route53/) or externally.

---

##### `alternateNames`<sup>Optional</sup> <a name="alternateNames" id="open-next-cdk.BaseSiteDomainProps.property.alternateNames"></a>

```typescript
public readonly alternateNames: string[];
```

- *Type:* string[]

Specify additional names that should route to the Cloudfront Distribution.

Note, certificates for these names will not be automatically generated so the `certificate` option must be specified.

---

##### `certificate`<sup>Optional</sup> <a name="certificate" id="open-next-cdk.BaseSiteDomainProps.property.certificate"></a>

```typescript
public readonly certificate: ICertificate;
```

- *Type:* aws-cdk-lib.aws_certificatemanager.ICertificate

Import the certificate for the domain.

By default, SST will create a certificate with the domain name. The certificate will be created in the `us-east-1`(N. Virginia) region as required by AWS CloudFront.

Set this option if you have an existing certificate in the `us-east-1` region in AWS Certificate Manager you want to use.

---

##### `domainAlias`<sup>Optional</sup> <a name="domainAlias" id="open-next-cdk.BaseSiteDomainProps.property.domainAlias"></a>

```typescript
public readonly domainAlias: string;
```

- *Type:* string

An alternative domain to be assigned to the website URL.

Visitors to the alias will be redirected to the main domain. (ie. `www.domain.com`).

Use this to create a `www.` version of your domain and redirect visitors to the root domain.

---

##### `hostedZone`<sup>Optional</sup> <a name="hostedZone" id="open-next-cdk.BaseSiteDomainProps.property.hostedZone"></a>

```typescript
public readonly hostedZone: IHostedZone;
```

- *Type:* aws-cdk-lib.aws_route53.IHostedZone

Import the underlying Route 53 hosted zone.

---

##### `isExternalDomain`<sup>Optional</sup> <a name="isExternalDomain" id="open-next-cdk.BaseSiteDomainProps.property.isExternalDomain"></a>

```typescript
public readonly isExternalDomain: boolean;
```

- *Type:* boolean

Set this option if the domain is not hosted on Amazon Route 53.

---

### BaseSiteEnvironmentOutputsInfo <a name="BaseSiteEnvironmentOutputsInfo" id="open-next-cdk.BaseSiteEnvironmentOutputsInfo"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.BaseSiteEnvironmentOutputsInfo.Initializer"></a>

```typescript
import { BaseSiteEnvironmentOutputsInfo } from 'open-next-cdk'

const baseSiteEnvironmentOutputsInfo: BaseSiteEnvironmentOutputsInfo = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.BaseSiteEnvironmentOutputsInfo.property.environmentOutputs">environmentOutputs</a></code> | <code>{[ key: string ]: string}</code> | *No description.* |
| <code><a href="#open-next-cdk.BaseSiteEnvironmentOutputsInfo.property.path">path</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.BaseSiteEnvironmentOutputsInfo.property.stack">stack</a></code> | <code>string</code> | *No description.* |

---

##### `environmentOutputs`<sup>Required</sup> <a name="environmentOutputs" id="open-next-cdk.BaseSiteEnvironmentOutputsInfo.property.environmentOutputs"></a>

```typescript
public readonly environmentOutputs: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

---

##### `path`<sup>Required</sup> <a name="path" id="open-next-cdk.BaseSiteEnvironmentOutputsInfo.property.path"></a>

```typescript
public readonly path: string;
```

- *Type:* string

---

##### `stack`<sup>Required</sup> <a name="stack" id="open-next-cdk.BaseSiteEnvironmentOutputsInfo.property.stack"></a>

```typescript
public readonly stack: string;
```

- *Type:* string

---

### BaseSiteReplaceProps <a name="BaseSiteReplaceProps" id="open-next-cdk.BaseSiteReplaceProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.BaseSiteReplaceProps.Initializer"></a>

```typescript
import { BaseSiteReplaceProps } from 'open-next-cdk'

const baseSiteReplaceProps: BaseSiteReplaceProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.BaseSiteReplaceProps.property.files">files</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.BaseSiteReplaceProps.property.replace">replace</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.BaseSiteReplaceProps.property.search">search</a></code> | <code>string</code> | *No description.* |

---

##### `files`<sup>Required</sup> <a name="files" id="open-next-cdk.BaseSiteReplaceProps.property.files"></a>

```typescript
public readonly files: string;
```

- *Type:* string

---

##### `replace`<sup>Required</sup> <a name="replace" id="open-next-cdk.BaseSiteReplaceProps.property.replace"></a>

```typescript
public readonly replace: string;
```

- *Type:* string

---

##### `search`<sup>Required</sup> <a name="search" id="open-next-cdk.BaseSiteReplaceProps.property.search"></a>

```typescript
public readonly search: string;
```

- *Type:* string

---

### CreateArchiveArgs <a name="CreateArchiveArgs" id="open-next-cdk.CreateArchiveArgs"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.CreateArchiveArgs.Initializer"></a>

```typescript
import { CreateArchiveArgs } from 'open-next-cdk'

const createArchiveArgs: CreateArchiveArgs = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.CreateArchiveArgs.property.directory">directory</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.CreateArchiveArgs.property.zipFileName">zipFileName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.CreateArchiveArgs.property.zipOutDir">zipOutDir</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.CreateArchiveArgs.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | *No description.* |
| <code><a href="#open-next-cdk.CreateArchiveArgs.property.fileGlob">fileGlob</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.CreateArchiveArgs.property.quiet">quiet</a></code> | <code>boolean</code> | *No description.* |

---

##### `directory`<sup>Required</sup> <a name="directory" id="open-next-cdk.CreateArchiveArgs.property.directory"></a>

```typescript
public readonly directory: string;
```

- *Type:* string

---

##### `zipFileName`<sup>Required</sup> <a name="zipFileName" id="open-next-cdk.CreateArchiveArgs.property.zipFileName"></a>

```typescript
public readonly zipFileName: string;
```

- *Type:* string

---

##### `zipOutDir`<sup>Required</sup> <a name="zipOutDir" id="open-next-cdk.CreateArchiveArgs.property.zipOutDir"></a>

```typescript
public readonly zipOutDir: string;
```

- *Type:* string

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.CreateArchiveArgs.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number

---

##### `fileGlob`<sup>Optional</sup> <a name="fileGlob" id="open-next-cdk.CreateArchiveArgs.property.fileGlob"></a>

```typescript
public readonly fileGlob: string;
```

- *Type:* string

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.CreateArchiveArgs.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

---

### ImageOptimizationProps <a name="ImageOptimizationProps" id="open-next-cdk.ImageOptimizationProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.ImageOptimizationProps.Initializer"></a>

```typescript
import { ImageOptimizationProps } from 'open-next-cdk'

const imageOptimizationProps: ImageOptimizationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.bucket">bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | The S3 bucket holding application images. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.nextBuild">nextBuild</a></code> | <code><a href="#open-next-cdk.NextjsBuild">NextjsBuild</a></code> | The `NextjsBuild` instance representing the built Nextjs application. |
| <code><a href="#open-next-cdk.ImageOptimizationProps.property.lambdaOptions">lambdaOptions</a></code> | <code>aws-cdk-lib.aws_lambda.FunctionOptions</code> | Override function properties. |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.ImageOptimizationProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.ImageOptimizationProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.ImageOptimizationProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.ImageOptimizationProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.ImageOptimizationProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.ImageOptimizationProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.ImageOptimizationProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.ImageOptimizationProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.ImageOptimizationProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.ImageOptimizationProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.ImageOptimizationProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

##### `bucket`<sup>Required</sup> <a name="bucket" id="open-next-cdk.ImageOptimizationProps.property.bucket"></a>

```typescript
public readonly bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

The S3 bucket holding application images.

---

##### `nextBuild`<sup>Required</sup> <a name="nextBuild" id="open-next-cdk.ImageOptimizationProps.property.nextBuild"></a>

```typescript
public readonly nextBuild: NextjsBuild;
```

- *Type:* <a href="#open-next-cdk.NextjsBuild">NextjsBuild</a>

The `NextjsBuild` instance representing the built Nextjs application.

---

##### `lambdaOptions`<sup>Optional</sup> <a name="lambdaOptions" id="open-next-cdk.ImageOptimizationProps.property.lambdaOptions"></a>

```typescript
public readonly lambdaOptions: FunctionOptions;
```

- *Type:* aws-cdk-lib.aws_lambda.FunctionOptions

Override function properties.

---

### NextjsAssetsCachePolicyProps <a name="NextjsAssetsCachePolicyProps" id="open-next-cdk.NextjsAssetsCachePolicyProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsAssetsCachePolicyProps.Initializer"></a>

```typescript
import { NextjsAssetsCachePolicyProps } from 'open-next-cdk'

const nextjsAssetsCachePolicyProps: NextjsAssetsCachePolicyProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsAssetsCachePolicyProps.property.staticMaxAgeDefault">staticMaxAgeDefault</a></code> | <code>aws-cdk-lib.Duration</code> | Cache-control max-age default for S3 static assets. |
| <code><a href="#open-next-cdk.NextjsAssetsCachePolicyProps.property.staticStaleWhileRevalidateDefault">staticStaleWhileRevalidateDefault</a></code> | <code>aws-cdk-lib.Duration</code> | Cache-control stale-while-revalidate default for S3 static assets. |

---

##### `staticMaxAgeDefault`<sup>Optional</sup> <a name="staticMaxAgeDefault" id="open-next-cdk.NextjsAssetsCachePolicyProps.property.staticMaxAgeDefault"></a>

```typescript
public readonly staticMaxAgeDefault: Duration;
```

- *Type:* aws-cdk-lib.Duration

Cache-control max-age default for S3 static assets.

Default: 30 days.

---

##### `staticStaleWhileRevalidateDefault`<sup>Optional</sup> <a name="staticStaleWhileRevalidateDefault" id="open-next-cdk.NextjsAssetsCachePolicyProps.property.staticStaleWhileRevalidateDefault"></a>

```typescript
public readonly staticStaleWhileRevalidateDefault: Duration;
```

- *Type:* aws-cdk-lib.Duration

Cache-control stale-while-revalidate default for S3 static assets.

Default: 1 day.

---

### NextjsAssetsDeploymentProps <a name="NextjsAssetsDeploymentProps" id="open-next-cdk.NextjsAssetsDeploymentProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsAssetsDeploymentProps.Initializer"></a>

```typescript
import { NextjsAssetsDeploymentProps } from 'open-next-cdk'

const nextjsAssetsDeploymentProps: NextjsAssetsDeploymentProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.bucket">bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | Properties for the S3 bucket containing the NextJS assets. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.nextBuild">nextBuild</a></code> | <code><a href="#open-next-cdk.NextjsBuild">NextjsBuild</a></code> | The `NextjsBuild` instance representing the built Nextjs application. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.cachePolicies">cachePolicies</a></code> | <code><a href="#open-next-cdk.NextjsAssetsCachePolicyProps">NextjsAssetsCachePolicyProps</a></code> | Override the default S3 cache policies created internally. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.distribution">distribution</a></code> | <code>aws-cdk-lib.aws_cloudfront.IDistribution</code> | Distribution to invalidate when assets change. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.ephemeralStorageSize">ephemeralStorageSize</a></code> | <code>aws-cdk-lib.Size</code> | ephemeralStorageSize for lambda function which been run by BucketDeployment. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.memoryLimit">memoryLimit</a></code> | <code>number</code> | memoryLimit for lambda function which been run by BucketDeployment. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.prune">prune</a></code> | <code>boolean</code> | Set to true to delete old assets (defaults to false). |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.useEfs">useEfs</a></code> | <code>boolean</code> | In case of useEfs, vpc is required. |
| <code><a href="#open-next-cdk.NextjsAssetsDeploymentProps.property.vpc">vpc</a></code> | <code>aws-cdk-lib.aws_ec2.IVpc</code> | In case of useEfs, vpc is required. |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.NextjsAssetsDeploymentProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.NextjsAssetsDeploymentProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.NextjsAssetsDeploymentProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.NextjsAssetsDeploymentProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.NextjsAssetsDeploymentProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.NextjsAssetsDeploymentProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.NextjsAssetsDeploymentProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.NextjsAssetsDeploymentProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.NextjsAssetsDeploymentProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.NextjsAssetsDeploymentProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsAssetsDeploymentProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

##### `bucket`<sup>Required</sup> <a name="bucket" id="open-next-cdk.NextjsAssetsDeploymentProps.property.bucket"></a>

```typescript
public readonly bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

Properties for the S3 bucket containing the NextJS assets.

---

##### `nextBuild`<sup>Required</sup> <a name="nextBuild" id="open-next-cdk.NextjsAssetsDeploymentProps.property.nextBuild"></a>

```typescript
public readonly nextBuild: NextjsBuild;
```

- *Type:* <a href="#open-next-cdk.NextjsBuild">NextjsBuild</a>

The `NextjsBuild` instance representing the built Nextjs application.

---

##### `cachePolicies`<sup>Optional</sup> <a name="cachePolicies" id="open-next-cdk.NextjsAssetsDeploymentProps.property.cachePolicies"></a>

```typescript
public readonly cachePolicies: NextjsAssetsCachePolicyProps;
```

- *Type:* <a href="#open-next-cdk.NextjsAssetsCachePolicyProps">NextjsAssetsCachePolicyProps</a>

Override the default S3 cache policies created internally.

---

##### `distribution`<sup>Optional</sup> <a name="distribution" id="open-next-cdk.NextjsAssetsDeploymentProps.property.distribution"></a>

```typescript
public readonly distribution: IDistribution;
```

- *Type:* aws-cdk-lib.aws_cloudfront.IDistribution

Distribution to invalidate when assets change.

---

##### `ephemeralStorageSize`<sup>Optional</sup> <a name="ephemeralStorageSize" id="open-next-cdk.NextjsAssetsDeploymentProps.property.ephemeralStorageSize"></a>

```typescript
public readonly ephemeralStorageSize: Size;
```

- *Type:* aws-cdk-lib.Size

ephemeralStorageSize for lambda function which been run by BucketDeployment.

---

##### `memoryLimit`<sup>Optional</sup> <a name="memoryLimit" id="open-next-cdk.NextjsAssetsDeploymentProps.property.memoryLimit"></a>

```typescript
public readonly memoryLimit: number;
```

- *Type:* number

memoryLimit for lambda function which been run by BucketDeployment.

---

##### `prune`<sup>Optional</sup> <a name="prune" id="open-next-cdk.NextjsAssetsDeploymentProps.property.prune"></a>

```typescript
public readonly prune: boolean;
```

- *Type:* boolean

Set to true to delete old assets (defaults to false).

Recommended to only set to true if you don't need the ability to roll back deployments.

---

##### `useEfs`<sup>Optional</sup> <a name="useEfs" id="open-next-cdk.NextjsAssetsDeploymentProps.property.useEfs"></a>

```typescript
public readonly useEfs: boolean;
```

- *Type:* boolean

In case of useEfs, vpc is required.

---

##### `vpc`<sup>Optional</sup> <a name="vpc" id="open-next-cdk.NextjsAssetsDeploymentProps.property.vpc"></a>

```typescript
public readonly vpc: IVpc;
```

- *Type:* aws-cdk-lib.aws_ec2.IVpc

In case of useEfs, vpc is required.

---

### NextjsBaseProps <a name="NextjsBaseProps" id="open-next-cdk.NextjsBaseProps"></a>

Common props shared across NextJS-related CDK constructs.

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsBaseProps.Initializer"></a>

```typescript
import { NextjsBaseProps } from 'open-next-cdk'

const nextjsBaseProps: NextjsBaseProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.NextjsBaseProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.NextjsBaseProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.NextjsBaseProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.NextjsBaseProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.NextjsBaseProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.NextjsBaseProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.NextjsBaseProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.NextjsBaseProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.NextjsBaseProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.NextjsBaseProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.NextjsBaseProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsBaseProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

### NextjsBuildProps <a name="NextjsBuildProps" id="open-next-cdk.NextjsBuildProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsBuildProps.Initializer"></a>

```typescript
import { NextjsBuildProps } from 'open-next-cdk'

const nextjsBuildProps: NextjsBuildProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.NextjsBuildProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.NextjsBuildProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.NextjsBuildProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.NextjsBuildProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.NextjsBuildProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.NextjsBuildProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.NextjsBuildProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.NextjsBuildProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.NextjsBuildProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.NextjsBuildProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.NextjsBuildProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsBuildProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

### NextjsCachePolicyProps <a name="NextjsCachePolicyProps" id="open-next-cdk.NextjsCachePolicyProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsCachePolicyProps.Initializer"></a>

```typescript
import { NextjsCachePolicyProps } from 'open-next-cdk'

const nextjsCachePolicyProps: NextjsCachePolicyProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsCachePolicyProps.property.imageCachePolicy">imageCachePolicy</a></code> | <code>aws-cdk-lib.aws_cloudfront.ICachePolicy</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsCachePolicyProps.property.lambdaCachePolicy">lambdaCachePolicy</a></code> | <code>aws-cdk-lib.aws_cloudfront.ICachePolicy</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsCachePolicyProps.property.staticCachePolicy">staticCachePolicy</a></code> | <code>aws-cdk-lib.aws_cloudfront.ICachePolicy</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsCachePolicyProps.property.staticClientMaxAgeDefault">staticClientMaxAgeDefault</a></code> | <code>aws-cdk-lib.Duration</code> | Cache-control max-age default for static assets (/_next/*). |

---

##### `imageCachePolicy`<sup>Optional</sup> <a name="imageCachePolicy" id="open-next-cdk.NextjsCachePolicyProps.property.imageCachePolicy"></a>

```typescript
public readonly imageCachePolicy: ICachePolicy;
```

- *Type:* aws-cdk-lib.aws_cloudfront.ICachePolicy

---

##### `lambdaCachePolicy`<sup>Optional</sup> <a name="lambdaCachePolicy" id="open-next-cdk.NextjsCachePolicyProps.property.lambdaCachePolicy"></a>

```typescript
public readonly lambdaCachePolicy: ICachePolicy;
```

- *Type:* aws-cdk-lib.aws_cloudfront.ICachePolicy

---

##### `staticCachePolicy`<sup>Optional</sup> <a name="staticCachePolicy" id="open-next-cdk.NextjsCachePolicyProps.property.staticCachePolicy"></a>

```typescript
public readonly staticCachePolicy: ICachePolicy;
```

- *Type:* aws-cdk-lib.aws_cloudfront.ICachePolicy

---

##### `staticClientMaxAgeDefault`<sup>Optional</sup> <a name="staticClientMaxAgeDefault" id="open-next-cdk.NextjsCachePolicyProps.property.staticClientMaxAgeDefault"></a>

```typescript
public readonly staticClientMaxAgeDefault: Duration;
```

- *Type:* aws-cdk-lib.Duration

Cache-control max-age default for static assets (/_next/*).

Default: 30 days.

---

### NextjsDefaultsProps <a name="NextjsDefaultsProps" id="open-next-cdk.NextjsDefaultsProps"></a>

Defaults for created resources.

Why `any`? see https://github.com/aws/jsii/issues/2901

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsDefaultsProps.Initializer"></a>

```typescript
import { NextjsDefaultsProps } from 'open-next-cdk'

const nextjsDefaultsProps: NextjsDefaultsProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsDefaultsProps.property.assetDeployment">assetDeployment</a></code> | <code>any</code> | Override static file deployment settings. |
| <code><a href="#open-next-cdk.NextjsDefaultsProps.property.distribution">distribution</a></code> | <code>any</code> | Override CloudFront distribution settings. |
| <code><a href="#open-next-cdk.NextjsDefaultsProps.property.lambda">lambda</a></code> | <code>aws-cdk-lib.aws_lambda.FunctionOptions</code> | Override server lambda function settings. |

---

##### `assetDeployment`<sup>Optional</sup> <a name="assetDeployment" id="open-next-cdk.NextjsDefaultsProps.property.assetDeployment"></a>

```typescript
public readonly assetDeployment: any;
```

- *Type:* any

Override static file deployment settings.

---

##### `distribution`<sup>Optional</sup> <a name="distribution" id="open-next-cdk.NextjsDefaultsProps.property.distribution"></a>

```typescript
public readonly distribution: any;
```

- *Type:* any

Override CloudFront distribution settings.

These properties should all be optional but cannot be due to a limitation in jsii.

---

##### `lambda`<sup>Optional</sup> <a name="lambda" id="open-next-cdk.NextjsDefaultsProps.property.lambda"></a>

```typescript
public readonly lambda: FunctionOptions;
```

- *Type:* aws-cdk-lib.aws_lambda.FunctionOptions

Override server lambda function settings.

---

### NextjsDistributionCdkProps <a name="NextjsDistributionCdkProps" id="open-next-cdk.NextjsDistributionCdkProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsDistributionCdkProps.Initializer"></a>

```typescript
import { NextjsDistributionCdkProps } from 'open-next-cdk'

const nextjsDistributionCdkProps: NextjsDistributionCdkProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsDistributionCdkProps.property.distribution">distribution</a></code> | <code>aws-cdk-lib.aws_cloudfront.DistributionProps</code> | Pass in a value to override the default settings this construct uses to create the CloudFront `Distribution` internally. |

---

##### `distribution`<sup>Optional</sup> <a name="distribution" id="open-next-cdk.NextjsDistributionCdkProps.property.distribution"></a>

```typescript
public readonly distribution: DistributionProps;
```

- *Type:* aws-cdk-lib.aws_cloudfront.DistributionProps

Pass in a value to override the default settings this construct uses to create the CloudFront `Distribution` internally.

---

### NextjsDistributionProps <a name="NextjsDistributionProps" id="open-next-cdk.NextjsDistributionProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsDistributionProps.Initializer"></a>

```typescript
import { NextjsDistributionProps } from 'open-next-cdk'

const nextjsDistributionProps: NextjsDistributionProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.imageOptFunction">imageOptFunction</a></code> | <code>aws-cdk-lib.aws_lambda.IFunction</code> | Lambda function to optimize images. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.nextBuild">nextBuild</a></code> | <code><a href="#open-next-cdk.NextjsBuild">NextjsBuild</a></code> | Built NextJS app. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.serverFunction">serverFunction</a></code> | <code>aws-cdk-lib.aws_lambda.IFunction</code> | Lambda function to route all non-static requests to. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.staticAssetsBucket">staticAssetsBucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | Bucket containing static assets. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.cachePolicies">cachePolicies</a></code> | <code><a href="#open-next-cdk.NextjsCachePolicyProps">NextjsCachePolicyProps</a></code> | Override the default CloudFront cache policies created internally. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.cdk">cdk</a></code> | <code><a href="#open-next-cdk.NextjsDistributionCdkProps">NextjsDistributionCdkProps</a></code> | Overrides for created CDK resources. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.customDomain">customDomain</a></code> | <code>string \| <a href="#open-next-cdk.NextjsDomainProps">NextjsDomainProps</a></code> | The customDomain for this website. Supports domains that are hosted either on [Route 53](https://aws.amazon.com/route53/) or externally. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.functionUrlAuthType">functionUrlAuthType</a></code> | <code>aws-cdk-lib.aws_lambda.FunctionUrlAuthType</code> | Override lambda function url auth type. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.originRequestPolicies">originRequestPolicies</a></code> | <code><a href="#open-next-cdk.NextjsOriginRequestPolicyProps">NextjsOriginRequestPolicyProps</a></code> | Override the default CloudFront origin request policies created internally. |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.stackPrefix">stackPrefix</a></code> | <code>string</code> | Optional value to prefix the edge function stack It defaults to "Nextjs". |
| <code><a href="#open-next-cdk.NextjsDistributionProps.property.stageName">stageName</a></code> | <code>string</code> | Include the name of your deployment stage if present. |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.NextjsDistributionProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.NextjsDistributionProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.NextjsDistributionProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.NextjsDistributionProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.NextjsDistributionProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.NextjsDistributionProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.NextjsDistributionProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.NextjsDistributionProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.NextjsDistributionProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.NextjsDistributionProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsDistributionProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

##### `imageOptFunction`<sup>Required</sup> <a name="imageOptFunction" id="open-next-cdk.NextjsDistributionProps.property.imageOptFunction"></a>

```typescript
public readonly imageOptFunction: IFunction;
```

- *Type:* aws-cdk-lib.aws_lambda.IFunction

Lambda function to optimize images.

Must be provided if you want to serve dynamic requests.

---

##### `nextBuild`<sup>Required</sup> <a name="nextBuild" id="open-next-cdk.NextjsDistributionProps.property.nextBuild"></a>

```typescript
public readonly nextBuild: NextjsBuild;
```

- *Type:* <a href="#open-next-cdk.NextjsBuild">NextjsBuild</a>

Built NextJS app.

---

##### `serverFunction`<sup>Required</sup> <a name="serverFunction" id="open-next-cdk.NextjsDistributionProps.property.serverFunction"></a>

```typescript
public readonly serverFunction: IFunction;
```

- *Type:* aws-cdk-lib.aws_lambda.IFunction

Lambda function to route all non-static requests to.

Must be provided if you want to serve dynamic requests.

---

##### `staticAssetsBucket`<sup>Required</sup> <a name="staticAssetsBucket" id="open-next-cdk.NextjsDistributionProps.property.staticAssetsBucket"></a>

```typescript
public readonly staticAssetsBucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

Bucket containing static assets.

Must be provided if you want to serve static files.

---

##### `cachePolicies`<sup>Optional</sup> <a name="cachePolicies" id="open-next-cdk.NextjsDistributionProps.property.cachePolicies"></a>

```typescript
public readonly cachePolicies: NextjsCachePolicyProps;
```

- *Type:* <a href="#open-next-cdk.NextjsCachePolicyProps">NextjsCachePolicyProps</a>

Override the default CloudFront cache policies created internally.

---

##### `cdk`<sup>Optional</sup> <a name="cdk" id="open-next-cdk.NextjsDistributionProps.property.cdk"></a>

```typescript
public readonly cdk: NextjsDistributionCdkProps;
```

- *Type:* <a href="#open-next-cdk.NextjsDistributionCdkProps">NextjsDistributionCdkProps</a>

Overrides for created CDK resources.

---

##### `customDomain`<sup>Optional</sup> <a name="customDomain" id="open-next-cdk.NextjsDistributionProps.property.customDomain"></a>

```typescript
public readonly customDomain: string | NextjsDomainProps;
```

- *Type:* string | <a href="#open-next-cdk.NextjsDomainProps">NextjsDomainProps</a>

The customDomain for this website. Supports domains that are hosted either on [Route 53](https://aws.amazon.com/route53/) or externally.

Note that you can also migrate externally hosted domains to Route 53 by
[following this guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html).

---

*Example*

```typescript
new NextjsDistribution(this, "Dist", {
  customDomain: "domain.com",
});

new NextjsDistribution(this, "Dist", {
  customDomain: {
    domainName: "domain.com",
    domainAlias: "www.domain.com",
    hostedZone: "domain.com"
  },
});
```


##### `functionUrlAuthType`<sup>Optional</sup> <a name="functionUrlAuthType" id="open-next-cdk.NextjsDistributionProps.property.functionUrlAuthType"></a>

```typescript
public readonly functionUrlAuthType: FunctionUrlAuthType;
```

- *Type:* aws-cdk-lib.aws_lambda.FunctionUrlAuthType
- *Default:* "NONE"

Override lambda function url auth type.

---

##### `originRequestPolicies`<sup>Optional</sup> <a name="originRequestPolicies" id="open-next-cdk.NextjsDistributionProps.property.originRequestPolicies"></a>

```typescript
public readonly originRequestPolicies: NextjsOriginRequestPolicyProps;
```

- *Type:* <a href="#open-next-cdk.NextjsOriginRequestPolicyProps">NextjsOriginRequestPolicyProps</a>

Override the default CloudFront origin request policies created internally.

---

##### `stackPrefix`<sup>Optional</sup> <a name="stackPrefix" id="open-next-cdk.NextjsDistributionProps.property.stackPrefix"></a>

```typescript
public readonly stackPrefix: string;
```

- *Type:* string

Optional value to prefix the edge function stack It defaults to "Nextjs".

---

##### `stageName`<sup>Optional</sup> <a name="stageName" id="open-next-cdk.NextjsDistributionProps.property.stageName"></a>

```typescript
public readonly stageName: string;
```

- *Type:* string

Include the name of your deployment stage if present.

Used to name the edge functions stack.
Required if using SST.

---

### NextjsDomainProps <a name="NextjsDomainProps" id="open-next-cdk.NextjsDomainProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsDomainProps.Initializer"></a>

```typescript
import { NextjsDomainProps } from 'open-next-cdk'

const nextjsDomainProps: NextjsDomainProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsDomainProps.property.domainName">domainName</a></code> | <code>string</code> | The domain to be assigned to the website URL (ie. domain.com). |
| <code><a href="#open-next-cdk.NextjsDomainProps.property.alternateNames">alternateNames</a></code> | <code>string[]</code> | Specify additional names that should route to the Cloudfront Distribution. |
| <code><a href="#open-next-cdk.NextjsDomainProps.property.certificate">certificate</a></code> | <code>aws-cdk-lib.aws_certificatemanager.ICertificate</code> | Import the certificate for the domain. |
| <code><a href="#open-next-cdk.NextjsDomainProps.property.domainAlias">domainAlias</a></code> | <code>string</code> | An alternative domain to be assigned to the website URL. |
| <code><a href="#open-next-cdk.NextjsDomainProps.property.hostedZone">hostedZone</a></code> | <code>aws-cdk-lib.aws_route53.IHostedZone</code> | Import the underlying Route 53 hosted zone. |
| <code><a href="#open-next-cdk.NextjsDomainProps.property.isExternalDomain">isExternalDomain</a></code> | <code>boolean</code> | Set this option if the domain is not hosted on Amazon Route 53. |

---

##### `domainName`<sup>Required</sup> <a name="domainName" id="open-next-cdk.NextjsDomainProps.property.domainName"></a>

```typescript
public readonly domainName: string;
```

- *Type:* string

The domain to be assigned to the website URL (ie. domain.com).

Supports domains that are hosted either on [Route 53](https://aws.amazon.com/route53/) or externally.

---

##### `alternateNames`<sup>Optional</sup> <a name="alternateNames" id="open-next-cdk.NextjsDomainProps.property.alternateNames"></a>

```typescript
public readonly alternateNames: string[];
```

- *Type:* string[]

Specify additional names that should route to the Cloudfront Distribution.

Note, certificates for these names will not be automatically generated so the `certificate` option must be specified.

---

##### `certificate`<sup>Optional</sup> <a name="certificate" id="open-next-cdk.NextjsDomainProps.property.certificate"></a>

```typescript
public readonly certificate: ICertificate;
```

- *Type:* aws-cdk-lib.aws_certificatemanager.ICertificate

Import the certificate for the domain.

By default, SST will create a certificate with the domain name. The certificate will be created in the `us-east-1`(N. Virginia) region as required by AWS CloudFront.

Set this option if you have an existing certificate in the `us-east-1` region in AWS Certificate Manager you want to use.

---

##### `domainAlias`<sup>Optional</sup> <a name="domainAlias" id="open-next-cdk.NextjsDomainProps.property.domainAlias"></a>

```typescript
public readonly domainAlias: string;
```

- *Type:* string

An alternative domain to be assigned to the website URL.

Visitors to the alias will be redirected to the main domain. (ie. `www.domain.com`).

Use this to create a `www.` version of your domain and redirect visitors to the root domain.

---

##### `hostedZone`<sup>Optional</sup> <a name="hostedZone" id="open-next-cdk.NextjsDomainProps.property.hostedZone"></a>

```typescript
public readonly hostedZone: IHostedZone;
```

- *Type:* aws-cdk-lib.aws_route53.IHostedZone

Import the underlying Route 53 hosted zone.

---

##### `isExternalDomain`<sup>Optional</sup> <a name="isExternalDomain" id="open-next-cdk.NextjsDomainProps.property.isExternalDomain"></a>

```typescript
public readonly isExternalDomain: boolean;
```

- *Type:* boolean

Set this option if the domain is not hosted on Amazon Route 53.

---

### NextjsLambdaProps <a name="NextjsLambdaProps" id="open-next-cdk.NextjsLambdaProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsLambdaProps.Initializer"></a>

```typescript
import { NextjsLambdaProps } from 'open-next-cdk'

const nextjsLambdaProps: NextjsLambdaProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.nextBuild">nextBuild</a></code> | <code><a href="#open-next-cdk.NextjsBuild">NextjsBuild</a></code> | Built nextJS application. |
| <code><a href="#open-next-cdk.NextjsLambdaProps.property.lambda">lambda</a></code> | <code>aws-cdk-lib.aws_lambda.FunctionOptions</code> | Override function properties. |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.NextjsLambdaProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.NextjsLambdaProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.NextjsLambdaProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.NextjsLambdaProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.NextjsLambdaProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.NextjsLambdaProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.NextjsLambdaProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.NextjsLambdaProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.NextjsLambdaProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.NextjsLambdaProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsLambdaProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

##### `nextBuild`<sup>Required</sup> <a name="nextBuild" id="open-next-cdk.NextjsLambdaProps.property.nextBuild"></a>

```typescript
public readonly nextBuild: NextjsBuild;
```

- *Type:* <a href="#open-next-cdk.NextjsBuild">NextjsBuild</a>

Built nextJS application.

---

##### `lambda`<sup>Optional</sup> <a name="lambda" id="open-next-cdk.NextjsLambdaProps.property.lambda"></a>

```typescript
public readonly lambda: FunctionOptions;
```

- *Type:* aws-cdk-lib.aws_lambda.FunctionOptions

Override function properties.

---

### NextjsLayerProps <a name="NextjsLayerProps" id="open-next-cdk.NextjsLayerProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsLayerProps.Initializer"></a>

```typescript
import { NextjsLayerProps } from 'open-next-cdk'

const nextjsLayerProps: NextjsLayerProps = { ... }
```


### NextjsOriginRequestPolicyProps <a name="NextjsOriginRequestPolicyProps" id="open-next-cdk.NextjsOriginRequestPolicyProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsOriginRequestPolicyProps.Initializer"></a>

```typescript
import { NextjsOriginRequestPolicyProps } from 'open-next-cdk'

const nextjsOriginRequestPolicyProps: NextjsOriginRequestPolicyProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsOriginRequestPolicyProps.property.fallbackOriginRequestPolicy">fallbackOriginRequestPolicy</a></code> | <code>aws-cdk-lib.aws_cloudfront.IOriginRequestPolicy</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsOriginRequestPolicyProps.property.imageOptimizationOriginRequestPolicy">imageOptimizationOriginRequestPolicy</a></code> | <code>aws-cdk-lib.aws_cloudfront.IOriginRequestPolicy</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsOriginRequestPolicyProps.property.lambdaOriginRequestPolicy">lambdaOriginRequestPolicy</a></code> | <code>aws-cdk-lib.aws_cloudfront.IOriginRequestPolicy</code> | *No description.* |

---

##### `fallbackOriginRequestPolicy`<sup>Optional</sup> <a name="fallbackOriginRequestPolicy" id="open-next-cdk.NextjsOriginRequestPolicyProps.property.fallbackOriginRequestPolicy"></a>

```typescript
public readonly fallbackOriginRequestPolicy: IOriginRequestPolicy;
```

- *Type:* aws-cdk-lib.aws_cloudfront.IOriginRequestPolicy

---

##### `imageOptimizationOriginRequestPolicy`<sup>Optional</sup> <a name="imageOptimizationOriginRequestPolicy" id="open-next-cdk.NextjsOriginRequestPolicyProps.property.imageOptimizationOriginRequestPolicy"></a>

```typescript
public readonly imageOptimizationOriginRequestPolicy: IOriginRequestPolicy;
```

- *Type:* aws-cdk-lib.aws_cloudfront.IOriginRequestPolicy

---

##### `lambdaOriginRequestPolicy`<sup>Optional</sup> <a name="lambdaOriginRequestPolicy" id="open-next-cdk.NextjsOriginRequestPolicyProps.property.lambdaOriginRequestPolicy"></a>

```typescript
public readonly lambdaOriginRequestPolicy: IOriginRequestPolicy;
```

- *Type:* aws-cdk-lib.aws_cloudfront.IOriginRequestPolicy

---

### NextjsProps <a name="NextjsProps" id="open-next-cdk.NextjsProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsProps.Initializer"></a>

```typescript
import { NextjsProps } from 'open-next-cdk'

const nextjsProps: NextjsProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.NextjsProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.NextjsProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.NextjsProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.NextjsProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.NextjsProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.NextjsProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.NextjsProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.NextjsProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.NextjsProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.NextjsProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |
| <code><a href="#open-next-cdk.NextjsProps.property.defaults">defaults</a></code> | <code><a href="#open-next-cdk.NextjsDefaultsProps">NextjsDefaultsProps</a></code> | Allows you to override defaults for the resources created by this construct. |
| <code><a href="#open-next-cdk.NextjsProps.property.imageOptimizationBucket">imageOptimizationBucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | Optional S3 Bucket to use, defaults to assets bucket. |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.NextjsProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.NextjsProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.NextjsProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.NextjsProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.NextjsProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.NextjsProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.NextjsProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.NextjsProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.NextjsProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.NextjsProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

##### `defaults`<sup>Optional</sup> <a name="defaults" id="open-next-cdk.NextjsProps.property.defaults"></a>

```typescript
public readonly defaults: NextjsDefaultsProps;
```

- *Type:* <a href="#open-next-cdk.NextjsDefaultsProps">NextjsDefaultsProps</a>

Allows you to override defaults for the resources created by this construct.

---

##### `imageOptimizationBucket`<sup>Optional</sup> <a name="imageOptimizationBucket" id="open-next-cdk.NextjsProps.property.imageOptimizationBucket"></a>

```typescript
public readonly imageOptimizationBucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

Optional S3 Bucket to use, defaults to assets bucket.

---

### NextjsS3EnvRewriterProps <a name="NextjsS3EnvRewriterProps" id="open-next-cdk.NextjsS3EnvRewriterProps"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.NextjsS3EnvRewriterProps.Initializer"></a>

```typescript
import { NextjsS3EnvRewriterProps } from 'open-next-cdk'

const nextjsS3EnvRewriterProps: NextjsS3EnvRewriterProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.nextjsPath">nextjsPath</a></code> | <code>string</code> | Relative path to the directory where the NextJS project is located. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.buildCommand">buildCommand</a></code> | <code>string</code> | Optional value used to install NextJS node dependencies. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.buildPath">buildPath</a></code> | <code>string</code> | The directory to execute `npm run build` from. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.compressionLevel">compressionLevel</a></code> | <code>number</code> | 0 - no compression, fastest 9 - maximum compression, slowest. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.environment">environment</a></code> | <code>{[ key: string ]: string}</code> | Custom environment variables to pass to the NextJS build and runtime. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.isPlaceholder">isPlaceholder</a></code> | <code>boolean</code> | Skip building app and deploy a placeholder. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.nodeEnv">nodeEnv</a></code> | <code>string</code> | Optional value for NODE_ENV during build and runtime. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.projectRoot">projectRoot</a></code> | <code>string</code> | Root of your project, if different from `nextjsPath`. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.quiet">quiet</a></code> | <code>boolean</code> | Less build output. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.sharpLayerArn">sharpLayerArn</a></code> | <code>string</code> | Optional arn for the sharp lambda layer. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.tempBuildDir">tempBuildDir</a></code> | <code>string</code> | Directory to store temporary build files in. |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.replacementConfig">replacementConfig</a></code> | <code><a href="#open-next-cdk.RewriteReplacementsConfig">RewriteReplacementsConfig</a></code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.s3Bucket">s3Bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.s3keys">s3keys</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.cloudfrontDistributionId">cloudfrontDistributionId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.NextjsS3EnvRewriterProps.property.debug">debug</a></code> | <code>boolean</code> | *No description.* |

---

##### `nextjsPath`<sup>Required</sup> <a name="nextjsPath" id="open-next-cdk.NextjsS3EnvRewriterProps.property.nextjsPath"></a>

```typescript
public readonly nextjsPath: string;
```

- *Type:* string

Relative path to the directory where the NextJS project is located.

Can be the root of your project (`.`) or a subdirectory (`packages/web`).

---

##### `buildCommand`<sup>Optional</sup> <a name="buildCommand" id="open-next-cdk.NextjsS3EnvRewriterProps.property.buildCommand"></a>

```typescript
public readonly buildCommand: string;
```

- *Type:* string

Optional value used to install NextJS node dependencies.

It defaults to 'npx --yes open-next@latest build'

---

##### `buildPath`<sup>Optional</sup> <a name="buildPath" id="open-next-cdk.NextjsS3EnvRewriterProps.property.buildPath"></a>

```typescript
public readonly buildPath: string;
```

- *Type:* string

The directory to execute `npm run build` from.

By default, it is `nextjsPath`.
Can be overridden, particularly useful for monorepos where `build` is expected to run
at the root of the project.

---

##### `compressionLevel`<sup>Optional</sup> <a name="compressionLevel" id="open-next-cdk.NextjsS3EnvRewriterProps.property.compressionLevel"></a>

```typescript
public readonly compressionLevel: number;
```

- *Type:* number
- *Default:* 1

0 - no compression, fastest 9 - maximum compression, slowest.

---

##### `environment`<sup>Optional</sup> <a name="environment" id="open-next-cdk.NextjsS3EnvRewriterProps.property.environment"></a>

```typescript
public readonly environment: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Custom environment variables to pass to the NextJS build and runtime.

---

##### `isPlaceholder`<sup>Optional</sup> <a name="isPlaceholder" id="open-next-cdk.NextjsS3EnvRewriterProps.property.isPlaceholder"></a>

```typescript
public readonly isPlaceholder: boolean;
```

- *Type:* boolean

Skip building app and deploy a placeholder.

Useful when using `next dev` for local development.

---

##### `nodeEnv`<sup>Optional</sup> <a name="nodeEnv" id="open-next-cdk.NextjsS3EnvRewriterProps.property.nodeEnv"></a>

```typescript
public readonly nodeEnv: string;
```

- *Type:* string

Optional value for NODE_ENV during build and runtime.

---

##### `projectRoot`<sup>Optional</sup> <a name="projectRoot" id="open-next-cdk.NextjsS3EnvRewriterProps.property.projectRoot"></a>

```typescript
public readonly projectRoot: string;
```

- *Type:* string

Root of your project, if different from `nextjsPath`.

Defaults to current working directory.

---

##### `quiet`<sup>Optional</sup> <a name="quiet" id="open-next-cdk.NextjsS3EnvRewriterProps.property.quiet"></a>

```typescript
public readonly quiet: boolean;
```

- *Type:* boolean

Less build output.

---

##### `sharpLayerArn`<sup>Optional</sup> <a name="sharpLayerArn" id="open-next-cdk.NextjsS3EnvRewriterProps.property.sharpLayerArn"></a>

```typescript
public readonly sharpLayerArn: string;
```

- *Type:* string

Optional arn for the sharp lambda layer.

If omitted, the layer will be created.

---

##### `tempBuildDir`<sup>Optional</sup> <a name="tempBuildDir" id="open-next-cdk.NextjsS3EnvRewriterProps.property.tempBuildDir"></a>

```typescript
public readonly tempBuildDir: string;
```

- *Type:* string

Directory to store temporary build files in.

Defaults to os.tmpdir().

---

##### `replacementConfig`<sup>Required</sup> <a name="replacementConfig" id="open-next-cdk.NextjsS3EnvRewriterProps.property.replacementConfig"></a>

```typescript
public readonly replacementConfig: RewriteReplacementsConfig;
```

- *Type:* <a href="#open-next-cdk.RewriteReplacementsConfig">RewriteReplacementsConfig</a>

---

##### `s3Bucket`<sup>Required</sup> <a name="s3Bucket" id="open-next-cdk.NextjsS3EnvRewriterProps.property.s3Bucket"></a>

```typescript
public readonly s3Bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

---

##### `s3keys`<sup>Required</sup> <a name="s3keys" id="open-next-cdk.NextjsS3EnvRewriterProps.property.s3keys"></a>

```typescript
public readonly s3keys: string[];
```

- *Type:* string[]

---

##### `cloudfrontDistributionId`<sup>Optional</sup> <a name="cloudfrontDistributionId" id="open-next-cdk.NextjsS3EnvRewriterProps.property.cloudfrontDistributionId"></a>

```typescript
public readonly cloudfrontDistributionId: string;
```

- *Type:* string

---

##### `debug`<sup>Optional</sup> <a name="debug" id="open-next-cdk.NextjsS3EnvRewriterProps.property.debug"></a>

```typescript
public readonly debug: boolean;
```

- *Type:* boolean

---

### RewriteReplacementsConfig <a name="RewriteReplacementsConfig" id="open-next-cdk.RewriteReplacementsConfig"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.RewriteReplacementsConfig.Initializer"></a>

```typescript
import { RewriteReplacementsConfig } from 'open-next-cdk'

const rewriteReplacementsConfig: RewriteReplacementsConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.RewriteReplacementsConfig.property.env">env</a></code> | <code>{[ key: string ]: string}</code> | *No description.* |
| <code><a href="#open-next-cdk.RewriteReplacementsConfig.property.jsonS3Bucket">jsonS3Bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | *No description.* |
| <code><a href="#open-next-cdk.RewriteReplacementsConfig.property.jsonS3Key">jsonS3Key</a></code> | <code>string</code> | *No description.* |

---

##### `env`<sup>Optional</sup> <a name="env" id="open-next-cdk.RewriteReplacementsConfig.property.env"></a>

```typescript
public readonly env: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

---

##### `jsonS3Bucket`<sup>Optional</sup> <a name="jsonS3Bucket" id="open-next-cdk.RewriteReplacementsConfig.property.jsonS3Bucket"></a>

```typescript
public readonly jsonS3Bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

---

##### `jsonS3Key`<sup>Optional</sup> <a name="jsonS3Key" id="open-next-cdk.RewriteReplacementsConfig.property.jsonS3Key"></a>

```typescript
public readonly jsonS3Key: string;
```

- *Type:* string

---

### RewriterParams <a name="RewriterParams" id="open-next-cdk.RewriterParams"></a>

#### Initializer <a name="Initializer" id="open-next-cdk.RewriterParams.Initializer"></a>

```typescript
import { RewriterParams } from 'open-next-cdk'

const rewriterParams: RewriterParams = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#open-next-cdk.RewriterParams.property.replacementConfig">replacementConfig</a></code> | <code><a href="#open-next-cdk.RewriteReplacementsConfig">RewriteReplacementsConfig</a></code> | *No description.* |
| <code><a href="#open-next-cdk.RewriterParams.property.s3Bucket">s3Bucket</a></code> | <code>aws-cdk-lib.aws_s3.IBucket</code> | *No description.* |
| <code><a href="#open-next-cdk.RewriterParams.property.s3keys">s3keys</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#open-next-cdk.RewriterParams.property.cloudfrontDistributionId">cloudfrontDistributionId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#open-next-cdk.RewriterParams.property.debug">debug</a></code> | <code>boolean</code> | *No description.* |

---

##### `replacementConfig`<sup>Required</sup> <a name="replacementConfig" id="open-next-cdk.RewriterParams.property.replacementConfig"></a>

```typescript
public readonly replacementConfig: RewriteReplacementsConfig;
```

- *Type:* <a href="#open-next-cdk.RewriteReplacementsConfig">RewriteReplacementsConfig</a>

---

##### `s3Bucket`<sup>Required</sup> <a name="s3Bucket" id="open-next-cdk.RewriterParams.property.s3Bucket"></a>

```typescript
public readonly s3Bucket: IBucket;
```

- *Type:* aws-cdk-lib.aws_s3.IBucket

---

##### `s3keys`<sup>Required</sup> <a name="s3keys" id="open-next-cdk.RewriterParams.property.s3keys"></a>

```typescript
public readonly s3keys: string[];
```

- *Type:* string[]

---

##### `cloudfrontDistributionId`<sup>Optional</sup> <a name="cloudfrontDistributionId" id="open-next-cdk.RewriterParams.property.cloudfrontDistributionId"></a>

```typescript
public readonly cloudfrontDistributionId: string;
```

- *Type:* string

---

##### `debug`<sup>Optional</sup> <a name="debug" id="open-next-cdk.RewriterParams.property.debug"></a>

```typescript
public readonly debug: boolean;
```

- *Type:* boolean

---



