// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_cloudwatch_metric_stream", metricStreamResourceType)
}

// metricStreamResourceType returns the Terraform aws_cloudwatch_metric_stream resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudWatch::MetricStream resource type.
func metricStreamResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Amazon Resource Name of the metric stream.",
			     "maxLength": 2048,
			     "minLength": 20,
			     "type": "string"
			   }
			*/
			Description: "Amazon Resource Name of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The date of creation of the metric stream.",
			     "format": "date-time",
			     "type": "string"
			   }
			*/
			Description: "The date of creation of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"exclude_filters": {
			// Property: ExcludeFilters
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			     "items": {
			       "additionalProperties": false,
			       "description": "This structure defines the metrics that will be streamed.",
			       "properties": {
			         "Namespace": {
			           "description": "Only metrics with Namespace matching this value will be streamed.",
			           "maxLength": 255,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/MetricStreamFilter",
			       "required": [
			         "Namespace"
			       ],
			       "type": "object"
			     },
			     "maxItems": 1000,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"namespace": {
						// Property: Namespace
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Only metrics with Namespace matching this value will be streamed.",
						     "maxLength": 255,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 1000,
				},
			),
			Optional: true,
		},
		"firehose_arn": {
			// Property: FirehoseArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the Kinesis Firehose where to stream the data.",
			     "maxLength": 2048,
			     "minLength": 20,
			     "type": "string"
			   }
			*/
			Description: "The ARN of the Kinesis Firehose where to stream the data.",
			Type:        types.StringType,
			Required:    true,
		},
		"include_filters": {
			// Property: IncludeFilters
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			     "items": {
			       "additionalProperties": false,
			       "description": "This structure defines the metrics that will be streamed.",
			       "properties": {
			         "Namespace": {
			           "description": "Only metrics with Namespace matching this value will be streamed.",
			           "maxLength": 255,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/MetricStreamFilter",
			       "required": [
			         "Namespace"
			       ],
			       "type": "object"
			     },
			     "maxItems": 1000,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"namespace": {
						// Property: Namespace
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Only metrics with Namespace matching this value will be streamed.",
						     "maxLength": 255,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 1000,
				},
			),
			Optional: true,
		},
		"last_update_date": {
			// Property: LastUpdateDate
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The date of the last update of the metric stream.",
			     "format": "date-time",
			     "type": "string"
			   }
			*/
			Description: "The date of the last update of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Name of the metric stream.",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Name of the metric stream.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"output_format": {
			// Property: OutputFormat
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The output format of the data streamed to the Kinesis Firehose.",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The output format of the data streamed to the Kinesis Firehose.",
			Type:        types.StringType,
			Required:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the role that provides access to the Kinesis Firehose.",
			     "maxLength": 2048,
			     "minLength": 20,
			     "type": "string"
			   }
			*/
			Description: "The ARN of the role that provides access to the Kinesis Firehose.",
			Type:        types.StringType,
			Required:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Displays the state of the Metric Stream.",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Displays the state of the Metric Stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A set of tags to assign to the delivery stream.",
			     "items": {
			       "additionalProperties": false,
			       "description": "Metadata that you can assign to a Metric Stream, consisting of a key-value pair.",
			       "properties": {
			         "Key": {
			           "description": "A unique identifier for the tag.",
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "An optional string, which you can use to describe or define the tag.",
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "A set of tags to assign to the delivery stream.",
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A unique identifier for the tag.",
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "A unique identifier for the tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "An optional string, which you can use to describe or define the tag.",
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "An optional string, which you can use to describe or define the tag.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
			// Tags is a write-only attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for Metric Stream",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::MetricStream").WithTerraformTypeName("aws_cloudwatch_metric_stream").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_cloudwatch_metric_stream", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}