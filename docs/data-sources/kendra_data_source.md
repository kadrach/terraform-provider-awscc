---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_kendra_data_source Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Kendra::DataSource
---

# awscc_kendra_data_source (Data Source)

Data Source schema for AWS::Kendra::DataSource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String)
- **data_source_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration))
- **description** (String) Description of data source
- **index_id** (String) ID of Index
- **name** (String) Name of data source
- **role_arn** (String) Role ARN
- **schedule** (String) Schedule
- **tags** (Attributes List) List of tags (see [below for nested schema](#nestedatt--tags))
- **type** (String) Data source type

<a id="nestedatt--data_source_configuration"></a>
### Nested Schema for `data_source_configuration`

Read-Only:

- **confluence_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration))
- **database_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--database_configuration))
- **google_drive_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--google_drive_configuration))
- **one_drive_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--one_drive_configuration))
- **s3_configuration** (Attributes) S3 data source configuration (see [below for nested schema](#nestedatt--data_source_configuration--s3_configuration))
- **salesforce_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration))
- **service_now_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--service_now_configuration))
- **share_point_configuration** (Attributes) SharePoint configuration (see [below for nested schema](#nestedatt--data_source_configuration--share_point_configuration))
- **web_crawler_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--web_crawler_configuration))
- **work_docs_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--work_docs_configuration))

<a id="nestedatt--data_source_configuration--confluence_configuration"></a>
### Nested Schema for `data_source_configuration.confluence_configuration`

Read-Only:

- **attachment_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--attachment_configuration))
- **blog_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--blog_configuration))
- **exclusion_patterns** (List of String)
- **inclusion_patterns** (List of String)
- **page_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--page_configuration))
- **secret_arn** (String)
- **server_url** (String)
- **space_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--space_configuration))
- **version** (String)
- **vpc_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--vpc_configuration))

<a id="nestedatt--data_source_configuration--confluence_configuration--attachment_configuration"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.attachment_configuration`

Read-Only:

- **attachment_field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--attachment_configuration--attachment_field_mappings))
- **crawl_attachments** (Boolean)

<a id="nestedatt--data_source_configuration--confluence_configuration--attachment_configuration--attachment_field_mappings"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.attachment_configuration.crawl_attachments`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--confluence_configuration--blog_configuration"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.blog_configuration`

Read-Only:

- **blog_field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--blog_configuration--blog_field_mappings))

<a id="nestedatt--data_source_configuration--confluence_configuration--blog_configuration--blog_field_mappings"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.blog_configuration.blog_field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--confluence_configuration--page_configuration"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.page_configuration`

Read-Only:

- **page_field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--page_configuration--page_field_mappings))

<a id="nestedatt--data_source_configuration--confluence_configuration--page_configuration--page_field_mappings"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.page_configuration.page_field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--confluence_configuration--space_configuration"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.space_configuration`

Read-Only:

- **crawl_archived_spaces** (Boolean)
- **crawl_personal_spaces** (Boolean)
- **exclude_spaces** (List of String)
- **include_spaces** (List of String)
- **space_field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--confluence_configuration--space_configuration--space_field_mappings))

<a id="nestedatt--data_source_configuration--confluence_configuration--space_configuration--space_field_mappings"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.space_configuration.space_field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--confluence_configuration--vpc_configuration"></a>
### Nested Schema for `data_source_configuration.confluence_configuration.vpc_configuration`

Read-Only:

- **security_group_ids** (List of String)
- **subnet_ids** (List of String)



<a id="nestedatt--data_source_configuration--database_configuration"></a>
### Nested Schema for `data_source_configuration.database_configuration`

Read-Only:

- **acl_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--database_configuration--acl_configuration))
- **column_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--database_configuration--column_configuration))
- **connection_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--database_configuration--connection_configuration))
- **database_engine_type** (String)
- **sql_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--database_configuration--sql_configuration))
- **vpc_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--database_configuration--vpc_configuration))

<a id="nestedatt--data_source_configuration--database_configuration--acl_configuration"></a>
### Nested Schema for `data_source_configuration.database_configuration.acl_configuration`

Read-Only:

- **allowed_groups_column_name** (String)


<a id="nestedatt--data_source_configuration--database_configuration--column_configuration"></a>
### Nested Schema for `data_source_configuration.database_configuration.column_configuration`

Read-Only:

