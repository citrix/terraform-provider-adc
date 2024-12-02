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

const testAccAppfwprofile_sqlinjection_binding_basic = `
	resource "citrixadc_appfwprofile_sqlinjection_binding" "demo_binding" {
		name                 = citrixadc_appfwprofile.demo_appfw.name
		sqlinjection         = "demo_binding"
		as_scan_location_sql = "HEADER"
		formactionurl_sql    = "www.example.com"
		state                = "ENABLED"
		isregex_sql          = "NOTREGEX"
		as_value_type_sql    = "Keyword"
		as_value_expr_sql    = "example1"
		isvalueregex_sql     = "NOTREGEX"
	}

	resource "citrixadc_appfwprofile" "demo_appfw" {
		name                     = "demo_appfwprofile"
		bufferoverflowaction     = ["none"]
		contenttypeaction        = ["none"]
		cookieconsistencyaction  = ["none"]
		creditcard               = ["none"]
		creditcardaction         = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction            = ["none"]
		denyurlaction            = ["none"]
		dynamiclearning          = ["none"]
		fieldconsistencyaction   = ["none"]
		fieldformataction        = ["none"]
		fileuploadtypesaction    = ["none"]
		inspectcontenttypes      = ["none"]
		jsondosaction            = ["none"]
		jsonsqlinjectionaction   = ["none"]
		jsonxssaction            = ["none"]
		multipleheaderaction     = ["none"]
		sqlinjectionaction       = ["none"]
		starturlaction           = ["none"]
		type                     = ["HTML"]
		xmlattachmentaction      = ["none"]
		xmldosaction             = ["none"]
		xmlformataction          = ["none"]
		xmlsoapfaultaction       = ["none"]
		xmlsqlinjectionaction    = ["none"]
		xmlvalidationaction      = ["none"]
		xmlwsiaction             = ["none"]
		xmlxssaction             = ["none"]
	}
`

const testAccAppfwprofile_sqlinjection_binding_as_value_type_sql = `
	resource citrixadc_appfwprofile_sqlinjection_binding demo_as_value_type_sql_binding {
		name = citrixadc_appfwprofile.demo_appfw_as_value_type_sql.name
		sqlinjection= "demo_binding"
		as_scan_location_sql= "HEADER"
		as_value_type_sql= "Keyword"
		as_value_expr_sql= "example1"
		formactionurl_sql= "www.example.com"
		state= "ENABLED"
		isregex_sql= "NOTREGEX"
	}

	resource citrixadc_appfwprofile demo_appfw_as_value_type_sql {
		name = "demo_appfw_as_value_type_sql"
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
`

func TestAccAppfwprofile_sqlinjection_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_sqlinjection_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_sqlinjection_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_sqlinjection_bindingExist("citrixadc_appfwprofile_sqlinjection_binding.demo_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_binding", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_binding", "as_scan_location_sql", "HEADER"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_binding", "formactionurl_sql", "www.example.com"),
				),
			},
			{
				Config: testAccAppfwprofile_sqlinjection_binding_as_value_type_sql,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_sqlinjection_bindingExist("citrixadc_appfwprofile_sqlinjection_binding.demo_as_value_type_sql_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_as_value_type_sql_binding", "name", "demo_appfw_as_value_type_sql"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_as_value_type_sql_binding", "as_scan_location_sql", "HEADER"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_as_value_type_sql_binding", "formactionurl_sql", "www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_as_value_type_sql_binding", "as_value_type_sql", "Keyword"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.demo_as_value_type_sql_binding", "as_value_expr_sql", "example1"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_sqlinjection_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_sqlinjection_binding name is set")
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
		appFwName := idSlice[0]
		sqlinjection := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_sqlinjection_binding.Type(),
			ResourceName:             appFwName,
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
			if v["sqlinjection"].(string) == sqlinjection {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find appfwprofile_sqlinjection_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_sqlinjection_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_sqlinjection_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwprofile_sqlinjection_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_sqlinjection_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
