package ns

type Nsparam struct {
	Aftpallowrandomsourceport string      `json:"aftpallowrandomsourceport,omitempty"`
	Cip                       string      `json:"cip,omitempty"`
	Cipheader                 string      `json:"cipheader,omitempty"`
	Cookieversion             string      `json:"cookieversion,omitempty"`
	Crportrange               string      `json:"crportrange,omitempty"`
	Exclusivequotamaxclient   int         `json:"exclusivequotamaxclient,omitempty"`
	Exclusivequotaspillover   int         `json:"exclusivequotaspillover,omitempty"`
	Ftpportrange              string      `json:"ftpportrange,omitempty"`
	Grantquotamaxclient       int         `json:"grantquotamaxclient,omitempty"`
	Grantquotaspillover       int         `json:"grantquotaspillover,omitempty"`
	Httpport                  interface{} `json:"httpport,omitempty"`
	Icaports                  interface{} `json:"icaports,omitempty"`
	Internaluserlogin         string      `json:"internaluserlogin,omitempty"`
	Maxconn                   int         `json:"maxconn,omitempty"`
	Maxreq                    int         `json:"maxreq,omitempty"`
	Pmtumin                   int         `json:"pmtumin,omitempty"`
	Pmtutimeout               int         `json:"pmtutimeout,omitempty"`
	Securecookie              string      `json:"securecookie,omitempty"`
	Timezone                  string      `json:"timezone,omitempty"`
	Useproxyport              string      `json:"useproxyport,omitempty"`
}
