package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/system"

	"github.com/chiradeep/go-nitro/netscaler"
	_ "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"path"
)

func resourceCitrixAdcSystemfile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemfileFunc,
		Read:          readSystemfileFunc,
		Delete:        deleteSystemfileFunc,
		Schema: map[string]*schema.Schema{
			"filecontent": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fileencoding": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Default:  "BASE64",
				Optional: true,
			},
			"filelocation": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSystemfileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemfileFunc")
	client := meta.(*NetScalerNitroClient).client

	filecontent := d.Get("filecontent").(string)
	fileencoding := d.Get("fileencoding").(string)
	filelocation := d.Get("filelocation").(string)
	filename := d.Get("filename").(string)

	fullPath := path.Join(filelocation, filename)

	if fileencoding != "BASE64" {
		return fmt.Errorf("file encoding %s is not supported", fileencoding)
	}

	// Encode file contents to base64
	b64filecontent := base64.StdEncoding.EncodeToString([]byte(filecontent))

	systemfile := system.Systemfile{
		Filecontent:  b64filecontent,
		Fileencoding: fileencoding,
		Filelocation: filelocation,
		Filename:     filename,
	}

	_, err := client.AddResource(netscaler.Systemfile.Type(), "", &systemfile)
	if err != nil {
		return err
	}

	d.SetId(fullPath)

	err = readSystemfileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemfile but we can't read it ?? %s", fullPath)
		return err
	}
	return nil
}

func readSystemfileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemfileFunc")
	client := meta.(*NetScalerNitroClient).client
	systemfileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systemfile state %s", systemfileName)
	argsMap := make(map[string]string)
	var err error
	argsMap["filelocation"] = url.QueryEscape(d.Get("filelocation").(string))
	if err != nil {
		return err
	}
	argsMap["filename"] = url.QueryEscape(d.Get("filename").(string))
	if err != nil {
		return err
	}
	findParams := netscaler.FindParams{
		ResourceType: "systemfile",
		ArgsMap:      argsMap,
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}

	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: systemfile does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	if len(dataArray) > 1 {
		return fmt.Errorf("multiple entries found for file")
	}

	data := dataArray[0]

	bytes, err := base64.StdEncoding.DecodeString(data["filecontent"].(string))
	if err != nil {
		return err
	}

	d.Set("filecontent", string(bytes))
	d.Set("fileencoding", data["fileencoding"])
	d.Set("filelocation", data["filelocation"])
	d.Set("filename", data["filename"])

	return nil

}

func deleteSystemfileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemfileFunc")
	client := meta.(*NetScalerNitroClient).client
	argsMap := make(map[string]string)
	argsMap["filelocation"] = url.QueryEscape(d.Get("filelocation").(string))
	filename := url.PathEscape(d.Get("filename").(string))
	err := client.DeleteResourceWithArgsMap("systemfile", filename, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
