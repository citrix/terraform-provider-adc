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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccAppfwprofile_denyurl_binding_basic = `
	resource citrixadc_appfwprofile demo_appfw {
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

	resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl1 {
		name = citrixadc_appfwprofile.demo_appfw.name
		denyurl = "debug[.][^/?]*(|[?].*)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}
`

func TestAccAppfwprofile_denyurl_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_denyurl_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppfwprofile_denyurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_denyurl_bindingExist("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "denyurl", "debug[.][^/?]*(|[?].*)$"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_denyurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_denyurl_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingID := rs.Primary.ID
		idSlice := strings.SplitN(bindingID, ",", 2)

		if len(idSlice) < 2 {
			return fmt.Errorf("Cannot deduce appfwprofile and denyurl from ID string")
		}

		profileName := idSlice[0]
		denyURL := idSlice[1]

		findParams := service.FindParams{
			ResourceType: service.Appfwprofile_denyurl_binding.Type(),
			ResourceName: profileName,
		}
		findParams.FilterMap = make(map[string]string)
		findParams.FilterMap["denyurl"] = url.QueryEscape(denyURL)
		data, err := nsClient.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwprofile_denyurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_denyurl_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_denyurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwprofile_denyurl_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_denyurl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
