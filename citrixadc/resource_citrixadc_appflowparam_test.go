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

const testAccAppflowparam_basic = `

resource "citrixadc_appflowparam" "tf_appflowparam" {
	templaterefresh     = 200
	flowrecordinterval  = 200
	httpcookie          = "ENABLED"
	httplocation        = "ENABLED"
  }
  
`
const testAccAppflowparam_update = `

resource "citrixadc_appflowparam" "tf_appflowparam" {
	templaterefresh     = 600
	flowrecordinterval  = 100
	httpcookie          = "DISABLED"
	httplocation        = "DISABLED"
  }  
`
func TestAccAppflowparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppflowparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh" , "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccAppflowparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh" , "600"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "100"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckAppflowparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Appflowparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowparam %s not found", n)
		}

		return nil
	}
}

