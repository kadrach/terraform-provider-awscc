---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_mediapackagev2_channel Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::MediaPackageV2::Channel Resource Type
---

# awscc_mediapackagev2_channel (Resource)

Definition of AWS::MediaPackageV2::Channel Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `channel_group_name` (String)
- `channel_name` (String)
- `description` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `ingest_endpoints` (Attributes List) (see [below for nested schema](#nestedatt--ingest_endpoints))
- `modified_at` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--ingest_endpoints"></a>
### Nested Schema for `ingest_endpoints`

Read-Only:

- `id` (String)
- `url` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_mediapackagev2_channel.example <resource ID>
```
