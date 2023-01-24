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

const testAccDnsnaptrrec_basic = `

resource "citrixadc_dnsnaptrrec" "dnsnaptrrec" {
	domain     = "example.com"
	order      = 10
	preference = 2
	ttl        = 3600
	replacement = "example1.com"
  }
`

func TestAccDnsnaptrrec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnsnaptrrecDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDnsnaptrrec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnaptrrecExist("citrixadc_dnsnaptrrec.dnsnaptrrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsnaptrrec.dnsnaptrrec", "domain", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnsnaptrrec.dnsnaptrrec", "order", "10"),
					resource.TestCheckResourceAttr("citrixadc_dnsnaptrrec.dnsnaptrrec", "preference", "2"),
					resource.TestCheckResourceAttr("citrixadc_dnsnaptrrec.dnsnaptrrec", "ttl", "3600"),
					resource.TestCheckResourceAttr("citrixadc_dnsnaptrrec.dnsnaptrrec", "replacement", "example1.com"),
				),
			},
		},
	})
}

func testAccCheckDnsnaptrrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsnaptrrec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnsnaptrrec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsnaptrrec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsnaptrrecDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsnaptrrec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnsnaptrrec.Type(), rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("dnsnaptrrec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
