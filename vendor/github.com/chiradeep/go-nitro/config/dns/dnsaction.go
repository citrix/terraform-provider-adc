package dns

type Dnsaction struct {
	Actionname       string      `json:"actionname,omitempty"`
	Actiontype       string      `json:"actiontype,omitempty"`
	Builtin          interface{} `json:"builtin,omitempty"`
	Cachebypass      string      `json:"cachebypass,omitempty"`
	Drop             string      `json:"drop,omitempty"`
	Ipaddress        interface{} `json:"ipaddress,omitempty"`
	Preferredloclist interface{} `json:"preferredloclist,omitempty"`
	Ttl              int         `json:"ttl,omitempty"`
	Viewname         string      `json:"viewname,omitempty"`
}
