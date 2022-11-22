# terraform-aws-platform

[![Maintenance Status][maintenance-image]](#maintenance-status)

## Purpose

This module is designed to output the AWS account ID, partition and region. In addition, it converts AWS region names into a shortened name. The shortened region name is created using a regex that matches the first part, the abbreviated ordinal direction and the numeric identifier. Examples of the `region_short` output are below.

| AWS Region Name | Shortened Name |
| --------------- | -------------- |
| ap-southeast-1  | apse1          |
| eu-central-1    | euc1           |
| us-east-1       | use1           |

Note: this conversion does not currently work for regions that contain four parts, like the `us-gov-east-1` regions.

## Example Usage

```
locals {
  account_id   = module.platform.account_id
  partition    = module.platform.partition
  region       = module.platform.region
  region_short = module.platform.region_short
}

module "platform" {
  source = "github.com/FormidableLabs/terraform-aws-platform?ref=v0.1"
}

resource "aws_lambda_function" "test_lambda" {
  filename      = "lambda_function_payload.zip"
  function_name = "${local.region_short}-lambda_function_name"
  handler       = "index.test"
  ...

}

```

[maintenance-image]: https://img.shields.io/badge/maintenance-active-green.svg?color=brightgreen&style=flat

## Maintenance Status

**Active:** Formidable is actively working on this project, and we expect to continue for work for the foreseeable future. Bug reports, feature requests and pull requests are welcome.
