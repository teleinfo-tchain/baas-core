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

type TelChainInfo struct {
	LeagueID  int    `json:"leagueId"`
	ChainID   uint64 `json:"chainId"`
	NodeIds   []int  `json:"nodeIds,omitempty"`
	NodeCount int    `json:"nodeCount,omitempty"`
}

type TelChainNodeInfo struct {
	LeagueID int    `json:"leagueId"`
	ChainID  uint64 `json:"chainId"`
	NodeId   int    `json:"nodeId,omitempty"`
}

type TelChainResource struct {
	LeagueID int     `json:"leagueId"`
	ChainID  uint64  `json:"chainId"`
	NodeId   int     `json:"nodeID,omitempty"`
	CPU      float64 `json:"cpu,omitempty"`
	Memory   int     `json:"memory,omitempty"`
	Storage  string  `json:"storage,omitempty"`
}
