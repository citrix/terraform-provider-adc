package citrixadc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

type changePasswordPayload struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	New_password string `json:"new_password,omitempty"`
}

func resourceCitrixAdcChangePassword() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createChangePassword,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"new_password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"first_time_password_reset": &schema.Schema{
				Type:      schema.TypeBool,
				Description: "Value is 'true' if the user is changing the default password, else value is 'false' if user wants to change password at any point later",
				Required:  true,
				ForceNew:  true,
			},
		},
	}
}

func createChangePassword(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createChangePassword")
	client := meta.(*NetScalerNitroClient).client
	id := resource.PrefixedUniqueId("tf-change-password-")

	payload := changePasswordPayload{
		Username:     d.Get("username").(string),
		Password:     d.Get("password").(string),
		New_password: d.Get("new_password").(string),
	}

	// first time default password resetter
	if d.Get("first_time_password_reset").(bool) == true {
		_, err := client.AddResource("login", "", &payload)
		if err != nil {
			return err
		} 
	} else {
		new_payload := changePasswordPayload{
			Username:     d.Get("username").(string),
			Password:     d.Get("new_password").(string),
		}
		err := client.UpdateUnnamedResource("systemuser", &new_payload)
		if err != nil {
			return err
		}
	}

	d.SetId(id)

	return nil
}

