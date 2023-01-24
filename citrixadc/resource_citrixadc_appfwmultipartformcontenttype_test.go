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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
	"net/url"
)

const testAccAppfwmultipartformcontenttype_basic = `
	resource "citrixadc_appfwmultipartformcontenttype" "tf_multipartform" {
		multipartformcontenttypevalue = "date/tf_multipartform"
		isregex                       = "REGEX"
	}
`

func TestAccAppfwmultipartformcontenttype_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwmultipartformcontenttypeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppfwmultipartformcontenttype_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwmultipartformcontenttypeExist("citrixadc_appfwmultipartformcontenttype.tf_multipartform", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwmultipartformcontenttype.tf_multipartform", "multipartformcontenttypevalue", "date/tf_multipartform"),
					resource.TestCheckResourceAttr("citrixadc_appfwmultipartformcontenttype.tf_multipartform", "isregex", "REGEX"),
				),
			},
		},
	})
}

func testAccCheckAppfwmultipartformcontenttypeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwmultipartformcontenttype name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		appfwmultipartformcontenttypeNameEscaped := url.PathEscape(url.QueryEscape(rs.Primary.ID))
		data, err := nsClient.FindResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeNameEscaped)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwmultipartformcontenttype %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwmultipartformcontenttypeDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwmultipartformcontenttype" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		
		appfwmultipartformcontenttypeNameEscaped := url.PathEscape(url.QueryEscape(rs.Primary.ID))
		_, err := nsClient.FindResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeNameEscaped)
		
		if err == nil {
			return fmt.Errorf("appfwmultipartformcontenttype %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
