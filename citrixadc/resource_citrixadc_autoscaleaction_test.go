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

const testAccAutoscaleaction_basic = `


resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
	name        = "my_autoscaleaction"
	type        = "SCALE_UP"
	profilename = "my_profile"
	vserver     = "my_vserver"
	parameters  = "my_parameters"
  }
`
const testAccAutoscaleaction_update = `


resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
	name        = "my_autoscaleaction"
	type        = "SCALE_DOWN"
	profilename = "my_profile"
	vserver     = "my_vserver2"
	parameters  = "my_parameters"
  }
`


func TestAccAutoscaleaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAutoscaleactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAutoscaleaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_autoscaleaction", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "name", "my_autoscaleaction"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "type", "SCALE_UP"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "profilename", "my_profile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "vserver", "my_vserver"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "parameters", "my_parameters"),
				),
			},
			resource.TestStep{
				Config: testAccAutoscaleaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_autoscaleaction", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "name", "my_autoscaleaction"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "type", "SCALE_DOWN"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "profilename", "my_profile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "vserver", "my_vserver2"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "parameters", "my_parameters"),
				),
			},
		},
	})
}

func testAccCheckAutoscaleactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No autoscaleaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Autoscaleaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("autoscaleaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAutoscaleactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_autoscaleaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Autoscaleaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("autoscaleaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
