package model

import (
	"fmt"
	"strings"
)

type PeerK8sConfig struct {
	Cpu          float64 `json:"cpu"`
	Memory       string  `json:"memory"`
	EndpointPort string  `json:"endpointPort"`
	WorkerLabel  string  `json:"workerLabel"`
	ClPort       string  `json:"clPort"`
	CouchdbPort  string  `json:"couchdbPort"`
	OutputPort   string  `json:"outputPort"`
	IP           string  `json:"ip"`
	Nodeid       string  `json:"nodeId"`
	Org          string  `json:"org"`
	PodName      string  `json:"podName"`
	Volume       Volumes `json:"volumes"`
}

type Volumes struct {
	ArtifactPV        string `json:"artifactPv"`
	ArtifactStorage   string `json:"artifactStorage"`
	DataPV            string `json:"dataPv"`
	DataStorage       string `json:"dataStorage"`
	ArtifactServiceIP string `json:"artifactServiceIp"`
	ArtifactMountPath string `json:"artifactMountPath"`
	DataServiceIP     string `json:"dataServiceIp"`
	DataMountPath     string `json:"dataMountPath"`
}

type PeerOrgConfig struct {
	Peercount       int    `json:"peerCount"`
	DeactivatedNode []int  `json:"deactivatedNode"`
	Country         string `json:"country"`
	Province        string `json:"province"`
	Locality        string `json:"locality"`
	Usercount       int    `json:"userCount"`
	PeerK8sconfig   map[int]*PeerK8sConfig
	CaK8s           CaK8sConfig `json:"caK8sConfig"`
}

type OrdererConfig struct {
	Cpu          float64 `json:"cpu"`
	Memory       string  `json:"memory"`
	EndPointPort string  `json:"endpointPort"`
	OutputPort   string  `json:"outputPort"`
	WorkerLabel  string  `json:"workerLabel"`
	IP           string  `json:"ip"`
	Nodeid       string  `json:"nodeId"`
	PodName      string  `json:"podName"`
	Volume       Volumes `json:"volumes"`
}

type OrdererOrgConfig struct {
	Country          string `json:"country"`
	Province         string `json:"province"`
	Locality         string `json:"locality"`
	DeactivatedNode  []int  `json:"deactivatedNode"`
	OrdererK8sConfig map[int]OrdererConfig
}

type CaK8sConfig struct {
	Cpu          float64 `json:"cpu"`
	Memory       string  `json:"memory"`
	EndPointPort string  `json:"endpointPort"`
}

type FabricChain struct {
	ChainId    int      `json:"chainId"`    //用户帐号
	LeagueId   int      `json:"leagueId"`   //用户帐号
	Consensus  string   `json:"consensus"`  //共识
	PeersOrgs  []string `json:"peersOrgs"`  //参与组织 除了orderer
	OrderCount int      `json:"orderCount"` //orderer节点个数
	//PeerCount   map[string]int      `json:"peerCount"`   //每个组织节点个数
	PeerOrgConfig    map[string]*PeerOrgConfig `json:"peerOrgConfig"`
	ChannelName      string                    `json:"channelName"` //channel 名
	TlsEnabled       string                    `json:"tlsEnabled"`  //是否开启tls  true or false
	KafkaCount       int                       `json:"kafkaCount"`
	ZkCount          string                    `json:"zkCount"`
	OrdererOrgconfig OrdererOrgConfig          `json:"ordererOrgConfig"`
	//Addedpeerorg       []string `json:"addedperorg"`
	//Addedpeerconfig    map[string]*PeerOrgConfig `json:"addedpeerconfig"`
	//Addedpeer          string   `json:"addedpeer"`//该peer节点隶属的组织名
	//Addedpeercount     int `json:"addedpeercount"`//该组织名下多增加几个peer节点
	//AddedOrderer       int `json:"added_orderer"`
	//AddedOrdererConfig map[int]*OrdererConfig
}

func (f FabricChain) GetHostDomain(org string) string {
	return strings.ToLower("fabric-league-" + fmt.Sprintf("%d", f.LeagueId) + "-" + fmt.Sprintf("%d", f.ChainId) + "-" + org)
}

type FabricChannel struct {
	FabricChain    `json:"fabricChain"`
	ChaincodeId    string   `json:"chaincodeId"`
	ChaincodePath  string   `json:"chaincodePath"`
	ChaincodeBytes []byte   `json:"chaincodeBytes"`
	Version        string   `json:"version"`
	Policy         string   `json:"policy"`
	Args           [][]byte `json:"args"`
}

type AddedPeerOrg struct {
}

type AddedNode struct {
}

type DeleteNode struct {
	ChainId  int    `json:"chainId"`  //用户帐号
	LeagueId int    `json:"leagueId"` //用户帐号
	Type     string `json:"type"`     //只能为orderer或者peer
	Index    string `json:"index"`
	Org      string `json:"org"`
}

func (f FabricChannel) GetChain() *FabricChain {
	return &f.FabricChain
}
