package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
)

type AwsAcmCertLifecycleRule struct{}

func NewAwsAcmCertLifecycleRule() *AwsAcmCertLifecycleRule {
	return &AwsAcmCertLifecycleRule{}
}

func (r *AwsAcmCertLifecycleRule) Name() string {
	return "aws_acm_cert_lifecycle_rule"
}

func (r *AwsAcmCertLifecycleRule) Enabled() bool {
	return true
}

func (r *AwsAcmCertLifecycleRule) Severity() string {
	return tflint.ERROR
}

func (r *AwsAcmCertLifecycleRule) Link() string {
	return "https://github.com/AleksaC/tflint-acm-cert-lifecycle"
}

func (r *AwsAcmCertLifecycleRule) Check(runner tflint.Runner) error {
	return runner.WalkResources("aws_acm_certificate", func(resource *configs.Resource) error {
		content, _, diags := resource.Config.PartialContent(&hcl.BodySchema{
			Blocks: []hcl.BlockHeaderSchema{
				{Type: "lifecycle"},
			},
		})
		if diags.HasErrors() {
			return diags
		}

		if len(content.Blocks) == 0 {
			if err := runner.EmitIssue(r, "`lifecycle` block not found, needs to contain `create_before_destroy = true`", resource.DeclRange); err != nil {
				return err
			}
			return nil
		}

		block := content.Blocks[0]
		lifecycle, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{Name: "create_before_destroy"},
			},
		})
		if diags.HasErrors() {
			return diags
		}

		if attr, exists := lifecycle.Attributes["create_before_destroy"]; exists {
			var createBeforeDestroy bool
			err := runner.EvaluateExpr(attr.Expr, &createBeforeDestroy, &cty.Bool)

			return runner.EnsureNoError(err, func() error {
				if !createBeforeDestroy {
					return runner.EmitIssueOnExpr(
						r,
						fmt.Sprintf("`create_before_destroy` is set to `%t`", createBeforeDestroy),
						attr.Expr,
					)
				}
				return nil
			})
		} else {
			if err := runner.EmitIssue(r, "create_before_destroy` not set, should be `create_before_destroy = true`", block.DefRange); err != nil {
				return err
			}
		}

		return nil
	})
}
