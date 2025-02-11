---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53recoverycontrol_cluster Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Route53RecoveryControl::Cluster
---

# awscc_route53recoverycontrol_cluster (Data Source)

Data Source schema for AWS::Route53RecoveryControl::Cluster



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **cluster_arn** (String) The Amazon Resource Name (ARN) of the cluster.
- **cluster_endpoints** (Attributes List) Endpoints for the cluster. (see [below for nested schema](#nestedatt--cluster_endpoints))
- **name** (String) Name of a Cluster. You can use any non-white space character in the name
- **status** (String) Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.

<a id="nestedatt--cluster_endpoints"></a>
### Nested Schema for `cluster_endpoints`

Read-Only:

- **endpoint** (String)
- **region** (String)


