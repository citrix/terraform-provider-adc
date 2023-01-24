package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSnmpengineid() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpengineidFunc,
		Read:          readSnmpengineidFunc,
		Update:        updateSnmpengineidFunc,
		Delete:        deleteSnmpengineidFunc, // Thought snmpengineid resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"engineid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ownernode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
func createSnmpengineidFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpengineidFunc")
	client := meta.(*NetScalerNitroClient).client
	
	snmpengineidName := resource.PrefixedUniqueId("tf-snmpengineid-")

	
	snmpengineid := snmp.Snmpengineid{
			Engineid:  d.Get("engineid").(string),
			Ownernode: d.Get("ownernode").(int),
	}

	err := client.UpdateUnnamedResource(service.Snmpengineid.Type(), &snmpengineid)
	if err != nil {
			return err
	}

	d.SetId(snmpengineidName)

	err = readSnmpengineidFunc(d, meta)
	if err != nil {
			log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpengineid but we can't read it ??", )
			return nil
	}
	return nil
}

func readSnmpengineidFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpengineidFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpengineid state",)
	data, err := client.FindResource(service.Snmpengineid.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpengineid state")
		d.SetId("")
		return nil
	}
	d.Set("engineid", data["engineid"])
	d.Set("ownernode", data["ownernode"])

	return nil

}

func updateSnmpengineidFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpengineidFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpengineid := snmp.Snmpengineid{}
	
	hasChange := false
	if d.HasChange("engineid") {
		log.Printf("[DEBUG]  citrixadc-provider: Engineid has changed for snmpengineid, starting update")
		snmpengineid.Engineid = d.Get("engineid").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for snmpengineid, starting update")
		snmpengineid.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpengineid.Type(), &snmpengineid)
		if err != nil {
			return fmt.Errorf("Error updating snmpengineid")
		}
	}
	return readSnmpengineidFunc(d, meta)
}

func deleteSnmpengineidFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpengineidFunc")
	// snmpenigneid  does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}