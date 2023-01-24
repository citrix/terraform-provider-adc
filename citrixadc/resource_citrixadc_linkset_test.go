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

const testAccLinkset_add_with_no_binding = `
	resource "citrixadc_linkset" "foo" {
		linkset_id = "LS/1"
	}
`
const testAccLinkset_update_with_binding = `
	resource "citrixadc_linkset" "foo" {
		linkset_id = "LS/1"
	
		interfacebinding = [
			"1/1/1",
			"2/1/1",
		]
	}
`

func TestAccLinkset_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("clustering not supported in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLinksetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLinkset_add_with_no_binding,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinksetExist("citrixadc_linkset.foo", nil),
				),
			},
			resource.TestStep{
				Config: testAccLinkset_update_with_binding,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinksetExist("citrixadc_linkset.foo", nil),
				),
			},
		},
	})
}

func testAccCheckLinksetExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No linkset ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Linkset.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Linkset ID %s not found", n)
		}

		return nil
	}
}

func testAccCheckLinksetDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_linkset" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Linkset.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Linset ID %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
