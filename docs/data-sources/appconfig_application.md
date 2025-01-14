---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appconfig_application Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::AppConfig::Application
---

# awscc_appconfig_application (Data Source)

Data Source schema for AWS::AppConfig::Application



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `application_id` (String) The application Id
- `description` (String) A description of the application.
- `name` (String) A name for the application.
- `tags` (Attributes Set) Metadata to assign to the application. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.
- `value` (String) The tag value can be up to 256 characters.
