package citrixadc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func dataSourceCitrixAdcNitroInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCitrixAdcNitroInfoRead,
		Schema: map[string]*schema.Schema{
			"workflow": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"query_args": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"primary_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nitro_list": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"object": {
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"nitro_object": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceCitrixAdcNitroInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In dataSourceCitrixAdcNitroInfoRead")
	var err error
	workflowMap := d.Get("workflow").(map[string]interface{})
	switch workflowMap["lifecycle"].(string) {
	case "binding_list":
		err = dataSourceCitrixAdcNitroInfoBindingListRead(d, meta)
	case "object_by_name":
		err = dataSourceCitrixAdcNitroInfoObjectByNameRead(d, meta)
	default:
		err = fmt.Errorf("Lifecycle %s is not implemented", workflowMap["lifecycle"].(string))
	}

	return err
}

func dataSourceCitrixAdcNitroInfoBindingListRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In dataSourceCitrixAdcNitroInfoBindingListRead")
	client := meta.(*NetScalerNitroClient).client
	workflowMap := d.Get("workflow").(map[string]interface{})
	primaryId := d.Get("primary_id").(string)
	argsMap := make(map[string]string)
	queryArgs := d.Get("query_args").(map[string]interface{})
	for k, v := range queryArgs {
		argsMap[k] = v.(string)
	}
	missingErrorCode, err := strconv.Atoi(workflowMap["bound_resource_missing_errorcode"].(string))
	if err != nil {
		return err
	}
	findParams := service.FindParams{
		ResourceType:             workflowMap["endpoint"].(string),
		ResourceName:             primaryId,
		ResourceMissingErrorCode: missingErrorCode,
		ArgsMap:                  argsMap,
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("dataArr %v ", dataArr)
	if err != nil {
		if strings.Contains(err.Error(), workflowMap["bound_resource_missing_errorcode"].(string)) {
			id := resource.PrefixedUniqueId("nitro-info-")
			d.SetId(id)
			emptyList := make([]interface{}, 0)
			d.Set("nitro_list", emptyList)
			return nil
		} else {
			return err
		}
	}
	// Fallthrough

	id := resource.PrefixedUniqueId("nitro-info-")
	d.SetId(id)

	output_list := make([]interface{}, 0)
	for _, item := range dataArr {
		object_map := make(map[string]interface{})
		item_map := make(map[string]string, 0)
		for k, v := range item {
			item_map[k] = fmt.Sprintf("%v", v)
		}
		object_map["object"] = item_map
		output_list = append(output_list, object_map)

	}
	log.Printf("output_list %v", output_list)
	d.Set("nitro_list", output_list)

	return nil
}

func dataSourceCitrixAdcNitroInfoObjectByNameRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In dataSourceCitrixAdcNitroInfoObjectByNameRead")
	client := meta.(*NetScalerNitroClient).client
	workflowMap := d.Get("workflow").(map[string]interface{})
	primaryId := d.Get("primary_id").(string)
	argsMap := make(map[string]string)
	queryArgs := d.Get("query_args").(map[string]interface{})
	for k, v := range queryArgs {
		argsMap[k] = v.(string)
	}
	missingErrorCode, err := strconv.Atoi(workflowMap["bound_resource_missing_errorcode"].(string))
	if err != nil {
		return err
	}
	findParams := service.FindParams{
		ResourceType:             workflowMap["endpoint"].(string),
		ResourceName:             primaryId,
		ResourceMissingErrorCode: missingErrorCode,
		ArgsMap:                  argsMap,
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("dataArr %v ", dataArr)
	if err != nil {
		if strings.Contains(err.Error(), workflowMap["bound_resource_missing_errorcode"].(string)) {
			id := resource.PrefixedUniqueId("nitro-info-")
			d.SetId(id)
			emptyMap := make(map[string]interface{})
			d.Set("nitro_object", emptyMap)
			return nil
		} else {
			return err
		}
	}
	// Fallthrough

	if len(dataArr) > 1 {
		return fmt.Errorf("Too many results %d", len(dataArr))
	}

	id := resource.PrefixedUniqueId("nitro-info-")
	d.SetId(id)

	outputMap := make(map[string]string)

	if len(dataArr) == 0 {
		d.Set("nitro_object", outputMap)
		return nil
	}
	// Fallthrough

	data := dataArr[0]
	for k, v := range data {
		outputMap[k] = fmt.Sprintf("%v", v)

	}
	d.Set("nitro_object", outputMap)

	return nil
}
