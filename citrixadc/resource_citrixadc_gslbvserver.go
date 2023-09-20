package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/mitchellh/mapstructure"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcGslbvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbvserverFunc,
		Read:          readGslbvserverFunc,
		Update:        updateGslbvserverFunc,
		Delete:        deleteGslbvserverFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"appflowlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backuplbmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupsessiontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"backupvserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"considereffectivestate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiedomain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookietimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disableprimaryondown": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsrecordtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domainname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamicweight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ecs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ecsaddrvalidation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"edr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iptype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mir": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistenceid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"persistencetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sitedomainttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sobackupaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"somethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistence": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistencetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sothreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tolerance": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v6netmasklen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v6persistmasklen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backupip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"backupipflag": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"cookiedomain": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cookiedomainflag": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"cookietimeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"domainname": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sitedomainttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"service": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domainname": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"servicename": {
							Type:     schema.TypeString,
							Required: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func createGslbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createGslbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var gslbvserverName string
	if v, ok := d.GetOk("name"); ok {
		gslbvserverName = v.(string)
	} else {
		gslbvserverName = resource.PrefixedUniqueId("tf-gslbvserver-")
		d.Set("name", gslbvserverName)
	}
	gslbvserver := gslb.Gslbvserver{
		Appflowlog:             d.Get("appflowlog").(string),
		Backupip:               d.Get("backupip").(string),
		Backuplbmethod:         d.Get("backuplbmethod").(string),
		Backupsessiontimeout:   d.Get("backupsessiontimeout").(int),
		Backupvserver:          d.Get("backupvserver").(string),
		Comment:                d.Get("comment").(string),
		Considereffectivestate: d.Get("considereffectivestate").(string),
		Cookiedomain:           d.Get("cookiedomain").(string),
		Cookietimeout:          d.Get("cookietimeout").(int),
		Disableprimaryondown:   d.Get("disableprimaryondown").(string),
		Dnsrecordtype:          d.Get("dnsrecordtype").(string),
		Domainname:             d.Get("domainname").(string),
		Dynamicweight:          d.Get("dynamicweight").(string),
		Ecs:                    d.Get("ecs").(string),
		Ecsaddrvalidation:      d.Get("ecsaddrvalidation").(string),
		Edr:                    d.Get("edr").(string),
		Iptype:                 d.Get("iptype").(string),
		Lbmethod:               d.Get("lbmethod").(string),
		Mir:                    d.Get("mir").(string),
		Name:                   d.Get("name").(string),
		Netmask:                d.Get("netmask").(string),
		Persistenceid:          d.Get("persistenceid").(int),
		Persistencetype:        d.Get("persistencetype").(string),
		Persistmask:            d.Get("persistmask").(string),
		Servicename:            d.Get("servicename").(string),
		Servicetype:            d.Get("servicetype").(string),
		Sitedomainttl:          d.Get("sitedomainttl").(int),
		Sobackupaction:         d.Get("sobackupaction").(string),
		Somethod:               d.Get("somethod").(string),
		Sopersistence:          d.Get("sopersistence").(string),
		Sopersistencetimeout:   d.Get("sopersistencetimeout").(int),
		Sothreshold:            d.Get("sothreshold").(int),
		State:                  d.Get("state").(string),
		Timeout:                d.Get("timeout").(int),
		Tolerance:              d.Get("tolerance").(int),
		Ttl:                    d.Get("ttl").(int),
		V6netmasklen:           d.Get("v6netmasklen").(int),
		V6persistmasklen:       d.Get("v6persistmasklen").(int),
		Weight:                 d.Get("weight").(int),
	}

	_, err := client.AddResource(service.Gslbvserver.Type(), gslbvserverName, &gslbvserver)
	if err != nil {
		return err
	}

	d.SetId(gslbvserverName)
	domains := d.Get("domain").(*schema.Set).List()
	for _, val := range domains {
		domain := val.(map[string]interface{})
		err = bindDomainToVserver(gslbvserverName, domain, meta)
		if err != nil {
			return err
		}
	}

	services := d.Get("service").(*schema.Set).List()
	for _, val := range services {
		svc := val.(map[string]interface{})
		err = bindGslbServiceToVserver(gslbvserverName, svc, meta)
		if err != nil {
			return err
		}
	}

	err = readGslbvserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbvserver but we can't read it ?? %s", gslbvserverName)
		return nil
	}
	return nil
}

func readGslbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readGslbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbvserverName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading gslbvserver state %s", gslbvserverName)
	data, err := client.FindResource(service.Gslbvserver.Type(), gslbvserverName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing gslbvserver state %s", gslbvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("backupip", data["backupip"])
	d.Set("backuplbmethod", data["backuplbmethod"])
	d.Set("backupsessiontimeout", data["backupsessiontimeout"])
	d.Set("backupvserver", data["backupvserver"])
	d.Set("comment", data["comment"])
	d.Set("considereffectivestate", data["considereffectivestate"])
	d.Set("cookiedomain", data["cookiedomain"])
	d.Set("cookietimeout", data["cookietimeout"])
	d.Set("disableprimaryondown", data["disableprimaryondown"])
	d.Set("dnsrecordtype", data["dnsrecordtype"])
	d.Set("domainname", data["domainname"])
	d.Set("dynamicweight", data["dynamicweight"])
	d.Set("ecs", data["ecs"])
	d.Set("ecsaddrvalidation", data["ecsaddrvalidation"])
	d.Set("edr", data["edr"])
	d.Set("iptype", data["iptype"])
	d.Set("lbmethod", data["lbmethod"])
	d.Set("mir", data["mir"])
	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])
	d.Set("persistenceid", data["persistenceid"])
	d.Set("persistencetype", data["persistencetype"])
	d.Set("persistmask", data["persistmask"])
	d.Set("servicename", data["servicename"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sitedomainttl", data["sitedomainttl"])
	d.Set("sobackupaction", data["sobackupaction"])
	d.Set("somethod", data["somethod"])
	d.Set("sopersistence", data["sopersistence"])
	d.Set("sopersistencetimeout", data["sopersistencetimeout"])
	d.Set("sothreshold", data["sothreshold"])
	d.Set("state", data["state"])
	d.Set("timeout", data["timeout"])
	d.Set("tolerance", data["tolerance"])
	d.Set("ttl", data["ttl"])
	d.Set("v6netmasklen", data["v6netmasklen"])
	d.Set("v6persistmasklen", data["v6persistmasklen"])
	d.Set("weight", data["weight"])

	data2, _ := client.FindResourceArray(service.Gslbvserver_domain_binding.Type(), gslbvserverName)
	domainBindings := make([]map[string]interface{}, len(data2))
	for i, binding := range data2 {
		domainBindings[i] = binding
	}
	d.Set("domain", domainBindings)

	data3, _ := client.FindResourceArray(service.Gslbvserver_gslbservice_binding.Type(), gslbvserverName)
	svcBindings := make([]map[string]interface{}, len(data3))
	for i, binding := range data3 {
		svcBindings[i] = binding
	}
	d.Set("service", svcBindings)
	return nil

}

func updateGslbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateGslbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbvserverName := d.Get("name").(string)

	gslbvserver := gslb.Gslbvserver{
		Name: d.Get("name").(string),
	}
	stateChange := false
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  netscaler-provider: Appflowlog has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("backupip") {
		log.Printf("[DEBUG]  netscaler-provider: Backupip has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Backupip = d.Get("backupip").(string)
		hasChange = true
	}
	if d.HasChange("backuplbmethod") {
		log.Printf("[DEBUG]  netscaler-provider: Backuplbmethod has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Backuplbmethod = d.Get("backuplbmethod").(string)
		hasChange = true
	}
	if d.HasChange("backupsessiontimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Backupsessiontimeout has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Backupsessiontimeout = d.Get("backupsessiontimeout").(int)
		hasChange = true
	}
	if d.HasChange("backupvserver") {
		log.Printf("[DEBUG]  netscaler-provider: Backupvserver has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Backupvserver = d.Get("backupvserver").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("considereffectivestate") {
		log.Printf("[DEBUG]  netscaler-provider: Considereffectivestate has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Considereffectivestate = d.Get("considereffectivestate").(string)
		hasChange = true
	}
	if d.HasChange("cookiedomain") {
		log.Printf("[DEBUG]  netscaler-provider: Cookiedomain has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Cookiedomain = d.Get("cookiedomain").(string)
		hasChange = true
	}
	if d.HasChange("cookietimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Cookietimeout has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Cookietimeout = d.Get("cookietimeout").(int)
		hasChange = true
	}
	if d.HasChange("disableprimaryondown") {
		log.Printf("[DEBUG]  netscaler-provider: Disableprimaryondown has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
		hasChange = true
	}
	if d.HasChange("dnsrecordtype") {
		log.Printf("[DEBUG]  netscaler-provider: Dnsrecordtype has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Dnsrecordtype = d.Get("dnsrecordtype").(string)
		hasChange = true
	}
	if d.HasChange("domainname") {
		log.Printf("[DEBUG]  netscaler-provider: Domainname has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Domainname = d.Get("domainname").(string)
		hasChange = true
	}
	if d.HasChange("dynamicweight") {
		log.Printf("[DEBUG]  netscaler-provider: Dynamicweight has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Dynamicweight = d.Get("dynamicweight").(string)
		hasChange = true
	}
	if d.HasChange("ecs") {
		log.Printf("[DEBUG]  netscaler-provider: Ecs has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Ecs = d.Get("ecs").(string)
		hasChange = true
	}
	if d.HasChange("ecsaddrvalidation") {
		log.Printf("[DEBUG]  netscaler-provider: Ecsaddrvalidation has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Ecsaddrvalidation = d.Get("ecsaddrvalidation").(string)
		hasChange = true
	}
	if d.HasChange("edr") {
		log.Printf("[DEBUG]  netscaler-provider: Edr has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Edr = d.Get("edr").(string)
		hasChange = true
	}
	if d.HasChange("iptype") {
		log.Printf("[DEBUG]  netscaler-provider: Iptype has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Iptype = d.Get("iptype").(string)
		hasChange = true
	}
	if d.HasChange("lbmethod") {
		log.Printf("[DEBUG]  netscaler-provider: Lbmethod has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Lbmethod = d.Get("lbmethod").(string)
		hasChange = true
	}
	if d.HasChange("mir") {
		log.Printf("[DEBUG]  netscaler-provider: Mir has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Mir = d.Get("mir").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  netscaler-provider: Name has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  netscaler-provider: Netmask has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Netmask = d.Get("netmask").(string)
		hasChange = true
	}
	if d.HasChange("persistenceid") {
		log.Printf("[DEBUG]  netscaler-provider: Persistenceid has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Persistenceid = d.Get("persistenceid").(int)
		hasChange = true
	}
	if d.HasChange("persistencetype") {
		log.Printf("[DEBUG]  netscaler-provider: Persistencetype has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Persistencetype = d.Get("persistencetype").(string)
		hasChange = true
	}
	if d.HasChange("persistmask") {
		log.Printf("[DEBUG]  netscaler-provider: Persistmask has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Persistmask = d.Get("persistmask").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG]  netscaler-provider: Servicename has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  netscaler-provider: Servicetype has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sitedomainttl") {
		log.Printf("[DEBUG]  netscaler-provider: Sitedomainttl has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Sitedomainttl = d.Get("sitedomainttl").(int)
		hasChange = true
	}
	if d.HasChange("sobackupaction") {
		log.Printf("[DEBUG]  netscaler-provider: Sobackupaction has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Sobackupaction = d.Get("sobackupaction").(string)
		hasChange = true
	}
	if d.HasChange("somethod") {
		log.Printf("[DEBUG]  netscaler-provider: Somethod has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Somethod = d.Get("somethod").(string)
		hasChange = true
	}
	if d.HasChange("sopersistence") {
		log.Printf("[DEBUG]  netscaler-provider: Sopersistence has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Sopersistence = d.Get("sopersistence").(string)
		hasChange = true
	}
	if d.HasChange("sopersistencetimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Sopersistencetimeout has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Sopersistencetimeout = d.Get("sopersistencetimeout").(int)
		hasChange = true
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG]  netscaler-provider: Sothreshold has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Sothreshold = d.Get("sothreshold").(int)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for gslbvserver %s, starting update", gslbvserverName)
		stateChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  netscaler-provider: Timeout has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("tolerance") {
		log.Printf("[DEBUG]  netscaler-provider: Tolerance has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Tolerance = d.Get("tolerance").(int)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  netscaler-provider: Ttl has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("v6netmasklen") {
		log.Printf("[DEBUG]  netscaler-provider: V6netmasklen has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.V6netmasklen = d.Get("v6netmasklen").(int)
		hasChange = true
	}
	if d.HasChange("v6persistmasklen") {
		log.Printf("[DEBUG]  netscaler-provider: V6persistmasklen has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.V6persistmasklen = d.Get("v6persistmasklen").(int)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  netscaler-provider: Weight has changed for gslbvserver %s, starting update", gslbvserverName)
		gslbvserver.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Gslbvserver.Type(), gslbvserverName, &gslbvserver)
		if err != nil {
			return fmt.Errorf("Error updating gslbvserver %s", gslbvserverName)
		}
	}

	if d.HasChange("domain") {
		log.Printf("[DEBUG]  netscaler-provider: Domain binding has changed for gslbvserver %s, starting update", gslbvserverName)
		orig, noo := d.GetChange("domain")
		if orig == nil {
			orig = new(schema.Set)
		}
		if noo == nil {
			noo = new(schema.Set)
		}
		oset := orig.(*schema.Set)
		nset := noo.(*schema.Set)

		remove := oset.Difference(nset).List()
		add := nset.Difference(oset).List()
		log.Printf("[DEBUG]  netscaler-provider: need to remove %d domain", len(remove))
		log.Printf("[DEBUG]  netscaler-provider: need to add %d domain", len(add))

		for _, val := range remove {
			domain := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to delete domain %v", domain)
			err := unbindDomain(gslbvserverName, domain, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error deleting domain %v", domain)
			}
		}

		for _, val := range add {
			domain := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to add domain %s", domain["domainname"].(string))
			err := bindDomainToVserver(gslbvserverName, domain, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error adding domain %s", domain["domainname"].(string))
			}
		}

	}

	if d.HasChange("service") {
		log.Printf("[DEBUG]  netscaler-provider: services binding has changed for gslbvserver %s, starting update", gslbvserverName)
		orig, noo := d.GetChange("service")
		if orig == nil {
			orig = new(schema.Set)
		}
		if noo == nil {
			noo = new(schema.Set)
		}
		oset := orig.(*schema.Set)
		nset := noo.(*schema.Set)

		remove := oset.Difference(nset).List()
		add := nset.Difference(oset).List()
		log.Printf("[DEBUG]  netscaler-provider: need to remove gslb vserver binding to %d service", len(remove))
		log.Printf("[DEBUG]  netscaler-provider: need to add gslb vserver binding to %d service", len(add))

		for _, val := range remove {
			service := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to delete gslb vserver binding to service %v", service)
			err := unbindGslbService(gslbvserverName, service, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error deleting gslb vserver binding to service %v", service)
				return err
			}
		}

		for _, val := range add {
			service := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to bind service %s", service["servicename"].(string))
			err := bindGslbServiceToVserver(gslbvserverName, service, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error binding service %s", service["servicename"].(string))
				return err
			}
		}

	}
	if stateChange {
		err := doGslbvserverStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling gslb vserver %s", gslbvserverName)
		}
	}

	return readGslbvserverFunc(d, meta)
}

func deleteGslbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteGslbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbvserverName := d.Id()
	domains := d.Get("domain").(*schema.Set).List()
	for _, val := range domains {
		domain := val.(map[string]interface{})
		_ = unbindDomain(gslbvserverName, domain, meta)
	}
	err := client.DeleteResource(service.Gslbvserver.Type(), gslbvserverName)
	if err != nil {
		return err
	}

	d.SetId("")
	//domain and bindings to gslb service are automatically deleted

	return nil
}

func bindDomainToVserver(vserver string, domain map[string]interface{}, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	domainname := domain["domainname"].(string)
	binding := gslb.Gslbvserverdomainbinding{}
	mapstructure.Decode(domain, &binding)
	binding.Name = vserver
	log.Printf("[INFO] netscaler-provider:  Binding domain %s to gslb vserver %s", domainname, vserver)
	_, err := client.AddResource(service.Gslbvserver_domain_binding.Type(), domainname, &binding)

	return err
}

func unbindDomain(gslbvserverName string, domain map[string]interface{}, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	domainname := domain["domainname"].(string)
	args := map[string]string{"domainname": domainname}
	log.Printf("[INFO] netscaler-provider:  Deleting binding of domain %s to gslb vserver %s", domainname, gslbvserverName)
	return client.DeleteResourceWithArgsMap(service.Gslbvserver_domain_binding.Type(), gslbvserverName, args)
}

func bindGslbServiceToVserver(vserver string, svc map[string]interface{}, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	servicename := svc["servicename"].(string)
	binding := gslb.Gslbvserverservicebinding{}
	mapstructure.Decode(svc, &binding)
	binding.Name = vserver
	log.Printf("[INFO] netscaler-provider:  Binding svc %s to gslb vserver %s", servicename, vserver)
	_, err := client.AddResource(service.Gslbvserver_gslbservice_binding.Type(), servicename, &binding)

	return err
}

func unbindGslbService(gslbvserverName string, svc map[string]interface{}, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	servicename := svc["servicename"].(string)
	args := map[string]string{"servicename": servicename}
	log.Printf("[INFO] netscaler-provider:  Deleting binding of svc %s to gslb vserver %s", servicename, gslbvserverName)
	return client.DeleteResourceWithArgsMap(service.Gslbvserver_gslbservice_binding.Type(), gslbvserverName, args)
}

func doGslbvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doServerStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	gslbvserver := gslb.Gslbvserver{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Gslbvserver.Type(), gslbvserver, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		err := client.ActOnResource(service.Gslbvserver.Type(), gslbvserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
