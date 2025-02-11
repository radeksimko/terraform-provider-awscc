---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_model_package_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SageMaker::ModelPackageGroup
---

# awscc_sagemaker_model_package_group (Resource)

Resource Type definition for AWS::SageMaker::ModelPackageGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **model_package_group_name** (String) The name of the model package group.

### Optional

- **model_package_group_description** (String) The description of the model package group.
- **model_package_group_policy** (String)
- **tags** (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **creation_time** (String) The time at which the model package group was created.
- **id** (String) Uniquely identifies the resource.
- **model_package_group_arn** (String) The Amazon Resource Name (ARN) of the model package group.
- **model_package_group_status** (String) The status of a modelpackage group job.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_sagemaker_model_package_group.example <resource ID>
```
