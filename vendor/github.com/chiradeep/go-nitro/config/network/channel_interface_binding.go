package network

type Channelinterfacebinding struct {
	Id           string      `json:"id,omitempty"`
	Ifnum        interface{} `json:"ifnum,omitempty"`
	Lamode       string      `json:"lamode,omitempty"`
	Slaveduplex  int         `json:"slaveduplex,omitempty"`
	Slaveflowctl int         `json:"slaveflowctl,omitempty"`
	Slavemedia   int         `json:"slavemedia,omitempty"`
	Slavespeed   int         `json:"slavespeed,omitempty"`
	Slavestate   int         `json:"slavestate,omitempty"`
	Slavetime    int         `json:"slavetime,omitempty"`
}
