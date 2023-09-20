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
	"os"
	"testing"
)

func TestAccPasswordresetter_basic(t *testing.T) {
	if os.Getenv("ADC_TEST_PASSWORD_RESETTER") == "" {
		t.Skip("Did not detect flag variable ADC_TEST_PASSWORD_RESETTER")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPasswordResetter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPasswordResetterExist("citrixadc_password_resetter.tf_resetter", nil),
				),
			},
		},
	})
}

func testAccCheckPasswordResetterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No syncer id is set")
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

const testAccPasswordResetter_basic = `
resource "citrixadc_password_resetter" "tf_resetter" {
    username = "nsroot"
    password = "nsroot"
    new_password = "newnsroot"
}
`
