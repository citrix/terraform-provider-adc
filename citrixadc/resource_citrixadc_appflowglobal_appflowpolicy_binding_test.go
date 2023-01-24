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
	"testing"
)

const testAccAppflowglobal_appflowpolicy_binding_basic = `

	resource "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
		policyname     = "test_policy"
		globalbindtype = "SYSTEM_GLOBAL"
		type           = "REQ_OVERRIDE"
		priority       = 55
	}
	
	# -------------------- ADC CLI ----------------------------
	#add appflow collector tf_collector -IPAddress 192.168.2.2
	#add appflowaction test_action -collectors tf_collector
	#add appflowpolicy test_policy client.TCP.DSTPORT.EQ(22) test_action
	
	
	# ---------------- NOT YET IMPLEMENTED -------------------
	# resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	#   name   = "test_policy"
	#   action = citrixadc_appflowaction.tf_appflowaction.name
	#   rule   = "client.TCP.DSTPORT.EQ(22)"
	# }
	# resource "citrixadc_appflowaction" "tf_appflowaction" {
	#   name            = "test_action"
	#   collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
	#   securityinsight = "ENABLED"
	#   botinsight      = "ENABLED"
	#   videoanalytics  = "ENABLED"
	# }
	# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	#   name      = "tf_collector"
	#   ipaddress = "192.168.2.2"
	#   port      = 80
	# }
`

const testAccAppflowglobal_appflowpolicy_binding_basic_step2 = `

	
	# -------------------- ADC CLI ----------------------------
	#add appflow collector tf_collector -IPAddress 192.168.2.2
	#add appflowaction test_action -collectors tf_collector
	#add appflowpolicy test_policy client.TCP.DSTPORT.EQ(22) test_action
	
	
	# ---------------- NOT YET IMPLEMENTED -------------------
	# resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	#   name   = "test_policy"
	#   action = citrixadc_appflowaction.tf_appflowaction.name
	#   rule   = "client.TCP.DSTPORT.EQ(22)"
	# }
	# resource "citrixadc_appflowaction" "tf_appflowaction" {
	#   name            = "test_action"
	#   collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
	#   securityinsight = "ENABLED"
	#   botinsight      = "ENABLED"
	#   videoanalytics  = "ENABLED"
	# }
	# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	#   name      = "tf_collector"
	#   ipaddress = "192.168.2.2"
	#   port      = 80
	# }
`

func TestAccAppflowglobal_appflowpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppflowglobal_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppflowglobal_appflowpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowglobal_appflowpolicy_bindingExist("citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccAppflowglobal_appflowpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowglobal_appflowpolicy_bindingNotExist("citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", "test3_policy", "REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckAppflowglobal_appflowpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowglobal_appflowpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := rs.Primary.ID
		typename := rs.Primary.Attributes["type"]

		findParams := service.FindParams{
			ResourceType:             "appflowglobal_appflowpolicy_binding",
			ArgsMap: 				  map[string]string{ "type":typename },
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appflowglobal_appflowpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowglobal_appflowpolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		policyname := id
		findParams := service.FindParams{
			ResourceType:             "appflowglobal_appflowpolicy_binding",
			ArgsMap: 				  map[string]string{ "type":typename },
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appflowglobal_appflowpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppflowglobal_appflowpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowglobal_appflowpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appflowglobal_appflowpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowglobal_appflowpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
