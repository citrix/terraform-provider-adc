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

const testAccSystemglobal_authenticationradiuspolicy_binding_basic = `

	resource "citrixadc_systemglobal_authenticationradiuspolicy_binding" "tf_systemglobal_authenticationradiuspolicy_binding" {
		policyname = citrixadc_authenticationradiuspolicy.tf_radiuspolicy.name
		priority   = 50
	}
	
	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "DISABLED"
		passencoding = "chap"
	}
	resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
		name      = "tf_radiuspolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
	}
`

const testAccSystemglobal_authenticationradiuspolicy_binding_basic_step2 = `
	
resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
	name         = "tf_radiusaction"
	radkey       = "secret"
	serverip     = "1.2.3.4"
	serverport   = 8080
	authtimeout  = 2
	radnasip     = "DISABLED"
	passencoding = "chap"
}
resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
	name      = "tf_radiuspolicy"
	rule      = "NS_TRUE"
	reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
}
`

func TestAccSystemglobal_authenticationradiuspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemglobal_authenticationradiuspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_authenticationradiuspolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_authenticationradiuspolicy_bindingExist("citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding", nil),
				),
			},
			{
				Config: testAccSystemglobal_authenticationradiuspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_authenticationradiuspolicy_bindingNotExist("citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding", "tf_radiuspolicy"),
				),
			},
		},
	})
}

func testAccCheckSystemglobal_authenticationradiuspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemglobal_authenticationradiuspolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "systemglobal_authenticationradiuspolicy_binding",
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
			return fmt.Errorf("systemglobal_authenticationradiuspolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_authenticationradiuspolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "systemglobal_authenticationradiuspolicy_binding",
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("systemglobal_authenticationradiuspolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_authenticationradiuspolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemglobal_authenticationradiuspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Systemglobal_authenticationradiuspolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemglobal_authenticationradiuspolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
