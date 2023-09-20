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

const testAccAuthenticationtacacsaction_add = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction"
		serverip        = "1.2.3.4"
		serverport      = 8080
		authtimeout     = 5
		authorization   = "ON"
		accounting      = "ON"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
`
const testAccAuthenticationtacacsaction_update = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction"
		serverip        = "1.2.3.4"
		serverport      = 8080
		authtimeout     = 5
		authorization   = "OFF"
		accounting      = "OFF"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
`

func TestAccAuthenticationtacacsaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationtacacsactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacsaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "name", "tf_tacacsaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "authorization", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "accounting", "ON"),
				),
			},
			{
				Config: testAccAuthenticationtacacsaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "name", "tf_tacacsaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "authorization", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "accounting", "OFF"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationtacacsactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationtacacsaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationtacacsaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationtacacsaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationtacacsactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationtacacsaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationtacacsaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationtacacsaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
