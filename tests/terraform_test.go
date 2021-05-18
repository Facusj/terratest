// Remember to : export GOROOT=/usr/local/go/

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraform(t *testing.T) {
	t.Parallel()
	// Tags
	tagName := "Flugel"
	tagOwner := "InfraTeam"

	// Bucket name
	bucketName := "flugel"
	region := "us-east-1"

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "/home/ubuntu/Desktop/terraform",

		Vars: map[string]interface{}{
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// We get the instance ID from the output...
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// ... And we pass it here.
	instanceTags := aws.GetTagsForEc2Instance(t, region, instanceID)

	// We check if the tag exists, and if it's the correct one.
	testTag, containsTag := instanceTags["Name"]
	testTagOwner, containsTagOwner := instanceTags["Owner"]
	assert.True(t, containsTag)
	assert.Equal(t, tagName, testTag)
	assert.True(t, containsTagOwner)
	assert.Equal(t, tagOwner, testTagOwner)

	// We find the name of the bucket with a Tag
	s3Name := aws.FindS3BucketWithTag(t, region, "Name", tagName)
	s3Owner := aws.FindS3BucketWithTag(t, region, "Owner", tagOwner)

	// We compare the name that we want to find, with the one
	// that is stored in s3 ...
	assert.Equal(t, bucketName, s3Name)

	// ... And we check if the Owner tag is the same.
	assert.Equal(t, bucketName, s3Owner)

	// Lastly, we seek that we are pointing the same bucket.
	assert.Equal(t, s3Name, s3Owner)
}