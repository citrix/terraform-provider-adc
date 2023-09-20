package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcServicegroup_lbmonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createServicegroup_lbmonitor_bindingFunc,
		Read:          readServicegroup_lbmonitor_bindingFunc,
		Delete:        deleteServicegroup_lbmonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"customserverid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dbsttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hashid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monitorname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"monstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nameserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passive": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"serverid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createServicegroup_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createServicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Get("servicegroupname")
	monitorName := d.Get("monitorname")
	// Use `,` as the separator since it is invalid character for servicegroup and monitor name
	servicegroupLbmonitorBindingId := fmt.Sprintf("%s,%s", servicegroupName, monitorName)

	servicegroup_lbmonitor_binding := basic.Servicegroupmonitorbinding{
		Customserverid:   d.Get("customserverid").(string),
		Dbsttl:           uint64(d.Get("dbsttl").(int)),
		Hashid:           uint32(d.Get("hashid").(int)),
		Monitorname:      d.Get("monitorname").(string),
		Monstate:         d.Get("monstate").(string),
		Nameserver:       d.Get("nameserver").(string),
		Passive:          d.Get("passive").(bool),
		Port:             int32(d.Get("port").(int)),
		Serverid:         uint32(d.Get("serverid").(int)),
		Servicegroupname: d.Get("servicegroupname").(string),
		State:            d.Get("state").(string),
		Weight:           uint32(d.Get("weight").(int)),
	}

	err := client.UpdateUnnamedResource(service.Servicegroup_lbmonitor_binding.Type(), &servicegroup_lbmonitor_binding)
	if err != nil {
		return err
	}

	d.SetId(servicegroupLbmonitorBindingId)

	err = readServicegroup_lbmonitor_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this servicegroup_lbmonitor_binding but we can't read it ?? %s", servicegroupLbmonitorBindingId)
		return nil
	}
	return nil
}

func readServicegroup_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readServicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupLbmonitorBindingId := d.Id()
	idSlice := strings.Split(servicegroupLbmonitorBindingId, ",")

	if len(idSlice) < 2 {
		return fmt.Errorf("Cannot deduce monitorname from id string")
	}

	if len(idSlice) > 2 {
		return fmt.Errorf("Too many separators \",\" in id string")
	}

	servicegroupName := idSlice[0]
	monitorName := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading servicegroup_lbmonitor_binding state %s", servicegroupLbmonitorBindingId)
	findParams := service.FindParams{
		ResourceType:             "servicegroup_lbmonitor_binding",
		ResourceName:             servicegroupName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_lbmonitor_binding state %s", servicegroupLbmonitorBindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["monitor_name"].(string) == monitorName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_lbmonitor_binding state %s", servicegroupLbmonitorBindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("customserverid", data["customserverid"])
	d.Set("dbsttl", data["dbsttl"])
	d.Set("hashid", data["hashid"])
	d.Set("monitorname", data["monitor_name"])
	d.Set("monstate", data["monstate"])
	d.Set("nameserver", data["nameserver"])
	d.Set("passive", data["passive"])
	d.Set("port", data["port"])
	d.Set("serverid", data["serverid"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("state", data["state"])
	d.Set("weight", data["weight"])

	return nil

}

func deleteServicegroup_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteServicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupLbmonitorBindingId := d.Id()
	idSlice := strings.Split(servicegroupLbmonitorBindingId, ",")

	servicegroupName := idSlice[0]
	monitorName := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("monitor_name:%s", monitorName))
	if v, ok := d.GetOk("port"); ok {
		args = append(args, fmt.Sprintf("port:%v", v))
	}

	err := client.DeleteResourceWithArgs(service.Servicegroup_lbmonitor_binding.Type(), servicegroupName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