- **change_detecting_columns** (List of String)
- **document_data_column_name** (String)
- **document_id_column_name** (String)
- **document_title_column_name** (String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--database_configuration--column_configuration--field_mappings))

<a id="nestedatt--data_source_configuration--database_configuration--column_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.database_configuration.column_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--database_configuration--connection_configuration"></a>
### Nested Schema for `data_source_configuration.database_configuration.connection_configuration`

Read-Only:

- **database_host** (String)
- **database_name** (String)
- **database_port** (Number)
- **secret_arn** (String)
- **table_name** (String)


<a id="nestedatt--data_source_configuration--database_configuration--sql_configuration"></a>
### Nested Schema for `data_source_configuration.database_configuration.sql_configuration`

Read-Only:

- **query_identifiers_enclosing_option** (String)


<a id="nestedatt--data_source_configuration--database_configuration--vpc_configuration"></a>
### Nested Schema for `data_source_configuration.database_configuration.vpc_configuration`

Read-Only:

- **security_group_ids** (List of String)
- **subnet_ids** (List of String)



<a id="nestedatt--data_source_configuration--google_drive_configuration"></a>
### Nested Schema for `data_source_configuration.google_drive_configuration`

Read-Only:

- **exclude_mime_types** (List of String)
- **exclude_shared_drives** (List of String)
- **exclude_user_accounts** (List of String)
- **exclusion_patterns** (List of String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--google_drive_configuration--field_mappings))
- **inclusion_patterns** (List of String)
- **secret_arn** (String)

<a id="nestedatt--data_source_configuration--google_drive_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.google_drive_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--one_drive_configuration"></a>
### Nested Schema for `data_source_configuration.one_drive_configuration`

Read-Only:

- **disable_local_groups** (Boolean)
- **exclusion_patterns** (List of String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--one_drive_configuration--field_mappings))
- **inclusion_patterns** (List of String)
- **one_drive_users** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--one_drive_configuration--one_drive_users))
- **secret_arn** (String)
- **tenant_domain** (String)

<a id="nestedatt--data_source_configuration--one_drive_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.one_drive_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)


<a id="nestedatt--data_source_configuration--one_drive_configuration--one_drive_users"></a>
### Nested Schema for `data_source_configuration.one_drive_configuration.one_drive_users`

Read-Only:

- **one_drive_user_list** (List of String)
- **one_drive_user_s3_path** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--one_drive_configuration--one_drive_users--one_drive_user_s3_path))

<a id="nestedatt--data_source_configuration--one_drive_configuration--one_drive_users--one_drive_user_s3_path"></a>
### Nested Schema for `data_source_configuration.one_drive_configuration.one_drive_users.one_drive_user_s3_path`

Read-Only:

- **bucket** (String)
- **key** (String)




<a id="nestedatt--data_source_configuration--s3_configuration"></a>
### Nested Schema for `data_source_configuration.s3_configuration`

Read-Only:

- **access_control_list_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--s3_configuration--access_control_list_configuration))
- **bucket_name** (String)
- **documents_metadata_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--s3_configuration--documents_metadata_configuration))
- **exclusion_patterns** (List of String)
- **inclusion_patterns** (List of String)
- **inclusion_prefixes** (List of String)

<a id="nestedatt--data_source_configuration--s3_configuration--access_control_list_configuration"></a>
### Nested Schema for `data_source_configuration.s3_configuration.access_control_list_configuration`

Read-Only:

- **key_path** (String)


<a id="nestedatt--data_source_configuration--s3_configuration--documents_metadata_configuration"></a>
### Nested Schema for `data_source_configuration.s3_configuration.documents_metadata_configuration`

Read-Only:

- **s3_prefix** (String)



<a id="nestedatt--data_source_configuration--salesforce_configuration"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration`

Read-Only:

- **chatter_feed_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--chatter_feed_configuration))
- **crawl_attachments** (Boolean)
- **exclude_attachment_file_patterns** (List of String)
- **include_attachment_file_patterns** (List of String)
- **knowledge_article_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration))
- **secret_arn** (String)
- **server_url** (String)
- **standard_object_attachment_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--standard_object_attachment_configuration))
- **standard_object_configurations** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--standard_object_configurations))

<a id="nestedatt--data_source_configuration--salesforce_configuration--chatter_feed_configuration"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.chatter_feed_configuration`

Read-Only:

