// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudformation_publisher", publisherDataSourceType)
}

// publisherDataSourceType returns the Terraform awscc_cloudformation_publisher data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudFormation::Publisher resource type.
func publisherDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"accept_terms_and_conditions": {
			// Property: AcceptTermsAndConditions
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf",
			//   "type": "boolean"
			// }
			Description: "Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf",
			Type:        types.BoolType,
			Computed:    true,
		},
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"identity_provider": {
			// Property: IdentityProvider
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of account used as the identity provider when registering this publisher with CloudFormation.",
			//   "enum": [
			//     "AWS_Marketplace",
			//     "GitHub",
			//     "Bitbucket"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of account used as the identity provider when registering this publisher with CloudFormation.",
			Type:        types.StringType,
			Computed:    true,
		},
		"publisher_id": {
			// Property: PublisherId
			// CloudFormation resource type schema:
			// {
			//   "description": "The publisher id assigned by CloudFormation for publishing in this region.",
			//   "maxLength": 40,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The publisher id assigned by CloudFormation for publishing in this region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"publisher_profile": {
			// Property: PublisherProfile
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL to the publisher's profile with the identity provider.",
			//   "maxLength": 1024,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL to the publisher's profile with the identity provider.",
			Type:        types.StringType,
			Computed:    true,
		},
		"publisher_status": {
			// Property: PublisherStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the publisher is verified.",
			//   "enum": [
			//     "VERIFIED",
			//     "UNVERIFIED"
			//   ],
			//   "type": "string"
			// }
			Description: "Whether the publisher is verified.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CloudFormation::Publisher",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::Publisher").WithTerraformTypeName("awscc_cloudformation_publisher")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accept_terms_and_conditions": "AcceptTermsAndConditions",
		"connection_arn":              "ConnectionArn",
		"identity_provider":           "IdentityProvider",
		"publisher_id":                "PublisherId",
		"publisher_profile":           "PublisherProfile",
		"publisher_status":            "PublisherStatus",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}