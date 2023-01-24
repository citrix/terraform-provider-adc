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

const testAccAuthenticationloginschema_add = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "NO"
		authenticationstrength  = "30"
		passwordcredentialindex = "10"
	}
`
const testAccAuthenticationloginschema_update = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "YES"
		authenticationstrength  = "20"
		passwordcredentialindex = "10"
	}
`

func TestAccAuthenticationloginschema_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationloginschemaDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationloginschema_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_loginschema", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "name", "tf_loginschema"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "ssocredentials", "NO"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "authenticationstrength", "30"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationloginschema_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_loginschema", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "name", "tf_loginschema"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "ssocredentials", "YES"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "authenticationstrength", "20"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationloginschemaExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationloginschema name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationloginschema.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationloginschema %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationloginschemaDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationloginschema" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationloginschema.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationloginschema %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
