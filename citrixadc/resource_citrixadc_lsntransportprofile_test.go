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
)

const testAccLsntransportprofile_basic = `

	resource "citrixadc_lsntransportprofile" "tf_lsntransportprofile" {
		transportprofilename = "my_lsn_transportprofile"
		transportprotocol    = "TCP"
		portquota            = 10
		sessionquota         = 10
		groupsessionlimit    = 100
	}
  
`

const testAccLsntransportprofile_update = `

	resource "citrixadc_lsntransportprofile" "tf_lsntransportprofile" {
		transportprofilename = "my_lsn_transportprofile"
		transportprotocol    = "TCP"
		portquota            = 20
		sessionquota         = 20
		groupsessionlimit    = 1000
	}
  
`

func TestAccLsntransportprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLsntransportprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsntransportprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_lsntransportprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprofilename", "my_lsn_transportprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "portquota", "10"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "sessionquota", "10"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "groupsessionlimit", "100"),
				),
			},
			{
				Config: testAccLsntransportprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_lsntransportprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprofilename", "my_lsn_transportprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "portquota", "20"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "sessionquota", "20"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "groupsessionlimit", "1000"),
				),
			},
		},
	})
}

func testAccCheckLsntransportprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsntransportprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("lsntransportprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsntransportprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsntransportprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsntransportprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("lsntransportprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsntransportprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
