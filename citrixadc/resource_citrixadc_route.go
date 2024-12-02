package citrixadc

import (
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcRoute() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRouteFunc,
		Read:          readRouteFunc,
		Update:        updateRouteFunc,
		Delete:        deleteRouteFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advertise": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cost": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cost1": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detail": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"distance": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"msr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	// setting ID as `network__netmask__gateway` for better logging
	routeName := d.Get("network").(string) + "__" + d.Get("netmask").(string) + "__" + d.Get("gateway").(string)

	route := network.Route{
		Advertise:  d.Get("advertise").(string),
		Cost:       d.Get("cost").(int),
		Cost1:      d.Get("cost1").(int),
		Detail:     d.Get("detail").(bool),
		Distance:   d.Get("distance").(int),
		Gateway:    d.Get("gateway").(string),
		Monitor:    d.Get("monitor").(string),
		Msr:        d.Get("msr").(string),
		Netmask:    d.Get("netmask").(string),
		Network:    d.Get("network").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Routetype:  d.Get("routetype").(string),
		Td:         d.Get("td").(int),
		Vlan:       d.Get("vlan").(int),
		Weight:     d.Get("weight").(int),
	}

	_, err := client.AddResource(service.Route.Type(), route.Network, &route)
	if err != nil {
		return err
	}

	d.SetId(routeName)

	err = readRouteFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this route but we can't read it ?? %s", routeName)
		return nil
	}
	return nil
}

func readRouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	routeName := d.Id()
	if routeName != "" {
		ip_net_gate := strings.Split(routeName, "__")
		if len(ip_net_gate) == 3 {
			d.Set("network", ip_net_gate[0])
			d.Set("netmask", ip_net_gate[1])
			d.Set("gateway", ip_net_gate[2])
		}
	}
	log.Printf("[DEBUG] citrixadc-provider: Reading route state %s", routeName)
	findParams := service.FindParams{
		ResourceType: service.Route.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing route state %s", routeName)
		d.SetId("")
		return nil
	}
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: route does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	idSlice := strings.SplitN(routeName, "__", 3)

	network := idSlice[0]
	netmask := idSlice[1]
	gateway := idSlice[2]

	foundIndex := -1
	for i, route := range dataArray {
		match := true
		if route["network"] != network {
			match = false
		}
		if route["netmask"] != netmask {
			match = false
		}
		if route["gateway"] != gateway {
			match = false
		}
		if val, ok := d.GetOk("ownergroup"); ok {
			if route["ownergroup"] != val.(string) {
				match = false
			}
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams route not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing route state %s", routeName)
		d.SetId("")
		return nil
	}

	data := dataArray[foundIndex]

	d.Set("advertise", data["advertise"])
	d.Set("cost", data["cost"])
	d.Set("cost1", data["cost1"])
	d.Set("detail", data["detail"])
	d.Set("distance", data["distance"])
	d.Set("gateway", data["gateway"])
	d.Set("monitor", data["monitor"])
	d.Set("msr", data["msr"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("routetype", data["routetype"])
	d.Set("td", data["td"])
	d.Set("vlan", data["vlan"])
	setToInt("weight", d, data["weight"])

	return nil

}

func updateRouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	routeName := d.Id()
	if routeName != "" {
		ip_net_gate := strings.Split(routeName, "__")
		if len(ip_net_gate) == 3 {
			d.Set("network", ip_net_gate[0])
			d.Set("netmask", ip_net_gate[1])
			d.Set("gateway", ip_net_gate[2])
		}
	}
	route := network.Route{}
	hasChange := false
	if d.HasChange("advertise") {
		log.Printf("[DEBUG]  citrixadc-provider: Advertise has changed for route %s, starting update", routeName)
		route.Advertise = d.Get("advertise").(string)
		hasChange = true
	}
	if d.HasChange("cost") {
		log.Printf("[DEBUG]  citrixadc-provider: Cost has changed for route %s, starting update", routeName)
		route.Cost = d.Get("cost").(int)
		hasChange = true
	}
	if d.HasChange("cost1") {
		log.Printf("[DEBUG]  citrixadc-provider: Cost1 has changed for route %s, starting update", routeName)
		route.Cost1 = d.Get("cost1").(int)
		hasChange = true
	}
	if d.HasChange("detail") {
		log.Printf("[DEBUG]  citrixadc-provider: Detail has changed for route %s, starting update", routeName)
		route.Detail = d.Get("detail").(bool)
		hasChange = true
	}
	if d.HasChange("distance") {
		log.Printf("[DEBUG]  citrixadc-provider: Distance has changed for route %s, starting update", routeName)
		route.Distance = d.Get("distance").(int)
		hasChange = true
	}
	if d.HasChange("gateway") {
		log.Printf("[DEBUG]  citrixadc-provider: Gateway has changed for route %s, starting update", routeName)
		route.Gateway = d.Get("gateway").(string)
		hasChange = true
	}
	if d.HasChange("monitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitor has changed for route %s, starting update", routeName)
		route.Monitor = d.Get("monitor").(string)
		hasChange = true
	}
	if d.HasChange("msr") {
		log.Printf("[DEBUG]  citrixadc-provider: Msr has changed for route %s, starting update", routeName)
		route.Msr = d.Get("msr").(string)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for route %s, starting update", routeName)
		route.Netmask = d.Get("netmask").(string)
		hasChange = true
	}
	if d.HasChange("network") {
		log.Printf("[DEBUG]  citrixadc-provider: Network has changed for route %s, starting update", routeName)
		route.Network = d.Get("network").(string)
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for route %s, starting update", routeName)
		route.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("routetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Routetype has changed for route %s, starting update", routeName)
		route.Routetype = d.Get("routetype").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for route %s, starting update", routeName)
		route.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for route %s, starting update", routeName)
		route.Vlan = d.Get("vlan").(int)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for route %s, starting update", routeName)
		route.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		// network, netmask, gateway attributes are mandatory in UPDATE data
		route.Network = d.Get("network").(string)
		route.Netmask = d.Get("netmask").(string)
		route.Gateway = d.Get("gateway").(string)
		err := client.UpdateUnnamedResource(service.Route.Type(), &route)
		if err != nil {
			return fmt.Errorf("Error updating route %s", routeName)
		}
	}
	return readRouteFunc(d, meta)
}

func deleteRouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRouteFunc")
	client := meta.(*NetScalerNitroClient).client

	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(d.Get("network").(string))
	argsMap["netmask"] = url.QueryEscape(d.Get("netmask").(string))
	argsMap["gateway"] = url.QueryEscape(d.Get("gateway").(string))
	if val, ok := d.GetOk("ownergroup"); ok {
		argsMap["ownergroup"] = url.QueryEscape(val.(string))
	}

	err := client.DeleteResourceWithArgsMap(service.Route.Type(), "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
