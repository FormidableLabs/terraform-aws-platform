locals {
  # Regex:
  #   Group 1 "([^-]*)" - Captures the country code up to the first "-"
  #   Group 2 "(.)" - Captures the first character of the region locality (ex "s" of southeast)
  #   Group 3 "(?:o[ru]th|[ae]st|entral)" - Groups, but does not capture (denoted by `?:`) the
  #           first part of the region locality, minus the first character (ex "outh" of southeast)
  #   Group 4 "(.?)" - Captures a single character of the second cardinal direction of an ordinal
  #           direction (ex "e" of southeast or "" of north)
  #   Ungrouped "[^-]*-" - Matches and throws away all other characters up to and including the
  #           last "-"
  #   Group 5 "(.*)" - Captures the remaining characters of the region, the numerical identifier
  #           (ex "2" in us-east-2)
  # This is used to convert the AWS region codes into a shortend code (ex "use1" from us-east-1
  # or "apse1" from ap-southeast-1)
  shortener_regex = "^([^-]*)-(.)(?:o[ru]th|[ae]st|entral)(.?)[^-]*-(.*)$"

  account_id = data.aws_caller_identity.this.account_id
  partition  = data.aws_partition.this.partition
  region     = data.aws_region.this.name
}

# Data Sources
data "aws_caller_identity" "this" {}
data "aws_partition" "this" {}
data "aws_region" "this" {}

# Outputs
output "account_id" {
  value = local.account_id
}

output "partition" {
  value = local.partition
}

output "region" {
  value = local.region
}

output "region_short" {
  # the regex function produces an array of matched groups, these
  # are joined together with an empty separator to form the shortened
  # AWS region code (ex (us)-(e)ast-(1) becomes "use1" or
  # (ap)-(s)outh(e)ast-(1) becomes "apse1")
  value = join("", regex(local.shortener_regex, local.region))
}
