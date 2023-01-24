package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationldapaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationldapactionFunc,
		Read:          readAuthenticationldapactionFunc,
		Update:        updateAuthenticationldapactionFunc,
		Delete:        deleteAuthenticationldapactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alternateemailattr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attributes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cloudattributes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"followreferrals": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupattrname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupnameidentifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchsubattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kbattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbinddn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbinddnpassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldaphostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldaploginname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxldapreferrals": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxnestinglevel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mssrvrecordlocation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nestedgroupextraction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"otpsecret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwdchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"referraldnslookup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requireuser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"searchfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sectype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sshpublickey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssonameattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subattributename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validateservercert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationldapactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Get("name").(string)
	authenticationldapaction := authentication.Authenticationldapaction{
		Alternateemailattr:         d.Get("alternateemailattr").(string),
		Attribute1:                 d.Get("attribute1").(string),
		Attribute10:                d.Get("attribute10").(string),
		Attribute11:                d.Get("attribute11").(string),
		Attribute12:                d.Get("attribute12").(string),
		Attribute13:                d.Get("attribute13").(string),
		Attribute14:                d.Get("attribute14").(string),
		Attribute15:                d.Get("attribute15").(string),
		Attribute16:                d.Get("attribute16").(string),
		Attribute2:                 d.Get("attribute2").(string),
		Attribute3:                 d.Get("attribute3").(string),
		Attribute4:                 d.Get("attribute4").(string),
		Attribute5:                 d.Get("attribute5").(string),
		Attribute6:                 d.Get("attribute6").(string),
		Attribute7:                 d.Get("attribute7").(string),
		Attribute8:                 d.Get("attribute8").(string),
		Attribute9:                 d.Get("attribute9").(string),
		Attributes:                 d.Get("attributes").(string),
		Authentication:             d.Get("authentication").(string),
		Authtimeout:                d.Get("authtimeout").(int),
		Cloudattributes:            d.Get("cloudattributes").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Email:                      d.Get("email").(string),
		Followreferrals:            d.Get("followreferrals").(string),
		Groupattrname:              d.Get("groupattrname").(string),
		Groupnameidentifier:        d.Get("groupnameidentifier").(string),
		Groupsearchattribute:       d.Get("groupsearchattribute").(string),
		Groupsearchfilter:          d.Get("groupsearchfilter").(string),
		Groupsearchsubattribute:    d.Get("groupsearchsubattribute").(string),
		Kbattribute:                d.Get("kbattribute").(string),
		Ldapbase:                   d.Get("ldapbase").(string),
		Ldapbinddn:                 d.Get("ldapbinddn").(string),
		Ldapbinddnpassword:         d.Get("ldapbinddnpassword").(string),
		Ldaphostname:               d.Get("ldaphostname").(string),
		Ldaploginname:              d.Get("ldaploginname").(string),
		Maxldapreferrals:           d.Get("maxldapreferrals").(int),
		Maxnestinglevel:            d.Get("maxnestinglevel").(int),
		Mssrvrecordlocation:        d.Get("mssrvrecordlocation").(string),
		Name:                       d.Get("name").(string),
		Nestedgroupextraction:      d.Get("nestedgroupextraction").(string),
		Otpsecret:                  d.Get("otpsecret").(string),
		Passwdchange:               d.Get("passwdchange").(string),
		Pushservice:                d.Get("pushservice").(string),
		Referraldnslookup:          d.Get("referraldnslookup").(string),
		Requireuser:                d.Get("requireuser").(string),
		Searchfilter:               d.Get("searchfilter").(string),
		Sectype:                    d.Get("sectype").(string),
		Serverip:                   d.Get("serverip").(string),
		Servername:                 d.Get("servername").(string),
		Serverport:                 d.Get("serverport").(int),
		Sshpublickey:               d.Get("sshpublickey").(string),
		Ssonameattribute:           d.Get("ssonameattribute").(string),
		Subattributename:           d.Get("subattributename").(string),
		Svrtype:                    d.Get("svrtype").(string),
		Validateservercert:         d.Get("validateservercert").(string),
	}

	_, err := client.AddResource(service.Authenticationldapaction.Type(), authenticationldapactionName, &authenticationldapaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationldapactionName)

	err = readAuthenticationldapactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationldapaction but we can't read it ?? %s", authenticationldapactionName)
		return nil
	}
	return nil
}

func readAuthenticationldapactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationldapaction state %s", authenticationldapactionName)
	data, err := client.FindResource(service.Authenticationldapaction.Type(), authenticationldapactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationldapaction state %s", authenticationldapactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("alternateemailattr", data["alternateemailattr"])
	d.Set("attribute1", data["attribute1"])
	d.Set("attribute10", data["attribute10"])
	d.Set("attribute11", data["attribute11"])
	d.Set("attribute12", data["attribute12"])
	d.Set("attribute13", data["attribute13"])
	d.Set("attribute14", data["attribute14"])
	d.Set("attribute15", data["attribute15"])
	d.Set("attribute16", data["attribute16"])
	d.Set("attribute2", data["attribute2"])
	d.Set("attribute3", data["attribute3"])
	d.Set("attribute4", data["attribute4"])
	d.Set("attribute5", data["attribute5"])
	d.Set("attribute6", data["attribute6"])
	d.Set("attribute7", data["attribute7"])
	d.Set("attribute8", data["attribute8"])
	d.Set("attribute9", data["attribute9"])
	d.Set("attributes", data["attributes"])
	d.Set("authentication", data["authentication"])
	d.Set("authtimeout", data["authtimeout"])
	d.Set("cloudattributes", data["cloudattributes"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("email", data["email"])
	d.Set("followreferrals", data["followreferrals"])
	d.Set("groupattrname", data["groupattrname"])
	d.Set("groupnameidentifier", data["groupnameidentifier"])
	d.Set("groupsearchattribute", data["groupsearchattribute"])
	d.Set("groupsearchfilter", data["groupsearchfilter"])
	d.Set("groupsearchsubattribute", data["groupsearchsubattribute"])
	d.Set("kbattribute", data["kbattribute"])
	d.Set("ldapbase", data["ldapbase"])
	d.Set("ldapbinddn", data["ldapbinddn"])
	d.Set("ldapbinddnpassword", data["ldapbinddnpassword"])
	d.Set("ldaphostname", data["ldaphostname"])
	d.Set("ldaploginname", data["ldaploginname"])
	d.Set("maxldapreferrals", data["maxldapreferrals"])
	d.Set("maxnestinglevel", data["maxnestinglevel"])
	d.Set("mssrvrecordlocation", data["mssrvrecordlocation"])
	d.Set("name", data["name"])
	d.Set("nestedgroupextraction", data["nestedgroupextraction"])
	d.Set("otpsecret", data["otpsecret"])
	d.Set("passwdchange", data["passwdchange"])
	d.Set("pushservice", data["pushservice"])
	d.Set("referraldnslookup", data["referraldnslookup"])
	d.Set("requireuser", data["requireuser"])
	d.Set("searchfilter", data["searchfilter"])
	d.Set("sectype", data["sectype"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])
	d.Set("serverport", data["serverport"])
	d.Set("sshpublickey", data["sshpublickey"])
	d.Set("ssonameattribute", data["ssonameattribute"])
	d.Set("subattributename", data["subattributename"])
	d.Set("svrtype", data["svrtype"])
	d.Set("validateservercert", data["validateservercert"])

	return nil

}

func updateAuthenticationldapactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Get("name").(string)

	authenticationldapaction := authentication.Authenticationldapaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("alternateemailattr") {
		log.Printf("[DEBUG]  citrixadc-provider: Alternateemailattr has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Alternateemailattr = d.Get("alternateemailattr").(string)
		hasChange = true
	}
	if d.HasChange("attribute1") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute1 = d.Get("attribute1").(string)
		hasChange = true
	}
	if d.HasChange("attribute10") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute10 = d.Get("attribute10").(string)
		hasChange = true
	}
	if d.HasChange("attribute11") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute11 = d.Get("attribute11").(string)
		hasChange = true
	}
	if d.HasChange("attribute12") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute12 = d.Get("attribute12").(string)
		hasChange = true
	}
	if d.HasChange("attribute13") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute13 = d.Get("attribute13").(string)
		hasChange = true
	}
	if d.HasChange("attribute14") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute14 = d.Get("attribute14").(string)
		hasChange = true
	}
	if d.HasChange("attribute15") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute15 = d.Get("attribute15").(string)
		hasChange = true
	}
	if d.HasChange("attribute16") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute16 = d.Get("attribute16").(string)
		hasChange = true
	}
	if d.HasChange("attribute2") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute2 = d.Get("attribute2").(string)
		hasChange = true
	}
	if d.HasChange("attribute3") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute3 = d.Get("attribute3").(string)
		hasChange = true
	}
	if d.HasChange("attribute4") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute4 = d.Get("attribute4").(string)
		hasChange = true
	}
	if d.HasChange("attribute5") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute5 = d.Get("attribute5").(string)
		hasChange = true
	}
	if d.HasChange("attribute6") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute6 = d.Get("attribute6").(string)
		hasChange = true
	}
	if d.HasChange("attribute7") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute7 = d.Get("attribute7").(string)
		hasChange = true
	}
	if d.HasChange("attribute8") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute8 = d.Get("attribute8").(string)
		hasChange = true
	}
	if d.HasChange("attribute9") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute9 = d.Get("attribute9").(string)
		hasChange = true
	}
	if d.HasChange("attributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Attributes has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attributes = d.Get("attributes").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Authtimeout = d.Get("authtimeout").(int)
		hasChange = true
	}
	if d.HasChange("cloudattributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Cloudattributes has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Cloudattributes = d.Get("cloudattributes").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("email") {
		log.Printf("[DEBUG]  citrixadc-provider: Email has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Email = d.Get("email").(string)
		hasChange = true
	}
	if d.HasChange("followreferrals") {
		log.Printf("[DEBUG]  citrixadc-provider: Followreferrals has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Followreferrals = d.Get("followreferrals").(string)
		hasChange = true
	}
	if d.HasChange("groupattrname") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupattrname has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupattrname = d.Get("groupattrname").(string)
		hasChange = true
	}
	if d.HasChange("groupnameidentifier") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupnameidentifier has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupnameidentifier = d.Get("groupnameidentifier").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupsearchattribute = d.Get("groupsearchattribute").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchfilter") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchfilter has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupsearchfilter = d.Get("groupsearchfilter").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchsubattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchsubattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupsearchsubattribute = d.Get("groupsearchsubattribute").(string)
		hasChange = true
	}
	if d.HasChange("kbattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Kbattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Kbattribute = d.Get("kbattribute").(string)
		hasChange = true
	}
	if d.HasChange("ldapbase") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbase has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldapbase = d.Get("ldapbase").(string)
		hasChange = true
	}
	if d.HasChange("ldapbinddn") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbinddn has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldapbinddn = d.Get("ldapbinddn").(string)
		hasChange = true
	}
	if d.HasChange("ldapbinddnpassword") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbinddnpassword has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldapbinddnpassword = d.Get("ldapbinddnpassword").(string)
		hasChange = true
	}
	if d.HasChange("ldaphostname") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldaphostname has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldaphostname = d.Get("ldaphostname").(string)
		hasChange = true
	}
	if d.HasChange("ldaploginname") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldaploginname has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldaploginname = d.Get("ldaploginname").(string)
		hasChange = true
	}
	if d.HasChange("maxldapreferrals") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxldapreferrals has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Maxldapreferrals = d.Get("maxldapreferrals").(int)
		hasChange = true
	}
	if d.HasChange("maxnestinglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxnestinglevel has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Maxnestinglevel = d.Get("maxnestinglevel").(int)
		hasChange = true
	}
	if d.HasChange("mssrvrecordlocation") {
		log.Printf("[DEBUG]  citrixadc-provider: Mssrvrecordlocation has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Mssrvrecordlocation = d.Get("mssrvrecordlocation").(string)
		hasChange = true
	}
	if d.HasChange("nestedgroupextraction") {
		log.Printf("[DEBUG]  citrixadc-provider: Nestedgroupextraction has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Nestedgroupextraction = d.Get("nestedgroupextraction").(string)
		hasChange = true
	}
	if d.HasChange("otpsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Otpsecret has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Otpsecret = d.Get("otpsecret").(string)
		hasChange = true
	}
	if d.HasChange("passwdchange") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdchange has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Passwdchange = d.Get("passwdchange").(string)
		hasChange = true
	}
	if d.HasChange("pushservice") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushservice has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Pushservice = d.Get("pushservice").(string)
		hasChange = true
	}
	if d.HasChange("referraldnslookup") {
		log.Printf("[DEBUG]  citrixadc-provider: Referraldnslookup has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Referraldnslookup = d.Get("referraldnslookup").(string)
		hasChange = true
	}
	if d.HasChange("requireuser") {
		log.Printf("[DEBUG]  citrixadc-provider: Requireuser has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Requireuser = d.Get("requireuser").(string)
		hasChange = true
	}
	if d.HasChange("searchfilter") {
		log.Printf("[DEBUG]  citrixadc-provider: Searchfilter has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Searchfilter = d.Get("searchfilter").(string)
		hasChange = true
	}
	if d.HasChange("sectype") {
		log.Printf("[DEBUG]  citrixadc-provider: Sectype has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Sectype = d.Get("sectype").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("sshpublickey") {
		log.Printf("[DEBUG]  citrixadc-provider: Sshpublickey has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Sshpublickey = d.Get("sshpublickey").(string)
		hasChange = true
	}
	if d.HasChange("ssonameattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssonameattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ssonameattribute = d.Get("ssonameattribute").(string)
		hasChange = true
	}
	if d.HasChange("subattributename") {
		log.Printf("[DEBUG]  citrixadc-provider: Subattributename has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Subattributename = d.Get("subattributename").(string)
		hasChange = true
	}
	if d.HasChange("svrtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Svrtype has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Svrtype = d.Get("svrtype").(string)
		hasChange = true
	}
	if d.HasChange("validateservercert") {
		log.Printf("[DEBUG]  citrixadc-provider: Validateservercert has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Validateservercert = d.Get("validateservercert").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationldapaction.Type(), authenticationldapactionName, &authenticationldapaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationldapaction %s", authenticationldapactionName)
		}
	}
	return readAuthenticationldapactionFunc(d, meta)
}

func deleteAuthenticationldapactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Id()
	err := client.DeleteResource(service.Authenticationldapaction.Type(), authenticationldapactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
