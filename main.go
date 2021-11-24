package main

import (
	"github.com/AleksaC/tflint-ruleset-acm-cert-lifecycle/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "acm_cert_lifecycle",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewAwsAcmCertLifecycleRule(),
			},
		},
	})
}
