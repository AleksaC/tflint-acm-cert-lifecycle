# tflint-acm-cert-lifecycle
[tflint](https://github.com/terraform-linters/tflint) plugin that ensures lifecycle `create_before_destroy` is set to
`true` for acm certificates as per [`aws_acm_certificate`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate)
official docs:

> It's recommended to specify `create_before_destroy = true` in a
> lifecycle block to replace a certificate which is currently in use

If you don't add this, terraform will timeout when trying to replace a certificate that's in use at the moment of apply.

## Contact 🙋‍♂️

- [Personal website](https://aleksac.me)
- <a target="_blank" href="http://twitter.com/aleksa_c_"><img alt='Twitter followers' src="https://img.shields.io/twitter/follow/aleksa_c_.svg?style=social"></a>
- aleksacukovic1@gmail.com
