// Terratest implementation that deploys the module and performs basic
// validations. Testify suite is used for as the testing framework.
package test

import (
	"fmt"
	"os"
	"testing"

	terratest_azure "github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/suite"
)

const (
	terraformDirectory = "../../"
	// All tests are conducted in westus2
	billToTagValue = "DEVOPS_TESTING_ACCOUNT"
	testLocation   = "westus2"
)

type E2ETestSuite struct {
	suite.Suite
	subscriptionID   string
	projectName      string
	projectShortName string
	terraformOptions *terraform.Options
}

// Test setup includes defining....
//   1. Defining test-specific identifiers to avoid collisions
//   2. Performing basic environment validation against required vars.
//
func (suite *E2ETestSuite) SetupSuite() {
	fmt.Printf("** End2End Setup Initiated ** \n\n")
	suite.subscriptionID = os.Getenv("ARM_SUBSCRIPTION_ID")

	if suite.subscriptionID == "" {
		suite.T().Fatal("ARM_SUBSCRIPTION_ID must be specified in the environment")
	}

	suite.projectShortName = random.UniqueId()
	suite.projectName = fmt.Sprintf("ModTestE2EProject-%s", suite.projectShortName)

	fmt.Printf(
		"Using Suite Values:\n\tProjectName=%s\n\tProjectShortName=%s\n",
		suite.projectName,
		suite.projectShortName)

	suite.terraformOptions = &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: terraformDirectory,
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"projectName":      suite.projectName,
			"projectShortName": suite.projectShortName,
			"billTo":           billToTagValue,
			"location":         testLocation,
			"extra_tags": map[string]string{
				"e2e_test": "true",
				"billTo":   "USER_TAG_OVERWRITE_FOR_BILL_TO",
			},
		},
	}

}

// Teardown only removes any required resources using the terragrunt "Destroy"
// utility.
func (suite *E2ETestSuite) TearDownTest() {
	fmt.Printf("\n** End2End Teardown Initiated ** \n\n")
	terraform.Destroy(suite.T(), suite.terraformOptions)
}

// All tests are run in this single test. It deploys the terraform module
// and performs basic validation against the created resources.
func (suite *E2ETestSuite) TestEnd2EndDeploy() {

	terraform.InitAndApply(suite.T(), suite.terraformOptions)

	resourceGroupOutPutMap := terraform.OutputMap(suite.T(),
		suite.terraformOptions, "resource_group")

	resourceGroupName := resourceGroupOutPutMap["name"]

	exists := terratest_azure.ResourceGroupExists(suite.T(),
		resourceGroupName, suite.subscriptionID)

	// Ensure the resource_group exists
	suite.Assert().True(exists)

	// Ensure the user provided extra tags do not overwrite protected tags
	tagSet := terraform.OutputMap(suite.T(),
		suite.terraformOptions, "tag_set")
	// TODO: Go get the tags off the actual resource_group, rather than rely
	//   on the terraform output to be accurate.
	suite.Assert().Equal(billToTagValue, tagSet["billTo"])
}

func TestEndToEndDeployment(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}
