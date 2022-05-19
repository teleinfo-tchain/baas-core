package model

type K8sData struct {
	Data [][]byte `json:"data"`
}

type ChainDomain struct {
	NodeIps   []string          `json:"nodeIps"`
	NodePorts map[string]string `json:"nodePorts"`
}

type HostNode struct {
	HostName string `json:"hostName"`
	HostIP   string `json:"hostIP"`
	Cpu      string `json:"cpu"   `
	Memory   string `json:"memory"`
	Storage  int64  `json:"storage"`
}

type ChainPod struct {
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
	Name       string `json:"name"`
	Port       int32  `json:"port"`
	URL        string `json:"url"`
	HostIP     string `json:"hostIP"`
	Type       string `json:"type"`
	Cpu        string `json:"cpu"`
	Memory     string `json:"memory"`
	Storage    string `json:"storage"`
}

type Resources struct {
	ChainId  int     `json:"chainId"`
	LeagueId int     `json:"leagueId"`
	Type     string  `json:"type"`
	Index    string  `json:"index"`
	Org      string  `json:"org"`
	CPU      float64 `json:"cpu"`
	Memory   int     `json:"memory"`
}
