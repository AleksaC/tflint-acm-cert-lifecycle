package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
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
	return "https://github.com/AleksaC/tflint-acm-cert-lifecycle/blob/main/README.md"
}

func (r *AwsAcmCertLifecycleRule) Check(runner tflint.Runner) error {
	return runner.WalkResources("aws_acm_certificate", func(resource *configs.Resource) error {
		if !resource.Managed.CreateBeforeDestroy {
			if err := runner.EmitIssue(r, "lifecycle {\n  create_before_destroy = true\n} needs to be set for `aws_acm_certificate`", resource.DeclRange); err != nil {
				return err
			}
		}
		return nil
	})
}
