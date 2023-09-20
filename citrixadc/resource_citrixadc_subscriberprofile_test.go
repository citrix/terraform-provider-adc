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

const testAccSubscriberprofile_basic = `


resource "citrixadc_subscriberprofile" "tf_subscriberprofile" {
	ip                  = "10.222.74.185"
	subscriptionidtype  = "E164"
	subscriptionidvalue = 5
  }
  
`

const testAccSubscriberprofile_update = `


resource "citrixadc_subscriberprofile" "tf_subscriberprofile" {
	ip                  = "10.222.74.185"
	subscriptionidtype  = "IMSI"
	subscriptionidvalue = 10
  }
  
`

func TestAccSubscriberprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubscriberprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_subscriberprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "ip", "10.222.74.185"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidtype", "E164"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidvalue", "5"),
				),
			},
			{
				Config: testAccSubscriberprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_subscriberprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "ip", "10.222.74.185"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidtype", "IMSI"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidvalue", "10"),
				),
			},
		},
	})
}

func testAccCheckSubscriberprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscriberprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("subscriberprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscriberprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSubscriberprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_subscriberprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("subscriberprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("subscriberprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
