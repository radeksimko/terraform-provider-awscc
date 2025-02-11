// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_servicecatalogappregistry_resource_association", resourceAssociationResourceType)
}

// resourceAssociationResourceType returns the Terraform awscc_servicecatalogappregistry_resource_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ServiceCatalogAppRegistry::ResourceAssociation resource type.
func resourceAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application": {
			// Property: Application
			// CloudFormation resource type schema:
			// {
			//   "description": "The name or the Id of the Application.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name or the Id of the Application.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"application_arn": {
			// Property: ApplicationArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"resource": {
			// Property: Resource
			// CloudFormation resource type schema:
			// {
			//   "description": "The name or the Id of the Resource.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name or the Id of the Resource.",
			Type:        types.StringType,
			Required:    true,
		},
		"resource_arn": {
			// Property: ResourceArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the CFN Resource for now it's enum CFN_STACK.",
			//   "enum": [
			//     "CFN_STACK"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of the CFN Resource for now it's enum CFN_STACK.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"CFN_STACK",
				}),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalogAppRegistry::ResourceAssociation").WithTerraformTypeName("awscc_servicecatalogappregistry_resource_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application":     "Application",
		"application_arn": "ApplicationArn",
		"id":              "Id",
		"resource":        "Resource",
		"resource_arn":    "ResourceArn",
		"resource_type":   "ResourceType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
