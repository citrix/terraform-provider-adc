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

const testAccTmsessionparameter_basic = `


resource "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
	sesstimeout                = 40
	defaultauthorizationaction = "ALLOW"
	sso                        = "ON"
	ssodomain                  = 3
  }
  
`
const testAccTmsessionparameter_update = `


resource "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
	sesstimeout                = 50
	defaultauthorizationaction = "DENY"
	sso                        = "OFF"
  }
  
`

func TestAccTmsessionparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTmsessionparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_tmsessionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sesstimeout", "40"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "defaultauthorizationaction", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sso", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "ssodomain", "3"),
				),
			},
			resource.TestStep{
				Config: testAccTmsessionparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_tmsessionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sesstimeout", "50"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "defaultauthorizationaction", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sso", "OFF"),
				),
			},
		},
	})
}

func testAccCheckTmsessionparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmsessionparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Tmsessionparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmsessionparameter %s not found", n)
		}

		return nil
	}
}