package tm

type Tmglobaltmsessionpolicybinding struct {
	Bindpolicytype int         `json:"bindpolicytype,omitempty"`
	Builtin        interface{} `json:"builtin,omitempty"`
	Policyname     string      `json:"policyname,omitempty"`
	Priority       int         `json:"priority,omitempty"`
}
