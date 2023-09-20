package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ipsec"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcIpsecprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIpsecprofileFunc,
		Read:          readIpsecprofileFunc,
		Delete:        deleteIpsecprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"encalgo": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hashalgo": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ikeretryinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ikeversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"livenesscheckinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peerpublickey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"perfectforwardsecrecy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"privatekey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"psk": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publickey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"replaywindowsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retransmissiontime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createIpsecprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsecprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecprofileName := d.Get("name").(string)

	ipsecprofile := ipsec.Ipsecprofile{
		Encalgo:               toStringList(d.Get("encalgo").([]interface{})),
		Hashalgo:              toStringList(d.Get("hashalgo").([]interface{})),
		Ikeretryinterval:      d.Get("ikeretryinterval").(int),
		Ikeversion:            d.Get("ikeversion").(string),
		Lifetime:              d.Get("lifetime").(int),
		Livenesscheckinterval: d.Get("livenesscheckinterval").(int),
		Name:                  d.Get("name").(string),
		Peerpublickey:         d.Get("peerpublickey").(string),
		Perfectforwardsecrecy: d.Get("perfectforwardsecrecy").(string),
		Privatekey:            d.Get("privatekey").(string),
		Psk:                   d.Get("psk").(string),
		Publickey:             d.Get("publickey").(string),
		Replaywindowsize:      d.Get("replaywindowsize").(int),
		Retransmissiontime:    d.Get("retransmissiontime").(int),
	}

	_, err := client.AddResource(service.Ipsecprofile.Type(), ipsecprofileName, &ipsecprofile)
	if err != nil {
		return err
	}

	d.SetId(ipsecprofileName)

	err = readIpsecprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ipsecprofile but we can't read it ?? %s", ipsecprofileName)
		return nil
	}
	return nil
}

func readIpsecprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsecprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ipsecprofile state %s", ipsecprofileName)
	data, err := client.FindResource(service.Ipsecprofile.Type(), ipsecprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipsecprofile state %s", ipsecprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("encalgo", data["encalgo"])
	d.Set("hashalgo", data["hashalgo"])
	d.Set("ikeretryinterval", data["ikeretryinterval"])
	d.Set("ikeversion", data["ikeversion"])
	d.Set("lifetime", data["lifetime"])
	d.Set("livenesscheckinterval", data["livenesscheckinterval"])
	d.Set("name", data["name"])
	d.Set("peerpublickey", data["peerpublickey"])
	//d.Set("perfectforwardsecrecy", data["perfectforwardsecrecy"])
	d.Set("privatekey", data["privatekey"])
	//d.Set("psk", data["psk"])
	d.Set("publickey", data["publickey"])
	d.Set("replaywindowsize", data["replaywindowsize"])
	d.Set("retransmissiontime", data["retransmissiontime"])

	return nil

}

func deleteIpsecprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsecprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecprofileName := d.Id()
	err := client.DeleteResource(service.Ipsecprofile.Type(), ipsecprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
