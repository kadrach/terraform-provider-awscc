// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSDataSyncLocationFSxWindows_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DataSync::LocationFSxWindows", "aws_datasync_location_f_sx_windows", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}