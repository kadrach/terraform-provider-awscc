// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_secretsmanager_secret", secretDataSource)
}

// secretDataSource returns the Terraform awscc_secretsmanager_secret data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecretsManager::Secret resource.
func secretDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Optional) Specifies a user-provided description of the secret.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "(Optional) Specifies a user-provided description of the secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GenerateSecretString
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "(Optional) Specifies text data that you want to encrypt and store in this new version of the secret.",
		//	  "properties": {
		//	    "ExcludeCharacters": {
		//	      "description": "A string that excludes characters in the generated password. By default, all characters from the included sets can be used. The string can be a minimum length of 0 characters and a maximum length of 7168 characters. ",
		//	      "type": "string"
		//	    },
		//	    "ExcludeLowercase": {
		//	      "description": "Specifies the generated password should not include lowercase letters. By default, ecrets Manager disables this parameter, and the generated password can include lowercase False, and the generated password can include lowercase letters.",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludeNumbers": {
		//	      "description": "Specifies that the generated password should exclude digits. By default, Secrets Manager does not enable the parameter, False, and the generated password can include digits.",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludePunctuation": {
		//	      "description": "Specifies that the generated password should not include punctuation characters. The default if you do not include this switch parameter is that punctuation characters can be included. ",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludeUppercase": {
		//	      "description": "Specifies that the generated password should not include uppercase letters. The default behavior is False, and the generated password can include uppercase letters. ",
		//	      "type": "boolean"
		//	    },
		//	    "GenerateStringKey": {
		//	      "description": "The JSON key name used to add the generated password to the JSON structure specified by the SecretStringTemplate parameter. If you specify this parameter, then you must also specify SecretStringTemplate. ",
		//	      "type": "string"
		//	    },
		//	    "IncludeSpace": {
		//	      "description": "Specifies that the generated password can include the space character. By default, Secrets Manager disables this parameter, and the generated password doesn't include space",
		//	      "type": "boolean"
		//	    },
		//	    "PasswordLength": {
		//	      "description": "The desired length of the generated password. The default value if you do not include this parameter is 32 characters. ",
		//	      "type": "integer"
		//	    },
		//	    "RequireEachIncludedType": {
		//	      "description": "Specifies whether the generated password must include at least one of every allowed character type. By default, Secrets Manager enables this parameter, and the generated password includes at least one of every character type.",
		//	      "type": "boolean"
		//	    },
		//	    "SecretStringTemplate": {
		//	      "description": "A properly structured JSON string that the generated password can be added to. If you specify this parameter, then you must also specify GenerateStringKey.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"generate_secret_string": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExcludeCharacters
				"exclude_characters": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A string that excludes characters in the generated password. By default, all characters from the included sets can be used. The string can be a minimum length of 0 characters and a maximum length of 7168 characters. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludeLowercase
				"exclude_lowercase": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies the generated password should not include lowercase letters. By default, ecrets Manager disables this parameter, and the generated password can include lowercase False, and the generated password can include lowercase letters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludeNumbers
				"exclude_numbers": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies that the generated password should exclude digits. By default, Secrets Manager does not enable the parameter, False, and the generated password can include digits.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludePunctuation
				"exclude_punctuation": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies that the generated password should not include punctuation characters. The default if you do not include this switch parameter is that punctuation characters can be included. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludeUppercase
				"exclude_uppercase": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies that the generated password should not include uppercase letters. The default behavior is False, and the generated password can include uppercase letters. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: GenerateStringKey
				"generate_string_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The JSON key name used to add the generated password to the JSON structure specified by the SecretStringTemplate parameter. If you specify this parameter, then you must also specify SecretStringTemplate. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IncludeSpace
				"include_space": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies that the generated password can include the space character. By default, Secrets Manager disables this parameter, and the generated password doesn't include space",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PasswordLength
				"password_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The desired length of the generated password. The default value if you do not include this parameter is 32 characters. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RequireEachIncludedType
				"require_each_included_type": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether the generated password must include at least one of every allowed character type. By default, Secrets Manager enables this parameter, and the generated password includes at least one of every character type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecretStringTemplate
				"secret_string_template": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A properly structured JSON string that the generated password can be added to. If you specify this parameter, then you must also specify GenerateStringKey.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "(Optional) Specifies text data that you want to encrypt and store in this new version of the secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "secret Id, the Arn of the resource.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "secret Id, the Arn of the resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Optional) Specifies the ARN, Key ID, or alias of the AWS KMS customer master key (CMK) used to encrypt the SecretString.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "(Optional) Specifies the ARN, Key ID, or alias of the AWS KMS customer master key (CMK) used to encrypt the SecretString.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The friendly name of the secret. You can use forward slashes in the name to represent a path hierarchy.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The friendly name of the secret. You can use forward slashes in the name to represent a path hierarchy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReplicaRegions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Optional) A list of ReplicaRegion objects. The ReplicaRegion type consists of a Region (required) and the KmsKeyId which can be an ARN, Key ID, or Alias.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A custom type that specifies a Region and the KmsKeyId for a replica secret.",
		//	    "properties": {
		//	      "KmsKeyId": {
		//	        "description": "The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses aws/secretsmanager.",
		//	        "type": "string"
		//	      },
		//	      "Region": {
		//	        "description": "(Optional) A string that represents a Region, for example \"us-east-1\".",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Region"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"replica_regions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: KmsKeyId
					"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses aws/secretsmanager.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Region
					"region": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "(Optional) A string that represents a Region, for example \"us-east-1\".",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "(Optional) A list of ReplicaRegion objects. The ReplicaRegion type consists of a Region (required) and the KmsKeyId which can be an ARN, Key ID, or Alias.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecretString
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Optional) Specifies text data that you want to encrypt and store in this new version of the secret.",
		//	  "type": "string"
		//	}
		"secret_string": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "(Optional) Specifies text data that you want to encrypt and store in this new version of the secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of user-defined tags associated with the secret. Use tags to manage your AWS resources. For additional information about tags, see TagResource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of user-defined tags associated with the secret. Use tags to manage your AWS resources. For additional information about tags, see TagResource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecretsManager::Secret",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecretsManager::Secret").WithTerraformTypeName("awscc_secretsmanager_secret")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                "Description",
		"exclude_characters":         "ExcludeCharacters",
		"exclude_lowercase":          "ExcludeLowercase",
		"exclude_numbers":            "ExcludeNumbers",
		"exclude_punctuation":        "ExcludePunctuation",
		"exclude_uppercase":          "ExcludeUppercase",
		"generate_secret_string":     "GenerateSecretString",
		"generate_string_key":        "GenerateStringKey",
		"id":                         "Id",
		"include_space":              "IncludeSpace",
		"key":                        "Key",
		"kms_key_id":                 "KmsKeyId",
		"name":                       "Name",
		"password_length":            "PasswordLength",
		"region":                     "Region",
		"replica_regions":            "ReplicaRegions",
		"require_each_included_type": "RequireEachIncludedType",
		"secret_string":              "SecretString",
		"secret_string_template":     "SecretStringTemplate",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}