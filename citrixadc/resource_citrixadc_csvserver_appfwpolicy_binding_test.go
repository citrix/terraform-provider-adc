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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccCsvserver_appfwpolicy_binding_basic = `
	resource citrixadc_csvserver_appfwpolicy_binding demo_binding {
		name = citrixadc_csvserver.demo_cs.name
		priority = 100
		policyname  = citrixadc_appfwpolicy.demo_appfwpolicy.name
		gotopriorityexpression = "END"
	}
	resource "citrixadc_csvserver" "demo_cs" {
		ipv46       = "10.10.10.33"
		name        = "demo_csvserver"
		port        = 80
		servicetype = "HTTP"
	}

	resource citrixadc_appfwprofile demo_appfwprofile {
		name = "demo_appfwprofile"
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

	resource citrixadc_appfwpolicy demo_appfwpolicy {
		name = "demo_appfwpolicy"
		profilename = citrixadc_appfwprofile.demo_appfwprofile.name
		rule = "true"
	}
`

func TestAccCsvserver_appfwpolicy_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsvserver_appfwpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appfwpolicy_bindingExist("citrixadc_csvserver_appfwpolicy_binding.demo_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "name", "demo_csvserver"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "policyname", "demo_appfwpolicy"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_appfwpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_appfwpolicy_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 2)
		csvserverName := idSlice[0]
		appfwPolicyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             service.Csvserver_appfwpolicy_binding.Type(),
			ResourceName:             csvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == appfwPolicyName {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find csvserver_appfwpolicy_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckCsvserver_appfwpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Csvserver_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
