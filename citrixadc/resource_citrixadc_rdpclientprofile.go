package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/rdp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRdpclientprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRdpclientprofileFunc,
		Read:          readRdpclientprofileFunc,
		Update:        updateRdpclientprofileFunc,
		Delete:        deleteRdpclientprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addusernameinrdpfile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"audiocapturemode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyboardhook": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multimonitorsupport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"randomizerdpfilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpcookievalidity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rdpcustomparams": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpfilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdphost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdplinkattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdplistener": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpurloverride": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectclipboard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectcomports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectdrives": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectpnpdevices": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectprinters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"videoplaybackmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRdpclientprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRdpclientprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpclientprofileName := d.Get("name").(string)
	rdpclientprofile := rdp.Rdpclientprofile{
		Addusernameinrdpfile: d.Get("addusernameinrdpfile").(string),
		Audiocapturemode:     d.Get("audiocapturemode").(string),
		Keyboardhook:         d.Get("keyboardhook").(string),
		Multimonitorsupport:  d.Get("multimonitorsupport").(string),
		Name:                 d.Get("name").(string),
		Psk:                  d.Get("psk").(string),
		Randomizerdpfilename: d.Get("randomizerdpfilename").(string),
		Rdpcookievalidity:    d.Get("rdpcookievalidity").(int),
		Rdpcustomparams:      d.Get("rdpcustomparams").(string),
		Rdpfilename:          d.Get("rdpfilename").(string),
		Rdphost:              d.Get("rdphost").(string),
		Rdplinkattribute:     d.Get("rdplinkattribute").(string),
		Rdplistener:          d.Get("rdplistener").(string),
		Rdpurloverride:       d.Get("rdpurloverride").(string),
		Redirectclipboard:    d.Get("redirectclipboard").(string),
		Redirectcomports:     d.Get("redirectcomports").(string),
		Redirectdrives:       d.Get("redirectdrives").(string),
		Redirectpnpdevices:   d.Get("redirectpnpdevices").(string),
		Redirectprinters:     d.Get("redirectprinters").(string),
		Videoplaybackmode:    d.Get("videoplaybackmode").(string),
	}

	_, err := client.AddResource("rdpclientprofile", rdpclientprofileName, &rdpclientprofile)
	if err != nil {
		return err
	}

	d.SetId(rdpclientprofileName)

	err = readRdpclientprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rdpclientprofile but we can't read it ?? %s", rdpclientprofileName)
		return nil
	}
	return nil
}

func readRdpclientprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRdpclientprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpclientprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rdpclientprofile state %s", rdpclientprofileName)
	data, err := client.FindResource("rdpclientprofile", rdpclientprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rdpclientprofile state %s", rdpclientprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("addusernameinrdpfile", data["addusernameinrdpfile"])
	d.Set("audiocapturemode", data["audiocapturemode"])
	d.Set("keyboardhook", data["keyboardhook"])
	d.Set("multimonitorsupport", data["multimonitorsupport"])
	d.Set("psk", data["psk"])
	d.Set("randomizerdpfilename", data["randomizerdpfilename"])
	d.Set("rdpcookievalidity", data["rdpcookievalidity"])
	d.Set("rdpcustomparams", data["rdpcustomparams"])
	d.Set("rdpfilename", data["rdpfilename"])
	d.Set("rdphost", data["rdphost"])
	d.Set("rdplinkattribute", data["rdplinkattribute"])
	d.Set("rdplistener", data["rdplistener"])
	d.Set("rdpurloverride", data["rdpurloverride"])
	d.Set("redirectclipboard", data["redirectclipboard"])
	d.Set("redirectcomports", data["redirectcomports"])
	d.Set("redirectdrives", data["redirectdrives"])
	d.Set("redirectpnpdevices", data["redirectpnpdevices"])
	d.Set("redirectprinters", data["redirectprinters"])
	d.Set("videoplaybackmode", data["videoplaybackmode"])

	return nil

}

func updateRdpclientprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRdpclientprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpclientprofileName := d.Get("name").(string)

	rdpclientprofile := rdp.Rdpclientprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("addusernameinrdpfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Addusernameinrdpfile has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Addusernameinrdpfile = d.Get("addusernameinrdpfile").(string)
		hasChange = true
	}
	if d.HasChange("audiocapturemode") {
		log.Printf("[DEBUG]  citrixadc-provider: Audiocapturemode has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Audiocapturemode = d.Get("audiocapturemode").(string)
		hasChange = true
	}
	if d.HasChange("keyboardhook") {
		log.Printf("[DEBUG]  citrixadc-provider: Keyboardhook has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Keyboardhook = d.Get("keyboardhook").(string)
		hasChange = true
	}
	if d.HasChange("multimonitorsupport") {
		log.Printf("[DEBUG]  citrixadc-provider: Multimonitorsupport has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Multimonitorsupport = d.Get("multimonitorsupport").(string)
		hasChange = true
	}
	if d.HasChange("psk") {
		log.Printf("[DEBUG]  citrixadc-provider: Psk has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Psk = d.Get("psk").(string)
		hasChange = true
	}
	if d.HasChange("randomizerdpfilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Randomizerdpfilename has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Randomizerdpfilename = d.Get("randomizerdpfilename").(string)
		hasChange = true
	}
	if d.HasChange("rdpcookievalidity") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpcookievalidity has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Rdpcookievalidity = d.Get("rdpcookievalidity").(int)
		hasChange = true
	}
	if d.HasChange("rdpcustomparams") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpcustomparams has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Rdpcustomparams = d.Get("rdpcustomparams").(string)
		hasChange = true
	}
	if d.HasChange("rdpfilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpfilename has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Rdpfilename = d.Get("rdpfilename").(string)
		hasChange = true
	}
	if d.HasChange("rdphost") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdphost has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Rdphost = d.Get("rdphost").(string)
		hasChange = true
	}
	if d.HasChange("rdplinkattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdplinkattribute has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Rdplinkattribute = d.Get("rdplinkattribute").(string)
		hasChange = true
	}
	if d.HasChange("rdplistener") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdplistener has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Rdplistener = d.Get("rdplistener").(string)
		hasChange = true
	}
	if d.HasChange("rdpurloverride") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpurloverride has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Rdpurloverride = d.Get("rdpurloverride").(string)
		hasChange = true
	}
	if d.HasChange("redirectclipboard") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectclipboard has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Redirectclipboard = d.Get("redirectclipboard").(string)
		hasChange = true
	}
	if d.HasChange("redirectcomports") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectcomports has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Redirectcomports = d.Get("redirectcomports").(string)
		hasChange = true
	}
	if d.HasChange("redirectdrives") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectdrives has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Redirectdrives = d.Get("redirectdrives").(string)
		hasChange = true
	}
	if d.HasChange("redirectpnpdevices") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectpnpdevices has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Redirectpnpdevices = d.Get("redirectpnpdevices").(string)
		hasChange = true
	}
	if d.HasChange("redirectprinters") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectprinters has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Redirectprinters = d.Get("redirectprinters").(string)
		hasChange = true
	}
	if d.HasChange("videoplaybackmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Videoplaybackmode has changed for rdpclientprofile %s, starting update", rdpclientprofileName)
		rdpclientprofile.Videoplaybackmode = d.Get("videoplaybackmode").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("rdpclientprofile", &rdpclientprofile)
		if err != nil {
			return fmt.Errorf("Error updating rdpclientprofile %s", rdpclientprofileName)
		}
	}
	return readRdpclientprofileFunc(d, meta)
}

func deleteRdpclientprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRdpclientprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpclientprofileName := d.Id()
	err := client.DeleteResource("rdpclientprofile", rdpclientprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
