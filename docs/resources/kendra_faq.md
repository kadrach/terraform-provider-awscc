---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_kendra_faq Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A Kendra FAQ resource
---

# awscc_kendra_faq (Resource)

A Kendra FAQ resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **index_id** (String) Unique ID of Index
- **name** (String)
- **role_arn** (String)
- **s3_path** (Attributes) (see [below for nested schema](#nestedatt--s3_path))

### Optional

- **description** (String) Description of the FAQ
- **file_format** (String) Format of the input file
- **tags** (Attributes List) List of tags (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String)
- **id** (String) Unique ID of the FAQ

<a id="nestedatt--s3_path"></a>
### Nested Schema for `s3_path`

Required:

- **bucket** (String)
- **key** (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) A string used to identify this tag
- **value** (String) A string containing the value for the tag

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_kendra_faq.example <resource ID>
```