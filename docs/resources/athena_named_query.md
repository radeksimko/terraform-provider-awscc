---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_athena_named_query Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Athena::NamedQuery
---

# awscc_athena_named_query (Resource)

Resource schema for AWS::Athena::NamedQuery



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **database** (String) The database to which the query belongs.
- **query_string** (String) The contents of the query with all query statements.

### Optional

- **description** (String) The query description.
- **name** (String) The query name.
- **work_group** (String) The name of the workgroup that contains the named query.

### Read-Only

- **id** (String) Uniquely identifies the resource.
- **named_query_id** (String) The unique ID of the query.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_athena_named_query.example <resource ID>
```
