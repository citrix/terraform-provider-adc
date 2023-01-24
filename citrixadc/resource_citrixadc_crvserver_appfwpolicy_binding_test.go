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

const testAccCrvserver_appfwpolicy_binding_basic = `

resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_appfwprofile" "demo_appfwprofile" {
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

resource "citrixadc_appfwpolicy" "demo_appfwpolicy1" {
    name = "demo_appfwpolicy1"
    profilename = citrixadc_appfwprofile.demo_appfwprofile.name
    rule = "true"
}

resource "citrixadc_crvserver_appfwpolicy_binding" "crvserver_appfwpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_appfwpolicy.demo_appfwpolicy1.name
    priority = 20
}
`

const testAccCrvserver_appfwpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_appfwprofile" "demo_appfwprofile" {
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
	
	resource "citrixadc_appfwpolicy" "demo_appfwpolicy1" {
		name = "demo_appfwpolicy1"
		profilename = citrixadc_appfwprofile.demo_appfwprofile.name
		rule = "true"
	}
`

func TestAccCrvserver_appfwpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCrvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCrvserver_appfwpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appfwpolicy_bindingExist("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccCrvserver_appfwpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appfwpolicy_bindingNotExist("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", "my_vserver,demo_appfwpolicy1"),
				),
			},
		},
	})
}

func testAccCheckCrvserver_appfwpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver_appfwpolicy_binding id is set")
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

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "crvserver_appfwpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("crvserver_appfwpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_appfwpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "crvserver_appfwpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("crvserver_appfwpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_appfwpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Crvserver_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
