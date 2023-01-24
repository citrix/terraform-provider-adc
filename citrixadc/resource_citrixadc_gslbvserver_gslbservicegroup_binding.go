package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcGslbvserver_gslbservicegroup_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbvserver_gslbservicegroup_bindingFunc,
		Read:          readGslbvserver_gslbservicegroup_bindingFunc,
		Delete:        deleteGslbvserver_gslbservicegroup_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createGslbvserver_gslbservicegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbvserver_gslbservicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	servicegroupname := d.Get("servicegroupname")
	
	bindingId := fmt.Sprintf("%s,%s", name, servicegroupname)
	gslbvserver_gslbservicegroup_binding := gslb.Gslbvservergslbservicegroupbinding{
		Name:             d.Get("name").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
	}

	err := client.UpdateUnnamedResource("gslbvserver_gslbservicegroup_binding", &gslbvserver_gslbservicegroup_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbvserver_gslbservicegroup_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbvserver_gslbservicegroup_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbvserver_gslbservicegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbvserver_gslbservicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicegroupname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbvserver_gslbservicegroup_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbvserver_gslbservicegroup_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_gslbservicegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicegroupname"].(string) == servicegroupname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_gslbservicegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("servicegroupname", data["servicegroupname"])

	return nil

}

func deleteGslbvserver_gslbservicegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbvserver_gslbservicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicegroupname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicegroupname:%s", servicegroupname))

	err := client.DeleteResourceWithArgs("gslbvserver_gslbservicegroup_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
