---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_servicecatalog_service_action_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Schema for AWS::ServiceCatalog::ServiceActionAssociation
---

# awscc_servicecatalog_service_action_association (Resource)

Resource Schema for AWS::ServiceCatalog::ServiceActionAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **product_id** (String)
- **provisioning_artifact_id** (String)
- **service_action_id** (String)

### Read-Only

- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_servicecatalog_service_action_association.example <resource ID>
```
