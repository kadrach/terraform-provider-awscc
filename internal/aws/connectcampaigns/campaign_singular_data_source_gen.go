// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connectcampaigns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_connectcampaigns_campaign", campaignDataSourceType)
}

// campaignDataSourceType returns the Terraform awscc_connectcampaigns_campaign data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ConnectCampaigns::Campaign resource type.
func campaignDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Connect Campaign Arn",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "^arn:aws[-a-z0-9]*:connect-campaigns:[-a-z0-9]*:[0-9]{12}:campaign/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "Amazon Connect Campaign Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"connect_instance_arn": {
			// Property: ConnectInstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Connect Instance Arn",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "Amazon Connect Instance Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"dialer_config": {
			// Property: DialerConfig
			// CloudFormation resource type schema:
			// {
			//   "description": "The possible types of dialer config parameters",
			//   "properties": {
			//     "PredictiveDialerConfig": {
			//       "additionalProperties": false,
			//       "description": "Predictive Dialer config",
			//       "properties": {
			//         "BandwidthAllocation": {
			//           "description": "The bandwidth allocation of a queue resource.",
			//           "maximum": 1,
			//           "minimum": 0,
			//           "type": "number"
			//         }
			//       },
			//       "required": [
			//         "BandwidthAllocation"
			//       ],
			//       "type": "object"
			//     },
			//     "ProgressiveDialerConfig": {
			//       "additionalProperties": false,
			//       "description": "Progressive Dialer config",
			//       "properties": {
			//         "BandwidthAllocation": {
			//           "description": "The bandwidth allocation of a queue resource.",
			//           "maximum": 1,
			//           "minimum": 0,
			//           "type": "number"
			//         }
			//       },
			//       "required": [
			//         "BandwidthAllocation"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The possible types of dialer config parameters",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"predictive_dialer_config": {
						// Property: PredictiveDialerConfig
						Description: "Predictive Dialer config",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bandwidth_allocation": {
									// Property: BandwidthAllocation
									Description: "The bandwidth allocation of a queue resource.",
									Type:        types.Float64Type,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"progressive_dialer_config": {
						// Property: ProgressiveDialerConfig
						Description: "Progressive Dialer config",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bandwidth_allocation": {
									// Property: BandwidthAllocation
									Description: "The bandwidth allocation of a queue resource.",
									Type:        types.Float64Type,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Connect Campaign Name",
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Amazon Connect Campaign Name",
			Type:        types.StringType,
			Computed:    true,
		},
		"outbound_call_config": {
			// Property: OutboundCallConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration used for outbound calls.",
			//   "properties": {
			//     "ConnectContactFlowArn": {
			//       "description": "The identifier of the contact flow for the outbound call.",
			//       "maxLength": 500,
			//       "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
			//       "type": "string"
			//     },
			//     "ConnectQueueArn": {
			//       "description": "The queue for the call. If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.",
			//       "maxLength": 500,
			//       "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/queue/[-a-zA-Z0-9]*$",
			//       "type": "string"
			//     },
			//     "ConnectSourcePhoneNumber": {
			//       "description": "The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.",
			//       "maxLength": 100,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ConnectContactFlowArn",
			//     "ConnectQueueArn"
			//   ],
			//   "type": "object"
			// }
			Description: "The configuration used for outbound calls.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"connect_contact_flow_arn": {
						// Property: ConnectContactFlowArn
						Description: "The identifier of the contact flow for the outbound call.",
						Type:        types.StringType,
						Computed:    true,
					},
					"connect_queue_arn": {
						// Property: ConnectQueueArn
						Description: "The queue for the call. If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.",
						Type:        types.StringType,
						Computed:    true,
					},
					"connect_source_phone_number": {
						// Property: ConnectSourcePhoneNumber
						Description: "The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more tags.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more tags.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ConnectCampaigns::Campaign",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ConnectCampaigns::Campaign").WithTerraformTypeName("awscc_connectcampaigns_campaign")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"bandwidth_allocation":        "BandwidthAllocation",
		"connect_contact_flow_arn":    "ConnectContactFlowArn",
		"connect_instance_arn":        "ConnectInstanceArn",
		"connect_queue_arn":           "ConnectQueueArn",
		"connect_source_phone_number": "ConnectSourcePhoneNumber",
		"dialer_config":               "DialerConfig",
		"key":                         "Key",
		"name":                        "Name",
		"outbound_call_config":        "OutboundCallConfig",
		"predictive_dialer_config":    "PredictiveDialerConfig",
		"progressive_dialer_config":   "ProgressiveDialerConfig",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}