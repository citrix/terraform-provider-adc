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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccSslrsakey_basic = `

	resource "citrixadc_systemfile" "tf_file" {
		filename 	 = "key1.pem"
		filelocation = "/nsconfig/ssl/"
		filecontent  = "hello"
	}
	resource "citrixadc_sslrsakey" "tf_sslrsakey" {
		reqfile          = "/nsconfig/ssl/test-ca.csr"
		keyfile          = "/nsconfig/ssl/key1.pem"
		countryname      = "in"
		statename        = "kar"
		organizationname = "xyz"
		depends_on = [citrixadc_systemfile.tf_file]
	}
`

func TestAccSslrsakey_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSslrsakey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslrsakeyExist("citrixadc_sslrsakey.tf_sslrsakey", nil),
				),
			},
		},
	})
}

func testAccCheckSslrsakeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslrsakey name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		return nil
	}
}
