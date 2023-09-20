package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcLinkset_channel_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLinkset_channel_bindingFunc,
		Read:          readLinkset_channel_bindingFunc,
		Delete:        deleteLinkset_channel_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"linkset_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ifnum": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLinkset_channel_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLinkset_channel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	id := d.Get("linkset_id").(string)
	ifnum := d.Get("ifnum").(string)
	bindingId := fmt.Sprintf("%s,%s", id, ifnum)
	linkset_channel_binding := network.Linksetchannelbinding{
		Id:    d.Get("linkset_id").(string),
		Ifnum: d.Get("ifnum").(string),
	}

	err := client.UpdateUnnamedResource(service.Linkset_channel_binding.Type(), &linkset_channel_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLinkset_channel_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this linkset_channel_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLinkset_channel_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLinkset_channel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	id := idSlice[0]
	ifnum := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading linkset_channel_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "linkset_channel_binding",
		ResourceName:             url.QueryEscape(url.QueryEscape(id)),
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
		log.Printf("[WARN] citrixadc-provider: Clearing linkset_channel_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ifnum"].(string) == ifnum {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ifnum not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing linkset_channel_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("linkset_id", data["id"])
	d.Set("ifnum", data["ifnum"])

	return nil

}

func deleteLinkset_channel_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLinkset_channel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ifnum := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ifnum:%v", url.QueryEscape(ifnum)))

	err := client.DeleteResourceWithArgs(service.Linkset_channel_binding.Type(), url.QueryEscape(url.QueryEscape(name)), args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
