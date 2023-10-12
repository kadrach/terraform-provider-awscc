// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_vpc_endpoint_connection_notification", vPCEndpointConnectionNotificationDataSource)
}

// vPCEndpointConnectionNotificationDataSource returns the Terraform awscc_ec2_vpc_endpoint_connection_notification data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::VPCEndpointConnectionNotification resource.
func vPCEndpointConnectionNotificationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConnectionEvents
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The endpoint events for which to receive notifications.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"connection_events": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The endpoint events for which to receive notifications.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectionNotificationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the SNS topic for the notifications.",
		//	  "type": "string"
		//	}
		"connection_notification_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the SNS topic for the notifications.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the endpoint service.",
		//	  "type": "string"
		//	}
		"service_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the endpoint service.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VPCEndpointConnectionNotificationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC Endpoint Connection ID generated by service",
		//	  "type": "string"
		//	}
		"vpc_endpoint_connection_notification_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "VPC Endpoint Connection ID generated by service",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VPCEndpointId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the endpoint.",
		//	  "type": "string"
		//	}
		"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::VPCEndpointConnectionNotification",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPCEndpointConnectionNotification").WithTerraformTypeName("awscc_ec2_vpc_endpoint_connection_notification")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connection_events":                       "ConnectionEvents",
		"connection_notification_arn":             "ConnectionNotificationArn",
		"service_id":                              "ServiceId",
		"vpc_endpoint_connection_notification_id": "VPCEndpointConnectionNotificationId",
		"vpc_endpoint_id":                         "VPCEndpointId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}