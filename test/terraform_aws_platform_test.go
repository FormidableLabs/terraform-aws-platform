package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAWSPlatformAfSouth1(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	os.Setenv("AWS_REGION", "af-south-1")
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
	assert.Equal(t, "000000000000", accountId)
	partition := terraform.Output(t, terraformOptions, "partition")
	assert.Equal(t, "aws", partition)
	region := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, "af-south-1", region)
	regionShort := terraform.Output(t, terraformOptions, "region_short")
	assert.Equal(t, "afs1", regionShort)
}

func TestTerraformAWSPlatformUsEast1(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	os.Setenv("AWS_REGION", "us-east-1")
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
	assert.Equal(t, "000000000000", accountId)
	partition := terraform.Output(t, terraformOptions, "partition")
	assert.Equal(t, "aws", partition)
	region := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, "us-east-1", region)
	regionShort := terraform.Output(t, terraformOptions, "region_short")
	assert.Equal(t, "use1", regionShort)
}

func TestTerraformAWSPlatformApSoutheast2(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	os.Setenv("AWS_REGION", "ap-southeast-2")
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
	assert.Equal(t, "000000000000", accountId)
	partition := terraform.Output(t, terraformOptions, "partition")
	assert.Equal(t, "aws", partition)
	region := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, "ap-southeast-2", region)
	regionShort := terraform.Output(t, terraformOptions, "region_short")
	assert.Equal(t, "apse2", regionShort)
}

func TestTerraformAWSPlatformEuNorth1(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	os.Setenv("AWS_REGION", "eu-north-1")
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
	assert.Equal(t, "000000000000", accountId)
	partition := terraform.Output(t, terraformOptions, "partition")
	assert.Equal(t, "aws", partition)
	region := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, "eu-north-1", region)
	regionShort := terraform.Output(t, terraformOptions, "region_short")
	assert.Equal(t, "eun1", regionShort)
}
