package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsAcmCertLifecycleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no lifecycle block",
			Content: `
resource "aws_acm_certificate" "assets" {
  domain_name       = local.assets_domain_name
  validation_method = "DNS"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsAcmCertLifecycleRule(),
					Message: "`lifecycle` block not found, needs to contain `create_before_destroy = true`",
					Range: hcl.Range{
						Filename: "cert.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 40},
					},
				},
			},
		},
		{
			Name: "no create_before_destroy attribute in lifecycle block",
			Content: `
resource "aws_acm_certificate" "assets" {
  domain_name       = local.assets_domain_name
  validation_method = "DNS"

  lifecycle {}
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsAcmCertLifecycleRule(),
					Message: "create_before_destroy` not set, should be `create_before_destroy = true`",
					Range: hcl.Range{
						Filename: "cert.tf",
						Start:    hcl.Pos{Line: 6, Column: 3},
						End:      hcl.Pos{Line: 6, Column: 12},
					},
				},
			},
		},
		{
			Name: "create_before_destroy = false",
			Content: `
resource "aws_acm_certificate" "assets" {
  domain_name       = local.assets_domain_name
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = false
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsAcmCertLifecycleRule(),
					Message: "`create_before_destroy` is set to `false`",
					Range: hcl.Range{
						Filename: "cert.tf",
						Start:    hcl.Pos{Line: 7, Column: 29},
						End:      hcl.Pos{Line: 7, Column: 34},
					},
				},
			},
		},
	}

	rule := NewAwsAcmCertLifecycleRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"cert.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
