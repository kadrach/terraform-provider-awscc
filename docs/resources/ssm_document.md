---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ssm_document Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.
---

# awscc_ssm_document (Resource)

The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **content** (String) The content for the Systems Manager document in JSON, YAML or String format.

### Optional

- **attachments** (Attributes List) A list of key and value pairs that describe attachments to a version of a document. (see [below for nested schema](#nestedatt--attachments))
- **document_format** (String) Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.
- **document_type** (String) The type of document to create.
- **name** (String) A name for the Systems Manager document.
- **requires** (Attributes List) A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document. (see [below for nested schema](#nestedatt--requires))
- **tags** (Attributes List) Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. (see [below for nested schema](#nestedatt--tags))
- **target_type** (String) Specify a target type to define the kinds of resources the document can run on.
- **version_name** (String) An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.

### Read-Only

- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--attachments"></a>
### Nested Schema for `attachments`

Optional:

- **key** (String) The key of a key-value pair that identifies the location of an attachment to a document.
- **name** (String) The name of the document attachment file.
- **values** (List of String) The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.


<a id="nestedatt--requires"></a>
### Nested Schema for `requires`

Optional:

- **name** (String) The name of the required SSM document. The name can be an Amazon Resource Name (ARN).
- **version** (String) The document version required by the current document.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The name of the tag.
- **value** (String) The value of the tag.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ssm_document.example <resource ID>
```