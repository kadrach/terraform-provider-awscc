---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ivs_recording_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IVS::RecordingConfiguration
---

# awscc_ivs_recording_configuration (Resource)

Resource Type definition for AWS::IVS::RecordingConfiguration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **destination_configuration** (Attributes) Recording Destination Configuration. (see [below for nested schema](#nestedatt--destination_configuration))

### Optional

- **name** (String) Recording Configuration Name.
- **tags** (Attributes Set) A list of key-value pairs that contain metadata for the asset model. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
- **id** (String) Uniquely identifies the resource.
- **state** (String) Recording Configuration State.

<a id="nestedatt--destination_configuration"></a>
### Nested Schema for `destination_configuration`

Required:

- **s3** (Attributes) Recording S3 Destination Configuration. (see [below for nested schema](#nestedatt--destination_configuration--s3))

<a id="nestedatt--destination_configuration--s3"></a>
### Nested Schema for `destination_configuration.s3`

Required:

- **bucket_name** (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ivs_recording_configuration.example <resource ID>
```