package model

import (
	"fmt"
	"strings"
)

type PeerConfig struct {
	Peercount int
	Country   string
	Province  string
	Locality  string
	Usercount int
}

type OrdererConfig struct {
	Country   string
	Province  string
	Locality  string
}

type FabricChain struct {
	ChainName   string   `json:"chainName"`
	ChainId     int      `json:"ChainId"`     //用户帐号
	LeagueId    int      `json:"LeagueId"`     //用户帐号
	Consensus   string   `json:"consensus"`   //共识
	PeersOrgs   []string `json:"peersOrgs"`   //参与组织 除了orderer
	OrderCount  int      `json:"orderCount"`  //orderer节点个数
	//PeerCount   map[string]int      `json:"peerCount"`   //每个组织节点个数
	Peerconfig  map[string]*PeerConfig
	ChannelName string   `json:"channelName"` //channel 名
	TlsEnabled  string   `json:"tlsEnabled"`  //是否开启tls  true or false
	KafkaCount  int       `json:"kafkaCount"`
	Ordererconfig OrdererConfig
	Addedpeerorg    []string `json:"addedperorg"`
	Addedpeerconfig  map[string]PeerConfig `json:"addedpeerconfig"`
	Addedpeer   string   `json:"addedpeer"`//该peer节点隶属的组织名
	Addedpeercount int `json:"addedpeercount"`//该组织名下多增加几个peer节点
}

func (f FabricChain) GetHostDomain(org string) string {
	return strings.ToLower(fmt.Sprintf("%d", f.LeagueId)+fmt.Sprintf("%d", f.ChainId)+ org)
}

type FabricChannel struct {
	FabricChain
	ChaincodeId    string   `json:"chaincodeId"`
	ChaincodePath  string   `json:"chaincodePath"`
	ChaincodeBytes []byte   `json:"chaincodeBytes"`
	Version        string   `json:"version"`
	Policy         string   `json:"policy"`
	Args           [][]byte `json:"args"`
}

type DeleteNode struct {
	ChainId     int      `json:"ChainId"`     //用户帐号
	LeagueId    int      `json:"LeagueId"`     //用户帐号
	Type        string   `json:"Type"`  //只能为orderer或者peer
	Index       int      `json:"Index"`
	Org         string   `json:"org"`
}

func (f FabricChannel) GetChain() *FabricChain {
	return &f.FabricChain
}
