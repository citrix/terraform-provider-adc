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

const testAccAuthenticationwebauthaction_add = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "TRUE"
		scheme                     = "https"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "old_group"
	}
`
const testAccAuthenticationwebauthaction_update = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "FALSE"
		scheme                     = "http"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "new_group"
	}
`

func TestAccAuthenticationwebauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationwebauthactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationwebauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_webauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "name", "tf_webauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "fullreqexpr", "TRUE"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "scheme", "https"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationwebauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_webauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "name", "tf_webauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "fullreqexpr", "FALSE"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "scheme", "http"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationwebauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationwebauthaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationwebauthaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationwebauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationwebauthactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationwebauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationwebauthaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationwebauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
