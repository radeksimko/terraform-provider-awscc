---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_account_audit_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.
---

# awscc_iot_account_audit_configuration (Resource)

Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **account_id** (String) Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
- **audit_check_configurations** (Attributes) Specifies which audit checks are enabled and disabled for this account. (see [below for nested schema](#nestedatt--audit_check_configurations))
- **role_arn** (String) The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.

### Optional

- **audit_notification_target_configurations** (Attributes) Information about the targets to which audit notifications are sent. (see [below for nested schema](#nestedatt--audit_notification_target_configurations))

### Read-Only

- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--audit_check_configurations"></a>
### Nested Schema for `audit_check_configurations`

Required:

- **authenticated_cognito_role_overly_permissive_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--authenticated_cognito_role_overly_permissive_check))
- **ca_certificate_expiring_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--ca_certificate_expiring_check))
- **ca_certificate_key_quality_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--ca_certificate_key_quality_check))
- **conflicting_client_ids_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--conflicting_client_ids_check))
- **device_certificate_expiring_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--device_certificate_expiring_check))
- **device_certificate_key_quality_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--device_certificate_key_quality_check))
- **device_certificate_shared_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--device_certificate_shared_check))
- **iot_policy_overly_permissive_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--iot_policy_overly_permissive_check))
- **iot_role_alias_allows_access_to_unused_services_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--iot_role_alias_allows_access_to_unused_services_check))
- **iot_role_alias_overly_permissive_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--iot_role_alias_overly_permissive_check))
- **logging_disabled_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--logging_disabled_check))
- **revoked_ca_certificate_still_active_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--revoked_ca_certificate_still_active_check))
- **revoked_device_certificate_still_active_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--revoked_device_certificate_still_active_check))
- **unauthenticated_cognito_role_overly_permissive_check** (Attributes) The configuration for a specific audit check. (see [below for nested schema](#nestedatt--audit_check_configurations--unauthenticated_cognito_role_overly_permissive_check))

<a id="nestedatt--audit_check_configurations--authenticated_cognito_role_overly_permissive_check"></a>
### Nested Schema for `audit_check_configurations.authenticated_cognito_role_overly_permissive_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--ca_certificate_expiring_check"></a>
### Nested Schema for `audit_check_configurations.ca_certificate_expiring_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--ca_certificate_key_quality_check"></a>
### Nested Schema for `audit_check_configurations.ca_certificate_key_quality_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--conflicting_client_ids_check"></a>
### Nested Schema for `audit_check_configurations.conflicting_client_ids_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--device_certificate_expiring_check"></a>
### Nested Schema for `audit_check_configurations.device_certificate_expiring_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--device_certificate_key_quality_check"></a>
### Nested Schema for `audit_check_configurations.device_certificate_key_quality_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--device_certificate_shared_check"></a>
### Nested Schema for `audit_check_configurations.device_certificate_shared_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--iot_policy_overly_permissive_check"></a>
### Nested Schema for `audit_check_configurations.iot_policy_overly_permissive_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--iot_role_alias_allows_access_to_unused_services_check"></a>
### Nested Schema for `audit_check_configurations.iot_role_alias_allows_access_to_unused_services_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--iot_role_alias_overly_permissive_check"></a>
### Nested Schema for `audit_check_configurations.iot_role_alias_overly_permissive_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--logging_disabled_check"></a>
### Nested Schema for `audit_check_configurations.logging_disabled_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--revoked_ca_certificate_still_active_check"></a>
### Nested Schema for `audit_check_configurations.revoked_ca_certificate_still_active_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--revoked_device_certificate_still_active_check"></a>
### Nested Schema for `audit_check_configurations.revoked_device_certificate_still_active_check`

Required:

- **enabled** (Boolean) True if the check is enabled.


<a id="nestedatt--audit_check_configurations--unauthenticated_cognito_role_overly_permissive_check"></a>
### Nested Schema for `audit_check_configurations.unauthenticated_cognito_role_overly_permissive_check`

Required:

- **enabled** (Boolean) True if the check is enabled.



<a id="nestedatt--audit_notification_target_configurations"></a>
### Nested Schema for `audit_notification_target_configurations`

Optional:

- **sns** (Attributes) (see [below for nested schema](#nestedatt--audit_notification_target_configurations--sns))

<a id="nestedatt--audit_notification_target_configurations--sns"></a>
### Nested Schema for `audit_notification_target_configurations.sns`

Optional:

- **enabled** (Boolean) True if notifications to the target are enabled.
- **role_arn** (String) The ARN of the role that grants permission to send notifications to the target.
- **target_arn** (String) The ARN of the target (SNS topic) to which audit notifications are sent.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iot_account_audit_configuration.example <resource ID>
```
