// Code generated by generators/resource/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_macie_custom_data_identifier", customDataIdentifierResourceType)
}

// customDataIdentifierResourceType returns the Terraform awscc_macie_custom_data_identifier resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Macie::CustomDataIdentifier resource type.
func customDataIdentifierResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Custom data identifier ARN.",
			//   "type": "string"
			// }
			Description: "Custom data identifier ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of custom data identifier.",
			//   "type": "string"
			// }
			Description: "Description of custom data identifier.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Custom data identifier ID.",
			//   "type": "string"
			// }
			Description: "Custom data identifier ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ignore_words": {
			// Property: IgnoreWords
			// CloudFormation resource type schema:
			// {
			//   "description": "Words to be ignored.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Words to be ignored.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"keywords": {
			// Property: Keywords
			// CloudFormation resource type schema:
			// {
			//   "description": "Keywords to be matched against.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Keywords to be matched against.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"maximum_match_distance": {
			// Property: MaximumMatchDistance
			// CloudFormation resource type schema:
			// {
			//   "description": "Maximum match distance.",
			//   "type": "integer"
			// }
			Description: "Maximum match distance.",
			Type:        types.NumberType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of custom data identifier.",
			//   "type": "string"
			// }
			Description: "Name of custom data identifier.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"regex": {
			// Property: Regex
			// CloudFormation resource type schema:
			// {
			//   "description": "Regular expression for custom data identifier.",
			//   "type": "string"
			// }
			Description: "Regular expression for custom data identifier.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Macie CustomDataIdentifier resource schema",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::CustomDataIdentifier").WithTerraformTypeName("awscc_macie_custom_data_identifier")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"description":            "Description",
		"id":                     "Id",
		"ignore_words":           "IgnoreWords",
		"keywords":               "Keywords",
		"maximum_match_distance": "MaximumMatchDistance",
		"name":                   "Name",
		"regex":                  "Regex",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
