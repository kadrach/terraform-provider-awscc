// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package databrew_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDataBrewJobDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DataBrew::Job", "awscc_databrew_job", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSDataBrewJobDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DataBrew::Job", "awscc_databrew_job", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}