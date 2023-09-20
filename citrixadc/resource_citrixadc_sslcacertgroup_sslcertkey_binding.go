package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslcacertgroup_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcacertgroup_sslcertkey_bindingFunc,
		Read:          readSslcacertgroup_sslcertkey_bindingFunc,
		Delete:        deleteSslcacertgroup_sslcertkey_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cacertgroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certkeyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"crlcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocspcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcacertgroup_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcacertgroup_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	cacertgroupname := d.Get("cacertgroupname")
	certkeyname := d.Get("certkeyname")
	bindingId := fmt.Sprintf("%s,%s", cacertgroupname, certkeyname)
	sslcacertgroup_sslcertkey_binding := ssl.Sslcacertgroupsslcertkeybinding{
		Cacertgroupname: d.Get("cacertgroupname").(string),
		Certkeyname:     d.Get("certkeyname").(string),
		Crlcheck:        d.Get("crlcheck").(string),
		Ocspcheck:       d.Get("ocspcheck").(string),
	}

	_, err := client.AddResource("sslcacertgroup_sslcertkey_binding", bindingId, &sslcacertgroup_sslcertkey_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslcacertgroup_sslcertkey_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslcacertgroup_sslcertkey_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslcacertgroup_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcacertgroup_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	cacertgroupname := idSlice[0]
	certkeyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslcacertgroup_sslcertkey_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslcacertgroup_sslcertkey_binding",
		ResourceName:             cacertgroupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslcacertgroup_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["certkeyname"].(string) == certkeyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams certkeyname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslcacertgroup_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("cacertgroupname", data["cacertgroupname"])
	d.Set("certkeyname", data["certkeyname"])
	d.Set("crlcheck", data["crlcheck"])
	d.Set("ocspcheck", data["ocspcheck"])

	return nil

}

func deleteSslcacertgroup_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcacertgroup_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	certkeyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("certkeyname:%s", certkeyname))

	err := client.DeleteResourceWithArgs("sslcacertgroup_sslcertkey_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
