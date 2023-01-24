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

const testAccLsnparameter_basic = `

	resource "citrixadc_lsnparameter" "tf_lsnparameter" {
		sessionsync          = "ENABLED"
		subscrsessionremoval = "ENABLED"
	}
`
const testAccLsnparameter_update = `

	resource "citrixadc_lsnparameter" "tf_lsnparameter" {
		sessionsync          = "DISABLED"
		subscrsessionremoval = "DISABLED"
	}
`

func TestAccLsnparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLsnparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_lsnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "sessionsync", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "subscrsessionremoval", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccLsnparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_lsnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "sessionsync", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "subscrsessionremoval", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckLsnparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("lsnparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnparameter %s not found", n)
		}

		return nil
	}
}