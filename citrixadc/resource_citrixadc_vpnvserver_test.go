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

const testAccVpnvserver_add = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "DISABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
	}
`

const testAccVpnvserver_update = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "ENABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
	}
`

func TestAccVpnvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "downstateflush", "DISABLED"),
				),
			},
			{
				Config: testAccVpnvserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "downstateflush", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckVpnvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vpnvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserverDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
