---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_wafv2_regex_pattern_set Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::WAFv2::RegexPatternSet
---

# awscc_wafv2_regex_pattern_set (Data Source)

Data Source schema for AWS::WAFv2::RegexPatternSet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) ARN of the WAF entity.
- **description** (String) Description of the entity.
- **name** (String) Name of the RegexPatternSet.
- **regular_expression_list** (List of String)
- **scope** (String) Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)


