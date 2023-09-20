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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccIcaparameter_basic = `


resource "citrixadc_icaparameter" "tf_icaparameter" {
	edtpmtuddf           = "ENABLED"
	edtpmtuddftimeout    = 200
	l7latencyfrequency   = 0
	enablesronhafailover = "YES"
  }
  
`
const testAccIcaparameter_update = `

resource "citrixadc_icaparameter" "tf_icaparameter" {
	edtpmtuddf           = "ENABLED"
	edtpmtuddftimeout    = 100
	l7latencyfrequency   = 30
	enablesronhafailover = "NO"
  }
  
`

func TestAccIcaparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_icaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddf", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddftimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "l7latencyfrequency", "0"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "enablesronhafailover", "YES"),
				),
			},
			{
				Config: testAccIcaparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_icaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddf", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddftimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "l7latencyfrequency", "30"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "enablesronhafailover", "NO"),
				),
			},
		},
	})
}

func testAccCheckIcaparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("icaparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icaparameter %s not found", n)
		}

		return nil
	}
}
