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

// GoogleBigqueryRoutineInvalidDataGovernanceTypeRule checks the pattern is valid
type GoogleBigqueryRoutineInvalidDataGovernanceTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleBigqueryRoutineInvalidDataGovernanceTypeRule returns new rule with default attributes
func NewGoogleBigqueryRoutineInvalidDataGovernanceTypeRule() *GoogleBigqueryRoutineInvalidDataGovernanceTypeRule {
	return &GoogleBigqueryRoutineInvalidDataGovernanceTypeRule{
		resourceType:  "google_bigquery_routine",
		attributeName: "data_governance_type",
	}
}

// Name returns the rule name
func (r *GoogleBigqueryRoutineInvalidDataGovernanceTypeRule) Name() string {
	return "google_bigquery_routine_invalid_data_governance_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleBigqueryRoutineInvalidDataGovernanceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleBigqueryRoutineInvalidDataGovernanceTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleBigqueryRoutineInvalidDataGovernanceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleBigqueryRoutineInvalidDataGovernanceTypeRule) Check(runner tflint.Runner) error {
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
			validateFunc := validation.StringInSlice([]string{"DATA_MASKING", ""}, false)

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