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

const testAccSslprofile_add = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`
const testAccSslprofile_update = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		hsts = "ENABLED"
		ecccurvebindings = []
	}
`

func TestAccSslprofile_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
			resource.TestStep{
				Config: testAccSslprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "hsts", "ENABLED"),
				),
			},
		},
	})
}

const testAccSslprofile_ecccurvebinding_bind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = ["P_256"]
	}
`
const testAccSslprofile_ecccurvebinding_unbind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccSslprofile_ecccurve_binding(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslprofile_ecccurvebinding_bind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
			resource.TestStep{
				Config: testAccSslprofile_ecccurvebinding_unbind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
		},
	})
}

const testAccSslprofile_cipherbinding_bind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
		cipherbindings {
			ciphername     = "HIGH"
			cipherpriority = 10
		  }
	}
`
const testAccSslprofile_cipherbinding_unbind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccSslprofile_cipher_binding(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslprofile_cipherbinding_bind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
			resource.TestStep{
				Config: testAccSslprofile_cipherbinding_unbind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
		},
	})
}

func testAccCheckSslprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SSL Profile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Sslprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Profile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Sslprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SSL Profile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
