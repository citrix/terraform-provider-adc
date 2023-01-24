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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccCsaction_create = `

resource "citrixadc_csaction" "foo" {
  name            = "tf_test_csaction"
  targetlbvserver = citrixadc_lbvserver.tf_image_lb.name
  comment         = "Forwards image requests to the image_lb"
}

resource "citrixadc_lbvserver" "tf_image_lb" {
  name        = "image_lb"
  ipv46       = "10.0.2.5"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_video_lb" {
  name        = "video_lb"
  ipv46       = "10.0.2.6"
  port        = "80"
  servicetype = "HTTP"
}

`

const testAccCsaction_update = `

resource "citrixadc_csaction" "foo" {
  name            = "tf_test_csaction"
  targetlbvserver = citrixadc_lbvserver.tf_video_lb.name
  comment         = "Forwards video requests to the video_lb"
}

resource "citrixadc_lbvserver" "tf_image_lb" {
  name        = "image_lb"
  ipv46       = "10.0.2.5"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_video_lb" {
  name        = "video_lb"
  ipv46       = "10.0.2.6"
  port        = "80"
  servicetype = "HTTP"
}

`

const testAccCsaction_update_name = `

resource "citrixadc_csaction" "foo" {
  name            = "tf_test_csaction_newname"
  targetlbvserver = citrixadc_lbvserver.tf_image_lb.name
  comment         = "Forwards video requests to the image_lb"
}

resource "citrixadc_lbvserver" "tf_image_lb" {
  name        = "image_lb"
  ipv46       = "10.0.2.5"
  port        = "80"
  servicetype = "HTTP"
}

`

func TestAccCsaction_create_update(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsaction_create,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsactionExist("citrixadc_csaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "name", "tf_test_csaction"),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "targetlbvserver", "image_lb"),
				),
			},
			resource.TestStep{
				Config: testAccCsaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsactionExist("citrixadc_csaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "name", "tf_test_csaction"),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "targetlbvserver", "video_lb"),
				),
			},
		},
	})
}

func TestAccCsaction_create_update_name(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsaction_create,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsactionExist("citrixadc_csaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "name", "tf_test_csaction"),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "targetlbvserver", "image_lb"),
				),
			},
			resource.TestStep{
				Config: testAccCsaction_update_name,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsactionExist("citrixadc_csaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "name", "tf_test_csaction_newname"),
					resource.TestCheckResourceAttr("citrixadc_csaction.foo", "targetlbvserver", "image_lb"),
				),
			},
		},
	})
}

func testAccCheckCsactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Csaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Csaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
