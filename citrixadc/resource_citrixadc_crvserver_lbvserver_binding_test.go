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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

const testAccCrvserver_lbvserver_binding_basic = `

resource "citrixadc_crvserver" "crvserver" {
	name        = "my_vserver"
	servicetype = "HTTP"
	arp         = "OFF"
  }
  resource "citrixadc_lbvserver" "foo_lbvserver" {
	name        = "test_lbvserver"
	servicetype = "HTTP"
	ipv46       = "192.0.0.0"
	port        = 8000
	comment     = "hello"
  }
  resource "citrixadc_service" "tf_service" {
	lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
	name        = "tf_service"
	port        = 8081
	ip          = "10.33.4.5"
	servicetype = "HTTP"
	cachetype   = "TRANSPARENT"
  }
  resource "citrixadc_crvserver_lbvserver_binding" "crvserver_lbvserver_binding" {
	name      = citrixadc_crvserver.crvserver.name
	lbvserver = citrixadc_lbvserver.foo_lbvserver.name
	depends_on = [
	  citrixadc_service.tf_service
	]
  }
`

const testAccCrvserver_lbvserver_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_crvserver" "crvserver" {
		name        = "my_vserver"
		servicetype = "HTTP"
		arp         = "OFF"
	  }
	  resource "citrixadc_lbvserver" "foo_lbvserver" {
		name        = "test_lbvserver"
		servicetype = "HTTP"
		ipv46       = "192.0.0.0"
		port        = 8000
		comment     = "hello"
	  }
	  resource "citrixadc_service" "tf_service" {
		lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
		name        = "tf_service"
		port        = 8081
		ip          = "10.33.4.5"
		servicetype = "HTTP"
		cachetype   = "TRANSPARENT"
	  }
`

func TestAccCrvserver_lbvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCrvserver_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCrvserver_lbvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_lbvserver_bindingExist("citrixadc_crvserver_lbvserver_binding.crvserver_lbvserver_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccCrvserver_lbvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_lbvserver_bindingNotExist("citrixadc_crvserver_lbvserver_binding.crvserver_lbvserver_binding", "my_vserver,test_lbvserver"),
				),
			},
		},
	})
}

func testAccCheckCrvserver_lbvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver_lbvserver_binding id is set")
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
		lbvserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "crvserver_lbvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching lbvserver
		found := false
		for _, v := range dataArr {
			if v["lbvserver"].(string) == lbvserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("crvserver_lbvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_lbvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		lbvserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "crvserver_lbvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching lbvserver
		found := false
		for _, v := range dataArr {
			if v["lbvserver"].(string) == lbvserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("crvserver_lbvserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_lbvserver_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver_lbvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Crvserver_lbvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver_lbvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
