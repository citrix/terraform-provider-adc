package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslrsakey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslrsakeyFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"bits": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aes256": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			}, "des": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"des3": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"exponent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"keyfile": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"keyform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"pkcs8": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  false,
				ForceNew:  true,
				Sensitive: true,
			},
		},
	}
}

func createSslrsakeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslrsakeyFunc")
	client := meta.(*NetScalerNitroClient).client

	sslrsakeyName := resource.PrefixedUniqueId("tf-sslrsakey-")
	sslrsakey := ssl.Sslrsakey{
		Bits:     d.Get("bits").(int),
		Des:      d.Get("des").(bool),
		Des3:     d.Get("des3").(bool),
		Aes256:   d.Get("aes256").(bool),
		Pkcs8:    d.Get("pkcs8").(bool),
		Password: d.Get("password").(string),
		Exponent: d.Get("exponent").(string),
		Keyfile:  d.Get("keyfile").(string),
		Keyform:  d.Get("keyform").(string),
	}

	err := client.ActOnResource(service.Sslrsakey.Type(), &sslrsakey, "create")
	if err != nil {
		return err
	}

	d.SetId(sslrsakeyName)

	return nil
}
