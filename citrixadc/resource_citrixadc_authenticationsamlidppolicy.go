package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationsamlidppolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationsamlidppolicyFunc,
		Read:          readAuthenticationsamlidppolicyFunc,
		Update:        updateAuthenticationsamlidppolicyFunc,
		Delete:        deleteAuthenticationsamlidppolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationsamlidppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationsamlidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidppolicyName := d.Get("name").(string)
	authenticationsamlidppolicy := authentication.Authenticationsamlidppolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Authenticationsamlidppolicy.Type(), authenticationsamlidppolicyName, &authenticationsamlidppolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationsamlidppolicyName)

	err = readAuthenticationsamlidppolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationsamlidppolicy but we can't read it ?? %s", authenticationsamlidppolicyName)
		return nil
	}
	return nil
}

func readAuthenticationsamlidppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationsamlidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidppolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationsamlidppolicy state %s", authenticationsamlidppolicyName)
	data, err := client.FindResource(service.Authenticationsamlidppolicy.Type(), authenticationsamlidppolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationsamlidppolicy state %s", authenticationsamlidppolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateAuthenticationsamlidppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationsamlidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidppolicyName := d.Get("name").(string)

	authenticationsamlidppolicy := authentication.Authenticationsamlidppolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authenticationsamlidppolicy %s, starting update", authenticationsamlidppolicyName)
		authenticationsamlidppolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for authenticationsamlidppolicy %s, starting update", authenticationsamlidppolicyName)
		authenticationsamlidppolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for authenticationsamlidppolicy %s, starting update", authenticationsamlidppolicyName)
		authenticationsamlidppolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for authenticationsamlidppolicy %s, starting update", authenticationsamlidppolicyName)
		authenticationsamlidppolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationsamlidppolicy %s, starting update", authenticationsamlidppolicyName)
		authenticationsamlidppolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for authenticationsamlidppolicy %s, starting update", authenticationsamlidppolicyName)
		authenticationsamlidppolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationsamlidppolicy.Type(), authenticationsamlidppolicyName, &authenticationsamlidppolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationsamlidppolicy %s", authenticationsamlidppolicyName)
		}
	}
	return readAuthenticationsamlidppolicyFunc(d, meta)
}

func deleteAuthenticationsamlidppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationsamlidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidppolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationsamlidppolicy.Type(), authenticationsamlidppolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
