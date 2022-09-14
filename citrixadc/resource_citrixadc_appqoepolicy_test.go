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

const testAccAppqoepolicy_basic = `

	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "my_appqoepolicy"
		rule   = "true"
		action = "my_act"
	}
`
const testAccAppqoepolicy_update = `

	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "my_appqoepolicy"
		rule   = "false"
		action = "my_act"
	}
`
func TestAccAppqoepolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppqoepolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppqoepolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoepolicyExist("citrixadc_appqoepolicy.tf_appqoepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoepolicy.tf_appqoepolicy", "name", "my_appqoepolicy"),
					resource.TestCheckResourceAttr("citrixadc_appqoepolicy.tf_appqoepolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_appqoepolicy.tf_appqoepolicy", "action", "my_act"),
				),
			},
			resource.TestStep{
				Config: testAccAppqoepolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoepolicyExist("citrixadc_appqoepolicy.tf_appqoepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoepolicy.tf_appqoepolicy", "name", "my_appqoepolicy"),
					resource.TestCheckResourceAttr("citrixadc_appqoepolicy.tf_appqoepolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_appqoepolicy.tf_appqoepolicy", "action", "my_act"),
				),
			},
		},
	})
}

func testAccCheckAppqoepolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appqoepolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Appqoepolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appqoepolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppqoepolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appqoepolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appqoepolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appqoepolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
