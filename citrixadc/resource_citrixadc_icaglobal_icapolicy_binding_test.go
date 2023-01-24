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

const testAccIcaglobal_icapolicy_binding_basic = `

resource "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
	policyname = "my_ica_policy"
	priority   = 100
	type       = "ICA_REQ_DEFAULT"
  }
`

const testAccIcaglobal_icapolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccIcaglobal_icapolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIcaglobal_icapolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIcaglobal_icapolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaglobal_icapolicy_bindingExist("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccIcaglobal_icapolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaglobal_icapolicy_bindingNotExist("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", "my_ica_policy", "ICA_REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckIcaglobal_icapolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaglobal_icapolicy_binding id is set")
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
			ResourceType:             "icaglobal_icapolicy_binding",
			ArgsMap:				  map[string]string {"type": rs.Primary.Attributes["type"] },
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
			return fmt.Errorf("icaglobal_icapolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcaglobal_icapolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "icaglobal_icapolicy_binding",
			ArgsMap:				  map[string]string {"type":typename },
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
			return fmt.Errorf("icaglobal_icapolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckIcaglobal_icapolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icaglobal_icapolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("icaglobal_icapolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icaglobal_icapolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
