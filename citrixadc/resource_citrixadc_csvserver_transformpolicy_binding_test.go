/*
Copyright 2016 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"fmt"
	"strings"
	"testing"
)

const testAccCsvserver_transformpolicy_binding_basic_step1 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.34"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_csvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 100
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
}
`

const testAccCsvserver_transformpolicy_binding_basic_step2 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.34"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_csvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 110
    bindpoint = "REQUEST"
    gotopriorityexpression = "NEXT"
}
`

func TestAccCsvserver_transformpolicy_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserver_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_transformpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_transformpolicy_bindingExist("citrixadc_csvserver_transformpolicy_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_transformpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_transformpolicy_bindingExist("citrixadc_csvserver_transformpolicy_binding.tf_binding", nil),
				),
			},
		},
	})
}

func testAccCheckCsvserver_transformpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_transformpolicy_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 2)
		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_transformpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == policyname {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("csvserver_transformpolicy_binding %s not found", bindingId)
		}

		return nil
	}
}

func testAccCheckCsvserver_transformpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_transformpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Csvserver_transformpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_transformpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