- **document_data_field_name** (String)
- **document_title_field_name** (String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--chatter_feed_configuration--field_mappings))
- **include_filter_types** (List of String)

<a id="nestedatt--data_source_configuration--salesforce_configuration--chatter_feed_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.chatter_feed_configuration.include_filter_types`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.knowledge_article_configuration`

Read-Only:

- **custom_knowledge_article_type_configurations** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--custom_knowledge_article_type_configurations))
- **included_states** (List of String)
- **standard_knowledge_article_type_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--standard_knowledge_article_type_configuration))

<a id="nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--custom_knowledge_article_type_configurations"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.knowledge_article_configuration.standard_knowledge_article_type_configuration`

Read-Only:

- **document_data_field_name** (String)
- **document_title_field_name** (String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--standard_knowledge_article_type_configuration--field_mappings))
- **name** (String)

<a id="nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--standard_knowledge_article_type_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.knowledge_article_configuration.standard_knowledge_article_type_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--standard_knowledge_article_type_configuration"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.knowledge_article_configuration.standard_knowledge_article_type_configuration`

Read-Only:

- **document_data_field_name** (String)
- **document_title_field_name** (String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--standard_knowledge_article_type_configuration--field_mappings))

<a id="nestedatt--data_source_configuration--salesforce_configuration--knowledge_article_configuration--standard_knowledge_article_type_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.knowledge_article_configuration.standard_knowledge_article_type_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)




<a id="nestedatt--data_source_configuration--salesforce_configuration--standard_object_attachment_configuration"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.standard_object_attachment_configuration`

Read-Only:

- **document_title_field_name** (String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--standard_object_attachment_configuration--field_mappings))

<a id="nestedatt--data_source_configuration--salesforce_configuration--standard_object_attachment_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.standard_object_attachment_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--salesforce_configuration--standard_object_configurations"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.standard_object_configurations`

Read-Only:

- **document_data_field_name** (String)
- **document_title_field_name** (String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--salesforce_configuration--standard_object_configurations--field_mappings))
- **name** (String)

<a id="nestedatt--data_source_configuration--salesforce_configuration--standard_object_configurations--field_mappings"></a>
### Nested Schema for `data_source_configuration.salesforce_configuration.standard_object_configurations.name`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)




<a id="nestedatt--data_source_configuration--service_now_configuration"></a>
### Nested Schema for `data_source_configuration.service_now_configuration`

Read-Only:

- **authentication_type** (String)
- **host_url** (String)
- **knowledge_article_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--service_now_configuration--knowledge_article_configuration))
- **secret_arn** (String)
- **service_catalog_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--service_now_configuration--service_catalog_configuration))
- **service_now_build_version** (String)

<a id="nestedatt--data_source_configuration--service_now_configuration--knowledge_article_configuration"></a>
### Nested Schema for `data_source_configuration.service_now_configuration.knowledge_article_configuration`

Read-Only:

- **crawl_attachments** (Boolean)
- **document_data_field_name** (String)
- **document_title_field_name** (String)
- **exclude_attachment_file_patterns** (List of String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--service_now_configuration--knowledge_article_configuration--field_mappings))
- **filter_query** (String)
- **include_attachment_file_patterns** (List of String)

<a id="nestedatt--data_source_configuration--service_now_configuration--knowledge_article_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.service_now_configuration.knowledge_article_configuration.include_attachment_file_patterns`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)



<a id="nestedatt--data_source_configuration--service_now_configuration--service_catalog_configuration"></a>
### Nested Schema for `data_source_configuration.service_now_configuration.service_catalog_configuration`

Read-Only:

- **crawl_attachments** (Boolean)
- **document_data_field_name** (String)
- **document_title_field_name** (String)
- **exclude_attachment_file_patterns** (List of String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--service_now_configuration--service_catalog_configuration--field_mappings))
- **include_attachment_file_patterns** (List of String)

<a id="nestedatt--data_source_configuration--service_now_configuration--service_catalog_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.service_now_configuration.service_catalog_configuration.include_attachment_file_patterns`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)




<a id="nestedatt--data_source_configuration--share_point_configuration"></a>
### Nested Schema for `data_source_configuration.share_point_configuration`

Read-Only:

