---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_glue_schema_version Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This resource represents an individual schema version of a schema defined in Glue Schema Registry.
---

# awscc_glue_schema_version (Resource)

This resource represents an individual schema version of a schema defined in Glue Schema Registry.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **schema** (Attributes) Identifier for the schema where the schema version will be created. (see [below for nested schema](#nestedatt--schema))
- **schema_definition** (String) Complete definition of the schema in plain-text.

### Read-Only

- **id** (String) Uniquely identifies the resource.
- **version_id** (String) Represents the version ID associated with the schema version.

<a id="nestedatt--schema"></a>
### Nested Schema for `schema`

Required:

- **registry_name** (String) Name of the registry to identify where the Schema is located.
- **schema_arn** (String) Amazon Resource Name for the Schema. This attribute can be used to uniquely represent the Schema.
- **schema_name** (String) Name of the schema. This parameter requires RegistryName to be provided.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_glue_schema_version.example <resource ID>
```