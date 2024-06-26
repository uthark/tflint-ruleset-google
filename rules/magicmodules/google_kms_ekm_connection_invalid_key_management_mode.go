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

// GoogleKmsEkmConnectionInvalidKeyManagementModeRule checks the pattern is valid
type GoogleKmsEkmConnectionInvalidKeyManagementModeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleKmsEkmConnectionInvalidKeyManagementModeRule returns new rule with default attributes
func NewGoogleKmsEkmConnectionInvalidKeyManagementModeRule() *GoogleKmsEkmConnectionInvalidKeyManagementModeRule {
	return &GoogleKmsEkmConnectionInvalidKeyManagementModeRule{
		resourceType:  "google_kms_ekm_connection",
		attributeName: "key_management_mode",
	}
}

// Name returns the rule name
func (r *GoogleKmsEkmConnectionInvalidKeyManagementModeRule) Name() string {
	return "google_kms_ekm_connection_invalid_key_management_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleKmsEkmConnectionInvalidKeyManagementModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleKmsEkmConnectionInvalidKeyManagementModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleKmsEkmConnectionInvalidKeyManagementModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleKmsEkmConnectionInvalidKeyManagementModeRule) Check(runner tflint.Runner) error {
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
			validateFunc := validation.StringInSlice([]string{"MANUAL", "CLOUD_KMS", ""}, false)

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
