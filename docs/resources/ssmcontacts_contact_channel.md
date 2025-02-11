---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ssmcontacts_contact_channel Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SSMContacts::ContactChannel
---

# awscc_ssmcontacts_contact_channel (Resource)

Resource Type definition for AWS::SSMContacts::ContactChannel



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **channel_address** (String) The details that SSM Incident Manager uses when trying to engage the contact channel.
- **channel_name** (String) The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
- **channel_type** (String) Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.
- **contact_id** (String) ARN of the contact resource
- **defer_activation** (Boolean) If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the engagement to a contact channel.
- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ssmcontacts_contact_channel.example <resource ID>
```
