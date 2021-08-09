package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsTags(t *testing.T) {

	// Construct the terraform options with default retryable errors
	// to handle the most common retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where the Terraform code is located
		TerraformDir: "../",
		// Configure a plan file path so we can introspect the plan and make assertions about it.
		PlanFilePath: "./test/plan.out",
	})

	// Get the plan
	plan := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

	// Check if the plan has required values
	terraform.RequirePlannedValuesMapKeyExists(t, plan, "aws_instance.flugel_test_instance")
	terraform.RequirePlannedValuesMapKeyExists(t, plan, "aws_s3_bucket.flugel_test_bucket")

	// Get the plan for every resource
	ec2Resource := plan.ResourcePlannedValuesMap["aws_instance.flugel_test_instance"]
	s3BucketResource := plan.ResourcePlannedValuesMap["aws_s3_bucket.flugel_test_bucket"]

	// Get the tags of every resource
	ec2Tags := ec2Resource.AttributeValues["tags"].(map[string]interface{})
	s3BucketTags := s3BucketResource.AttributeValues["tags"].(map[string]interface{})

	// Check the tags as it were supposed to be
	expectedTags := map[string]interface{}{
		"Name":  "Flugel",
		"Owner": "InfraTeam",
	}

	assert.Equal(t, ec2Tags, expectedTags)
	assert.Equal(t, s3BucketTags, expectedTags)
}
