---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_glue_schema_version_metadata Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This resource adds Key-Value metadata to a Schema version of Glue Schema Registry.
---

# awscc_glue_schema_version_metadata (Resource)

This resource adds Key-Value metadata to a Schema version of Glue Schema Registry.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **key** (String) Metadata key
- **schema_version_id** (String) Represents the version ID associated with the schema version.
- **value** (String) Metadata value

### Read-Only

- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_glue_schema_version_metadata.example <resource ID>
```
