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
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"fmt"
	"strings"
	"testing"
)

const testAccDnsnsrec_basic_step1 = `

resource "citrixadc_dnsnsrec" "tf_dnsnsrec1" {
    domain = "www.test.com"
    nameserver = "192.168.1.100"
	ttl = 4000
}

resource "citrixadc_dnsnsrec" "tf_dnsnsrec2" {
    domain = "www.test.com"
    nameserver = "192.168.1.99"
	ttl = 4000
}
`

const testAccDnsnsrec_basic_step2 = `

resource "citrixadc_dnsnsrec" "tf_dnsnsrec1" {
    domain = "www.test.com"
    nameserver = "192.168.1.100"
	ttl = 4000
}

resource "citrixadc_dnsnsrec" "tf_dnsnsrec2" {
    domain = "www.test.com"
    nameserver = "192.168.1.98"
	ttl = 4000
}
`

func TestAccDnsnsrec_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnsnsrecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsnsrec_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec1", nil),
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec2", nil),
				),
			},
			{
				Config: testAccDnsnsrec_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec1", nil),
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec2", nil),
				),
			},
		},
	})
}

func testAccCheckDnsnsrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsnsrec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		dnsnsrecId := rs.Primary.ID

		idSlice := strings.SplitN(dnsnsrecId, ",", 2)
		domain := idSlice[0]
		nameserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType: "dnsnsrec",
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		foundIndex := -1
		for i, v := range dataArr {
			if v["domain"] == domain && v["nameserver"] == nameserver {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("Cannot find dnsnsrec with id %v", dnsnsrecId)
		}

		return nil
	}
}

func testAccCheckDnsnsrecDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsnsrec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnsnsrec.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsnsrec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
