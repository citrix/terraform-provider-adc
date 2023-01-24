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

const testAccAppfwpolicy_add = `
	resource citrixadc_appfwprofile tfAcc_appfwprofile {
		name = "tfAcc_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	resource citrixadc_appfwpolicy tfAcc_appfwpolicy1 {
		name = "tfAcc_appfwpolicy1"
		profilename = citrixadc_appfwprofile.tfAcc_appfwprofile.name
		rule = "true"
	}
`
const testAccAppfwpolicy_update = `
	resource citrixadc_appfwprofile tfAcc_appfwprofile {
		name = "tfAcc_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	resource citrixadc_appfwpolicy tfAcc_appfwpolicy1 {
		name = "tfAcc_appfwpolicy1"
		profilename = citrixadc_appfwprofile.tfAcc_appfwprofile.name
		rule = "true"
        comment = "test comment"
	}
`

func TestAccAppfwpolicy_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwpolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppfwpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "name", "tfAcc_appfwpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "profilename", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "rule", "true"),
				),
			},
			resource.TestStep{
				Config: testAccAppfwpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "name", "tfAcc_appfwpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "profilename", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "comment", "test comment"),
				),
			},
		},
	})
}

func testAccCheckAppfwpolicyExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Appfwpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
