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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccAaapreauthenticationpolicy_basic = `

	resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
		name = "my_policy"
		rule = "REQ.VLANID == 5"
		reqaction = "my_action"
	}
`
const testAccAaapreauthenticationpolicy_update = `

	resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
		name = "my_policy"
		rule = "REQ.VLANID == 10"
		reqaction = "my_action2"
	}
`

func TestAccAaapreauthenticationpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAaapreauthenticationpolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAaapreauthenticationpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationpolicyExist("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", "name", "my_policy"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", "rule", "REQ.VLANID == 5"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", "reqaction", "my_action"),
				),
			},
			resource.TestStep{
				Config: testAccAaapreauthenticationpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationpolicyExist("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", "name", "my_policy"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", "rule", "REQ.VLANID == 10"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy", "reqaction", "my_action2"),
				),
			},
		},
	})
}

func testAccCheckAaapreauthenticationpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaapreauthenticationpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Aaapreauthenticationpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaapreauthenticationpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaapreauthenticationpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaapreauthenticationpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Aaapreauthenticationpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaapreauthenticationpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
