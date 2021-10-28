package model

import (
	"fmt"
	"strings"
)

type PeerK8sConfig struct {
	Cpu          string `json:"cpu"`
	Memory       string `json:"memory"`
	EndpointPort string `json:"endpoint_port"`
	WorkerLabel  string `json:"worker_label"`
	ClPort       string  `json:"cl_port"`
	CouchdbPort  string  `json:"couchdb_port"`
	OutputPort   string  `json:"output_port"`
	IP           string   `json:"ip"`
	Nodeid       string   `json:"nodeid"`
	Org          string   `json:"org"`
	PodName      string   `json:"pod_name"`
	Volume       Volumes
}

type Volumes struct {
	ArtifactPV         string `json:"artifact_pv"`
	ArtifactStorage    string `json:"artifact_storage"`
	DataPV             string  `json:"data_pv"`
	DataStorage        string  `json:"data_storage"`
	ArtifactServiceIP  string `json:"artifact_service_ip"`
	ArtifactMountPath  string `json:"artifact_mount_path"`
	DataServiceIP      string `json:"data_service_ip"`
	DataMountPath      string `json:"data_mount_path"`
}


type PeerOrgConfig struct {
	Peercount      int
	Country        string
	Province       string
	Locality       string
	Usercount      int
	PeerK8sconfig  map[int]*PeerK8sConfig
	CaK8s          CaK8sConfig
}

type OrdererConfig struct {
	Cpu            string
	Memory         string
	EndPointPort   string
	OutputPort     string
	WorkerLabel    string
	IP             string
	Nodeid         string
	PodName        string
	Volume         Volumes
}

type OrdererOrgConfig struct {
	Country        string
	Province       string
	Locality       string
	OrdererK8sConfig  map[int]OrdererConfig
}

type CaK8sConfig struct {
	Cpu            string
	Memory         string
	EndPointPort   string
}

type FabricChain struct {
	ChainId     int      `json:"ChainId"`     //用户帐号
	LeagueId    int      `json:"LeagueId"`     //用户帐号
	Consensus   string   `json:"consensus"`   //共识
	PeersOrgs   []string `json:"peersOrgs"`   //参与组织 除了orderer
	OrderCount  int      `json:"orderCount"`  //orderer节点个数
	//PeerCount   map[string]int      `json:"peerCount"`   //每个组织节点个数
	PeerOrgConfig      map[string]*PeerOrgConfig
	ChannelName        string   `json:"channelName"` //channel 名
	TlsEnabled         string   `json:"tlsEnabled"`  //是否开启tls  true or false
	KafkaCount         int       `json:"kafkaCount"`
	OrdererOrgconfig   OrdererOrgConfig
	//Addedpeerorg       []string `json:"addedperorg"`
	//Addedpeerconfig    map[string]*PeerOrgConfig `json:"addedpeerconfig"`
	//Addedpeer          string   `json:"addedpeer"`//该peer节点隶属的组织名
	//Addedpeercount     int `json:"addedpeercount"`//该组织名下多增加几个peer节点
	//AddedOrderer       int `json:"added_orderer"`
	//AddedOrdererConfig map[int]*OrdererConfig
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


type AddedPeerOrg struct {

}

type AddedNode struct {

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
