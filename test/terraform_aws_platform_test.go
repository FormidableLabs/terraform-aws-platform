package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAWSPlatform(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	accountId := terraform.Output(t, terraformOptions, "account_id")
	assert.Equal(t, "912039392020", accountId)
	partition := terraform.Output(t, terraformOptions, "partition")
	assert.Equal(t, "aws", partition)
	region := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, "us-east-1", region)
	regionShort := terraform.Output(t, terraformOptions, "region_short")
	assert.Equal(t, "use1", regionShort)
}
