---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_greengrassv2_component_version Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource for Greengrass component version.
---

# awscc_greengrassv2_component_version (Resource)

Resource for Greengrass component version.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **inline_recipe** (String)
- **lambda_function** (Attributes) (see [below for nested schema](#nestedatt--lambda_function))
- **tags** (Map of String)

### Read-Only

- **arn** (String)
- **component_name** (String)
- **component_version** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--lambda_function"></a>
### Nested Schema for `lambda_function`

Optional:

- **component_dependencies** (Attributes Map) (see [below for nested schema](#nestedatt--lambda_function--component_dependencies))
- **component_lambda_parameters** (Attributes) (see [below for nested schema](#nestedatt--lambda_function--component_lambda_parameters))
- **component_name** (String)
- **component_platforms** (Attributes List) (see [below for nested schema](#nestedatt--lambda_function--component_platforms))
- **component_version** (String)
- **lambda_arn** (String)

<a id="nestedatt--lambda_function--component_dependencies"></a>
### Nested Schema for `lambda_function.component_dependencies`

Optional:

- **dependency_type** (String)
- **version_requirement** (String)


<a id="nestedatt--lambda_function--component_lambda_parameters"></a>
### Nested Schema for `lambda_function.component_lambda_parameters`

Optional:

- **environment_variables** (Map of String)
- **event_sources** (Attributes List) (see [below for nested schema](#nestedatt--lambda_function--component_lambda_parameters--event_sources))
- **exec_args** (List of String)
- **input_payload_encoding_type** (String)
- **linux_process_params** (Attributes) (see [below for nested schema](#nestedatt--lambda_function--component_lambda_parameters--linux_process_params))
- **max_idle_time_in_seconds** (Number)
- **max_instances_count** (Number)
- **max_queue_size** (Number)
- **pinned** (Boolean)
- **status_timeout_in_seconds** (Number)
- **timeout_in_seconds** (Number)

<a id="nestedatt--lambda_function--component_lambda_parameters--event_sources"></a>
### Nested Schema for `lambda_function.component_lambda_parameters.event_sources`

Optional:

- **topic** (String)
- **type** (String)


<a id="nestedatt--lambda_function--component_lambda_parameters--linux_process_params"></a>
### Nested Schema for `lambda_function.component_lambda_parameters.linux_process_params`

Optional:

- **container_params** (Attributes) (see [below for nested schema](#nestedatt--lambda_function--component_lambda_parameters--linux_process_params--container_params))
- **isolation_mode** (String)

<a id="nestedatt--lambda_function--component_lambda_parameters--linux_process_params--container_params"></a>
### Nested Schema for `lambda_function.component_lambda_parameters.linux_process_params.isolation_mode`

Optional:

- **devices** (Attributes List) (see [below for nested schema](#nestedatt--lambda_function--component_lambda_parameters--linux_process_params--isolation_mode--devices))
- **memory_size_in_kb** (Number)
- **mount_ro_sysfs** (Boolean)
- **volumes** (Attributes List) (see [below for nested schema](#nestedatt--lambda_function--component_lambda_parameters--linux_process_params--isolation_mode--volumes))

<a id="nestedatt--lambda_function--component_lambda_parameters--linux_process_params--isolation_mode--devices"></a>
### Nested Schema for `lambda_function.component_lambda_parameters.linux_process_params.isolation_mode.devices`

Optional:

- **add_group_owner** (Boolean)
- **path** (String)
- **permission** (String)


<a id="nestedatt--lambda_function--component_lambda_parameters--linux_process_params--isolation_mode--volumes"></a>
### Nested Schema for `lambda_function.component_lambda_parameters.linux_process_params.isolation_mode.volumes`

Optional:

- **add_group_owner** (Boolean)
- **destination_path** (String)
- **permission** (String)
- **source_path** (String)





<a id="nestedatt--lambda_function--component_platforms"></a>
### Nested Schema for `lambda_function.component_platforms`

Optional:

- **attributes** (Map of String)
- **name** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_greengrassv2_component_version.example <resource ID>
```
