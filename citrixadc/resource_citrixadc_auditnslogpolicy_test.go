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
	"testing"
)

const testAccAuditnslogpolicy_basic = `


resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
	name   = "my_auditnslogpolicy"
	rule   = "true"
	action = "SETASLEARNNSLOG_ACT"
  }
  
`
const testAccAuditnslogpolicy_update = `


resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
	name   = "my_auditnslogpolicy"
	rule   = "false"
	action = "SETASLEARNNSLOG_ACT"
  }
  
`

func TestAccAuditnslogpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuditnslogpolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuditnslogpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogpolicyExist("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", "name", "my_auditnslogpolicy"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", "action", "SETASLEARNNSLOG_ACT"),
				),
			},
			resource.TestStep{
				Config: testAccAuditnslogpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogpolicyExist("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", "name", "my_auditnslogpolicy"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogpolicy.tf_auditnslogpolicy", "action", "SETASLEARNNSLOG_ACT"),
				),
			},
		},
	})
}

func testAccCheckAuditnslogpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditnslogpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Auditnslogpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("auditnslogpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditnslogpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditnslogpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Auditnslogpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("auditnslogpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
