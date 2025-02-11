---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_fms_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Creates an AWS Firewall Manager policy.
---

# awscc_fms_policy (Resource)

Creates an AWS Firewall Manager policy.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **exclude_resource_tags** (Boolean)
- **policy_name** (String)
- **remediation_enabled** (Boolean)
- **resource_type** (String) An AWS resource type
- **security_service_policy_data** (Attributes) (see [below for nested schema](#nestedatt--security_service_policy_data))

### Optional

- **delete_all_policy_resources** (Boolean)
- **exclude_map** (Attributes) An FMS includeMap or excludeMap. (see [below for nested schema](#nestedatt--exclude_map))
- **include_map** (Attributes) An FMS includeMap or excludeMap. (see [below for nested schema](#nestedatt--include_map))
- **resource_tags** (Attributes List) (see [below for nested schema](#nestedatt--resource_tags))
- **resource_type_list** (List of String)
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) A resource ARN.
- **id** (String) The ID of this resource.

<a id="nestedatt--security_service_policy_data"></a>
### Nested Schema for `security_service_policy_data`

Required:

- **managed_service_data** (String)
- **type** (String)


<a id="nestedatt--exclude_map"></a>
### Nested Schema for `exclude_map`

Optional:

- **account** (List of String)
- **orgunit** (List of String)


<a id="nestedatt--include_map"></a>
### Nested Schema for `include_map`

Optional:

- **account** (List of String)
- **orgunit** (List of String)


<a id="nestedatt--resource_tags"></a>
### Nested Schema for `resource_tags`

Optional:

- **key** (String)
- **value** (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_fms_policy.example <resource ID>
```
