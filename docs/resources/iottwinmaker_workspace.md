---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iottwinmaker_workspace Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::IoTTwinMaker::Workspace
---

# awscc_iottwinmaker_workspace (Resource)

Resource schema for AWS::IoTTwinMaker::Workspace



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `role` (String) The ARN of the execution role associated with the workspace.
- `s3_location` (String) The ARN of the S3 bucket where resources associated with the workspace are stored.
- `workspace_id` (String) The ID of the workspace.

### Optional

- `description` (String) The description of the workspace.
- `tags` (Map of String) A map of key-value pairs to associate with a resource.

### Read-Only

- `arn` (String) The ARN of the workspace.
- `creation_date_time` (String) The date and time when the workspace was created.
- `id` (String) Uniquely identifies the resource.
- `update_date_time` (String) The date and time of the current update.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iottwinmaker_workspace.example <resource ID>
```