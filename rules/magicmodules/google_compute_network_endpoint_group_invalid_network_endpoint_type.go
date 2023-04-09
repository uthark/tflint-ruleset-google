// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule checks the pattern is valid
type GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule returns new rule with default attributes
func NewGoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule() *GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule {
	return &GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule{
		resourceType:  "google_compute_network_endpoint_group",
		attributeName: "network_endpoint_type",
	}
}

// Name returns the rule name
func (r *GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule) Name() string {
	return "google_compute_network_endpoint_group_invalid_network_endpoint_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			validateFunc := validation.StringInSlice([]string{"GCE_VM_IP", "GCE_VM_IP_PORT", "NON_GCP_PRIVATE_IP_PORT", ""}, false)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
