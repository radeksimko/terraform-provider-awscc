// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package devopsguru

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_devopsguru_resource_collection", resourceCollectionDataSourceType)
}

// resourceCollectionDataSourceType returns the Terraform awscc_devopsguru_resource_collection data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DevOpsGuru::ResourceCollection resource type.
func resourceCollectionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"resource_collection_filter": {
			// Property: ResourceCollectionFilter
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
			//   "properties": {
			//     "CloudFormation": {
			//       "additionalProperties": false,
			//       "description": "CloudFormation resource for DevOps Guru to monitor",
			//       "properties": {
			//         "StackNames": {
			//           "description": "An array of CloudFormation stack names.",
			//           "items": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "maxItems": 200,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cloudformation": {
						// Property: CloudFormation
						Description: "CloudFormation resource for DevOps Guru to monitor",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"stack_names": {
									// Property: StackNames
									Description: "An array of CloudFormation stack names.",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"resource_collection_type": {
			// Property: ResourceCollectionType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of ResourceCollection",
			//   "enum": [
			//     "AWS_CLOUD_FORMATION"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of ResourceCollection",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::DevOpsGuru::ResourceCollection",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::ResourceCollection").WithTerraformTypeName("awscc_devopsguru_resource_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudformation":             "CloudFormation",
		"resource_collection_filter": "ResourceCollectionFilter",
		"resource_collection_type":   "ResourceCollectionType",
		"stack_names":                "StackNames",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