- **crawl_attachments** (Boolean)
- **disable_local_groups** (Boolean)
- **document_title_field_name** (String)
- **exclusion_patterns** (List of String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--share_point_configuration--field_mappings))
- **inclusion_patterns** (List of String)
- **secret_arn** (String)
- **share_point_version** (String)
- **ssl_certificate_s3_path** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--share_point_configuration--ssl_certificate_s3_path))
- **urls** (List of String)
- **use_change_log** (Boolean)
- **vpc_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--share_point_configuration--vpc_configuration))

<a id="nestedatt--data_source_configuration--share_point_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.share_point_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)


<a id="nestedatt--data_source_configuration--share_point_configuration--ssl_certificate_s3_path"></a>
### Nested Schema for `data_source_configuration.share_point_configuration.ssl_certificate_s3_path`

Read-Only:

- **bucket** (String)
- **key** (String)


<a id="nestedatt--data_source_configuration--share_point_configuration--vpc_configuration"></a>
### Nested Schema for `data_source_configuration.share_point_configuration.vpc_configuration`

Read-Only:

- **security_group_ids** (List of String)
- **subnet_ids** (List of String)



<a id="nestedatt--data_source_configuration--web_crawler_configuration"></a>
### Nested Schema for `data_source_configuration.web_crawler_configuration`

Read-Only:

- **authentication_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--web_crawler_configuration--authentication_configuration))
- **crawl_depth** (Number)
- **max_content_size_per_page_in_mega_bytes** (Number)
- **max_links_per_page** (Number)
- **max_urls_per_minute_crawl_rate** (Number)
- **proxy_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--web_crawler_configuration--proxy_configuration))
- **url_exclusion_patterns** (List of String)
- **url_inclusion_patterns** (List of String)
- **urls** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--web_crawler_configuration--urls))

<a id="nestedatt--data_source_configuration--web_crawler_configuration--authentication_configuration"></a>
### Nested Schema for `data_source_configuration.web_crawler_configuration.authentication_configuration`

Read-Only:

- **basic_authentication** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--web_crawler_configuration--authentication_configuration--basic_authentication))

<a id="nestedatt--data_source_configuration--web_crawler_configuration--authentication_configuration--basic_authentication"></a>
### Nested Schema for `data_source_configuration.web_crawler_configuration.authentication_configuration.basic_authentication`

Read-Only:

- **credentials** (String)
- **host** (String)
- **port** (Number)



<a id="nestedatt--data_source_configuration--web_crawler_configuration--proxy_configuration"></a>
### Nested Schema for `data_source_configuration.web_crawler_configuration.proxy_configuration`

Read-Only:

- **credentials** (String)
- **host** (String)
- **port** (Number)


<a id="nestedatt--data_source_configuration--web_crawler_configuration--urls"></a>
### Nested Schema for `data_source_configuration.web_crawler_configuration.urls`

Read-Only:

- **seed_url_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--web_crawler_configuration--urls--seed_url_configuration))
- **site_maps_configuration** (Attributes) (see [below for nested schema](#nestedatt--data_source_configuration--web_crawler_configuration--urls--site_maps_configuration))

<a id="nestedatt--data_source_configuration--web_crawler_configuration--urls--seed_url_configuration"></a>
### Nested Schema for `data_source_configuration.web_crawler_configuration.urls.site_maps_configuration`

Read-Only:

- **seed_urls** (List of String)
- **web_crawler_mode** (String)


<a id="nestedatt--data_source_configuration--web_crawler_configuration--urls--site_maps_configuration"></a>
### Nested Schema for `data_source_configuration.web_crawler_configuration.urls.site_maps_configuration`

Read-Only:

- **site_maps** (List of String)




<a id="nestedatt--data_source_configuration--work_docs_configuration"></a>
### Nested Schema for `data_source_configuration.work_docs_configuration`

Read-Only:

- **crawl_comments** (Boolean)
- **exclusion_patterns** (List of String)
- **field_mappings** (Attributes List) (see [below for nested schema](#nestedatt--data_source_configuration--work_docs_configuration--field_mappings))
- **inclusion_patterns** (List of String)
- **organization_id** (String)
- **use_change_log** (Boolean)

<a id="nestedatt--data_source_configuration--work_docs_configuration--field_mappings"></a>
### Nested Schema for `data_source_configuration.work_docs_configuration.field_mappings`

Read-Only:

- **data_source_field_name** (String)
- **date_field_format** (String)
- **index_field_name** (String)




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) A string used to identify this tag
- **value** (String) A string containing the value for the tag

