---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_wafv2_regex_pattern_set Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Contains a list of Regular expressions based on the provided inputs. RegexPatternSet can be used with other WAF entities with RegexPatternSetReferenceStatement to perform other actions .
---

# awscc_wafv2_regex_pattern_set (Resource)

Contains a list of Regular expressions based on the provided inputs. RegexPatternSet can be used with other WAF entities with RegexPatternSetReferenceStatement to perform other actions .



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **regular_expression_list** (List of String)
- **scope** (String) Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.

### Optional

- **description** (String) Description of the entity.
- **name** (String) Name of the RegexPatternSet.
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) ARN of the WAF entity.
- **id** (String) Id of the RegexPatternSet

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_wafv2_regex_pattern_set.example <resource ID>
```
