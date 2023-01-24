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
	"fmt"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

const testAccLbvserver_spilloverpolicy_binding_basic = `
	# Since the spilloverpolicy resource is not yet available on Terraform,
	# the tf_spilloverpolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli command:
	# add spillover policy tf_spilloverpolicy -rule TRUE -action SPILLOVER

	resource "citrixadc_lbvserver_spilloverpolicy_binding" "tf_lbvserver_spilloverpolicy_binding" {
		name = citrixadc_lbvserver.tf_lbvserver.name
        policyname = "tf_spilloverpolicy"
        bindpoint = "REQUEST"
        gotopriorityexpression = "END"
        invoke = false
        priority = 1
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

const testAccLbvserver_spilloverpolicy_binding_basic_step2 = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

func TestAccLbvserver_spilloverpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserver_spilloverpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLbvserver_spilloverpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_spilloverpolicy_bindingExist("citrixadc_lbvserver_spilloverpolicy_binding.tf_lbvserver_spilloverpolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccLbvserver_spilloverpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_spilloverpolicy_bindingNotExist("citrixadc_lbvserver_spilloverpolicy_binding.tf_lbvserver_spilloverpolicy_binding", "tf_lbvserver,tf_spilloverpolicy"),
				),
			},
		},
	})
}

func testAccCheckLbvserver_spilloverpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_spilloverpolicy_binding id is set")
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

		lbvserverName := idSlice[0]
		policyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_spilloverpolicy_binding",
			ResourceName:             lbvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbvserver_spilloverpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_spilloverpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		lbvserverName := idSlice[0]
		policyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_spilloverpolicy_binding",
			ResourceName:             lbvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lbvserver_spilloverpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_spilloverpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_spilloverpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Lbvserver_spilloverpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_spilloverpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
