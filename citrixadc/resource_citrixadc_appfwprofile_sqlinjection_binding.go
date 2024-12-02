package citrixadc

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwprofile_sqlinjection_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_sqlinjection_bindingFunc,
		Read:          readAppfwprofile_sqlinjection_bindingFunc,
		Delete:        deleteAppfwprofile_sqlinjection_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sqlinjection": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"as_scan_location_sql": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"formactionurl_sql": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alertonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_expr_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isregex_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_sqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appFwName := d.Get("name").(string)
	sqlinjection := d.Get("sqlinjection").(string)
	bindingId := fmt.Sprintf("%s,%s", appFwName, sqlinjection)

	appfwprofile_sqlinjection_binding := appfw.Appfwprofilesqlinjectionbinding{
		Alertonly:         d.Get("alertonly").(string),
		Asscanlocationsql: d.Get("as_scan_location_sql").(string),
		Asvalueexprsql:    d.Get("as_value_expr_sql").(string),
		Asvaluetypesql:    d.Get("as_value_type_sql").(string),
		Comment:           d.Get("comment").(string),
		Formactionurlsql:  d.Get("formactionurl_sql").(string),
		Isautodeployed:    d.Get("isautodeployed").(string),
		Isregexsql:        d.Get("isregex_sql").(string),
		Isvalueregexsql:   d.Get("isvalueregex_sql").(string),
		Name:              appFwName,
		Sqlinjection:      sqlinjection,
		State:             d.Get("state").(string),
	}

	_, err := client.AddResource(service.Appfwprofile_sqlinjection_binding.Type(), sqlinjection, &appfwprofile_sqlinjection_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_sqlinjection_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_sqlinjection_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_sqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readAppfwprofile_sqlinjection_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.SplitN(bindingId, ",", 2)
	appFwName := idSlice[0]
	sqlinjection := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_sqlinjection_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_sqlinjection_binding.Type(),
		ResourceName:             appFwName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_sqlinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["sqlinjection"].(string) == sqlinjection {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_sqlinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough
	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_sql", data["as_scan_location_sql"])
	d.Set("as_value_expr_sql", data["as_value_expr_sql"])
	d.Set("as_value_type_sql", data["as_value_type_sql"])
	d.Set("comment", data["comment"])
	d.Set("formactionurl_sql", data["formactionurl_sql"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_sql", data["isregex_sql"])
	d.Set("isvalueregex_sql", data["isvalueregex_sql"])
	d.Set("name", data["name"])
	d.Set("sqlinjection", data["sqlinjection"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_sqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	appFwName := idSlice[0]
	sqlinjection := idSlice[1]

	args := make(map[string]string)
	args["sqlinjection"] = sqlinjection
	args["formactionurl_sql"] = url.QueryEscape(d.Get("formactionurl_sql").(string))
	args["as_scan_location_sql"] = d.Get("as_scan_location_sql").(string)
	if val, ok := d.GetOk("as_value_type_sql"); ok {
		args["as_value_type_sql"] = val.(string)
	}
	if val, ok := d.GetOk("as_value_expr_sql"); ok {
		args["as_value_expr_sql"] = val.(string)
	}

	if val, ok := d.GetOk("as_value_type_sql"); ok {
		args["as_value_type_sql"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("as_value_expr_sql"); ok {
		args["as_value_expr_sql"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args["ruletype"] = url.QueryEscape(val.(string))
	}

	err := client.DeleteResourceWithArgsMap(service.Appfwprofile_sqlinjection_binding.Type(), appFwName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
