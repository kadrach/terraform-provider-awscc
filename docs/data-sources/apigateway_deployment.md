---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigateway_deployment Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ApiGateway::Deployment
---

# awscc_apigateway_deployment (Data Source)

Data Source schema for AWS::ApiGateway::Deployment



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `deployment_canary_settings` (Attributes) The input configuration for a canary deployment. (see [below for nested schema](#nestedatt--deployment_canary_settings))
- `deployment_id` (String) Primary Id for this resource
- `description` (String) The description for the Deployment resource to create.
- `rest_api_id` (String) The string identifier of the associated RestApi.
- `stage_description` (Attributes) The description of the Stage resource for the Deployment resource to create. To specify a stage description, you must also provide a stage name. (see [below for nested schema](#nestedatt--stage_description))
- `stage_name` (String) The name of the Stage resource for the Deployment resource to create.

<a id="nestedatt--deployment_canary_settings"></a>
### Nested Schema for `deployment_canary_settings`

Read-Only:

- `percent_traffic` (Number) The percentage (0.0-100.0) of traffic routed to the canary deployment.
- `stage_variable_overrides` (Map of String) A stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.
- `use_stage_cache` (Boolean) A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.


<a id="nestedatt--stage_description"></a>
### Nested Schema for `stage_description`

Read-Only:

- `access_log_setting` (Attributes) Specifies settings for logging access in this stage. (see [below for nested schema](#nestedatt--stage_description--access_log_setting))
- `cache_cluster_enabled` (Boolean) Specifies whether a cache cluster is enabled for the stage.
- `cache_cluster_size` (String) The size of the stage's cache cluster. For more information, see [cacheClusterSize](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateStage.html#apigw-CreateStage-request-cacheClusterSize) in the *API Gateway API Reference*.
- `cache_data_encrypted` (Boolean) Indicates whether the cached responses are encrypted.
- `cache_ttl_in_seconds` (Number) The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.
- `caching_enabled` (Boolean) Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide*.
- `canary_setting` (Attributes) Specifies settings for the canary deployment in this stage. (see [below for nested schema](#nestedatt--stage_description--canary_setting))
- `client_certificate_id` (String) The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.
- `data_trace_enabled` (Boolean) Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.
- `description` (String) A description of the purpose of the stage.
- `documentation_version` (String) The version identifier of the API documentation snapshot.
- `logging_level` (String) The logging level for this method. For valid values, see the ``loggingLevel`` property of the [MethodSetting](https://docs.aws.amazon.com/apigateway/latest/api/API_MethodSetting.html) resource in the *Amazon API Gateway API Reference*.
- `method_settings` (Attributes Set) Configures settings for all of the stage's methods. (see [below for nested schema](#nestedatt--stage_description--method_settings))
- `metrics_enabled` (Boolean) Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.
- `tags` (Attributes List) An array of arbitrary tags (key-value pairs) to associate with the stage. (see [below for nested schema](#nestedatt--stage_description--tags))
- `throttling_burst_limit` (Number) The target request burst rate limit. This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.
- `throttling_rate_limit` (Number) The target request steady-state rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.
- `tracing_enabled` (Boolean) Specifies whether active tracing with X-ray is enabled for this stage.
 For more information, see [Trace API Gateway API Execution with X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide*.
- `variables` (Map of String) A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&=,]+``.

<a id="nestedatt--stage_description--access_log_setting"></a>
### Nested Schema for `stage_description.access_log_setting`

Read-Only:

- `destination_arn` (String) The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``.
- `format` (String) A single line format of the access logs of data, as specified by selected $context variables. The format must include at least ``$context.requestId``.


<a id="nestedatt--stage_description--canary_setting"></a>
### Nested Schema for `stage_description.canary_setting`

Read-Only:

- `percent_traffic` (Number) The percent (0-100) of traffic diverted to a canary deployment.
- `stage_variable_overrides` (Map of String) Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.
- `use_stage_cache` (Boolean) A Boolean flag to indicate whether the canary deployment uses the stage cache or not.


<a id="nestedatt--stage_description--method_settings"></a>
### Nested Schema for `stage_description.method_settings`

Read-Only:

- `cache_data_encrypted` (Boolean) Specifies whether the cached responses are encrypted.
- `cache_ttl_in_seconds` (Number) Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
- `caching_enabled` (Boolean) Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.
- `data_trace_enabled` (Boolean) Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.
- `http_method` (String) The HTTP method.
- `logging_level` (String) Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are ``OFF``, ``ERROR``, and ``INFO``. Choose ``ERROR`` to write only error-level entries to CloudWatch Logs, or choose ``INFO`` to include all ``ERROR`` events as well as extra informational events.
- `metrics_enabled` (Boolean) Specifies whether Amazon CloudWatch metrics are enabled for this method.
- `resource_path` (String) The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``).
- `throttling_burst_limit` (Number) Specifies the throttling burst limit.
- `throttling_rate_limit` (Number) Specifies the throttling rate limit.


<a id="nestedatt--stage_description--tags"></a>
### Nested Schema for `stage_description.tags`

Read-Only:

- `key` (String) The key name of the tag
- `value` (String) The value for the tag
