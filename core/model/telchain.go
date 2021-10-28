package model

// TelChainInitInfo 链基础信息
type TelChainInitInfo struct {
	ChainID         int
	ChainCode       string
	TotalNodesCount int
	Generator       string
	Period          int    `json:"period"`         // 出块时间
	RegulatoryData  string `json:"regulatoryData"` // 监管账号

	WitnessNum int
}

type Dpos struct {
	Period       int       `json:"period"`
	WitnessesNum int       `json:"witnessesNum"`
	Witnesses    []Witness `json:"witnesses"`
}

type Witness struct {
	Address string `json:"address"`
	Url     string `json:"url"`
}

type Config struct {
	ChainId   int    `json:"chainId"`
	Chaincode string `json:"chaincode"`
	Dpos      `json:"dpos"`
}

type GenesisAccount struct {
	Balance string `json:"balance"`
}

type Alloc map[string]GenesisAccount

type Genesis struct {
	Config         `json:"config"`
	Timestamp      string `json:"timestamp"`
	ExtraData      string `json:"extraData"`
	Generator      string `json:"generator"`
	RegulatoryData string `json:"regulatoryData"`
	Alloc          `json:"alloc"`
	Number         string `json:"number"`
	ParentHash     string `json:"parentHash"`
}

type TelChain struct {
	ChainID   int
	ChainName string
	ChainDesc string

	NameSpace string

	ChainConfigPath string
	K8sConfigPath   string
	ChainDataFPath  string

	TelChainInitInfo
}

type TelChainDeployInfo struct {
	LeagueID  int
	ChainID   uint64
	NodeId    int
	NodeCount int
}

type TelChainResource struct {
	LeagueID int
	ChainID  uint64
	NodeId   int
	CPU      float64 `json:"cpu"`
	Memory   int     `json:"memory"`
	Storage  string
}
