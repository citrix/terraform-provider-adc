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

const testAccNshttpprofile_add = `
	resource "citrixadc_nshttpprofile" "foo" {
		name  = "tf_httpprofile"
		http2 = "ENABLED"
        markrfc7230noncompliantinval = "ENABLED"
        markhttpheaderextrawserror = "ENABLED"
        dropinvalreqs = "ENABLED"
	}  
`
const testAccNshttpprofile_update = `
	resource "citrixadc_nshttpprofile" "foo" {
		name  = "tf_httpprofile"
		http2 = "DISABLED"
        markrfc7230noncompliantinval = "DISABLED"
        markhttpheaderextrawserror = "DISABLED"
        dropinvalreqs = "DISABLED"
	}  
`

func TestAccNshttpprofile_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNshttpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshttpprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "name", "tf_httpprofile"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markrfc7230noncompliantinval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markhttpheaderextrawserror", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "dropinvalreqs", "ENABLED"),
				),
			},
			{
				Config: testAccNshttpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "name", "tf_httpprofile"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markrfc7230noncompliantinval", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markhttpheaderextrawserror", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "dropinvalreqs", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckNshttpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No NS HTTP Profile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nshttpprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("NS HTTP Profile %s not found", n)
		}

		return nil
	}
}

func testAccCheckNshttpprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nshttpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nshttpprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("NS HTTP Profile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
