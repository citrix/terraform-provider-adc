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
	"strings"
	"testing"
)

const testAccVxlan_srcip_binding_basic = `
	resource "citrixadc_nsip" "tf_srcip" {
		ipaddress = "11.22.33.44"
		type      = "SNIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_vxlan_srcip_binding" "tf_binding" {
		vxlanid = citrixadc_vxlan.tf_vxlan.vxlanid
		srcip   = citrixadc_nsip.tf_srcip.ipaddress
	}
`

const testAccVxlan_srcip_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_nsip" "tf_srcip" {
		ipaddress = "11.22.33.44"
		type      = "SNIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
`

func TestAccVxlan_srcip_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVxlan_srcip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_srcip_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_srcip_bindingExist("citrixadc_vxlan_srcip_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccVxlan_srcip_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_srcip_bindingNotExist("citrixadc_vxlan_srcip_binding.tf_binding", "123,11.22.33.44"),
				),
			},
		},
	})
}

func testAccCheckVxlan_srcip_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vxlan_srcip_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		vxlanid := idSlice[0]
		srcip := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vxlan_srcip_binding",
			ResourceName:             vxlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["srcip"].(string) == srcip {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vxlan_srcip_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVxlan_srcip_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		vxlanid := idSlice[0]
		srcip := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vxlan_srcip_binding",
			ResourceName:             vxlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["srcip"].(string) == srcip {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vxlan_srcip_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVxlan_srcip_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vxlan_srcip_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("vxlan_srcip_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vxlan_srcip_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
