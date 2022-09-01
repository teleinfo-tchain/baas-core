package path

import (
	"fmt"
	"github.com/bif/telbaas/baas-core/common/constant"
	"github.com/bif/telbaas/baas-core/common/fileutil"
	"github.com/bif/telbaas/baas-core/core/model"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
)

func NewUserChainPath(artifactPath, k8sConfig, dataPath, templatePath string) UserChainPath {
	return UserChainPath{
		ArtifactPath:  artifactPath,
		K8sConfigPath: k8sConfig,
		DataPath:      dataPath,
		TemplatePath:  templatePath,
	}
}

type UserChainPath struct {
	ArtifactPath  string
	K8sConfigPath string
	DataPath      string
	TemplatePath  string
}

type ProjectDir struct {
	BaasArtifactsDir       string
	BaasK8sFabricConfigDir string
	BaasFabricDataDir      string
}

func (p ProjectDir) GetProjectDir(chain *model.FabricChain, Config *viper.Viper) UserChainPath {
	artifactPath := filepath.Join(p.BaasArtifactsDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-k8s-config/[leagueid]/[chainid]/
	k8sConfig := filepath.Join(p.BaasK8sFabricConfigDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/
	dataPath := filepath.Join(p.BaasFabricDataDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	//模板
	// /baas-fabricEngine/baas-template
	templatePath := filepath.Join(Config.GetString("FabricRootPath"), Config.GetString("FabricTemplate"))
	return NewUserChainPath(artifactPath, k8sConfig, dataPath, templatePath)
}

func (p ProjectDir) GetProjectDirByNode(node *model.DeleteNode, Config *viper.Viper) UserChainPath {
	artifactPath := filepath.Join(p.BaasArtifactsDir, fmt.Sprintf("%d", node.LeagueId), fmt.Sprintf("%d", node.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-k8s-config/[leagueid]/[chainid]/
	k8sConfig := filepath.Join(p.BaasK8sFabricConfigDir, fmt.Sprintf("%d", node.LeagueId), fmt.Sprintf("%d", node.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/
	dataPath := filepath.Join(p.BaasFabricDataDir, fmt.Sprintf("%d", node.LeagueId), fmt.Sprintf("%d", node.ChainId))
	//模板
	// /baas-fabricEngine/baas-template
	templatePath := filepath.Join(Config.GetString("FabricRootPath"), Config.GetString("FabricTemplate"))
	return NewUserChainPath(artifactPath, k8sConfig, dataPath, templatePath)
}

func (p ProjectDir) BuildProjectDir(chain *model.FabricChain, Config *viper.Viper) UserChainPath {
	//nfs shared
	//改了
	// /baas-fabricEngine/baas-nfsshared/baas-artifacts/[leagueid]/[chainid]/
	artifactPath := filepath.Join(p.BaasArtifactsDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-k8s-config/[leagueid]/[chainid]/
	k8sConfig := filepath.Join(p.BaasK8sFabricConfigDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/
	dataPath := filepath.Join(p.BaasFabricDataDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	//模板
	// /baas-fabricEngine/baas-template
	templatePath := filepath.Join(Config.GetString("FabricRootPath"), Config.GetString("FabricTemplate"))

	fileutil.CreatedDir(artifactPath)
	fileutil.CreatedDir(k8sConfig)
	fileutil.CreatedDir(dataPath)
	//创建artifact文件夹
	///baas-fabricEngine/baas-nfsshared/baas-artifacts/[leagueid]/[chainid]/channel-artifacts
	fileutil.CreatedDir(filepath.Join(artifactPath, constant.ChannelArtifactsDir))

	switch chain.Consensus {
	case constant.OrdererSolo:
		domain := "orderer0." + chain.GetHostDomain(constant.OrdererSuffix)
		// /baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/order0.[Leagueid].[chainid]orderer
		fileutil.CreatedDir(filepath.Join(dataPath, domain))
	case constant.OrdererKafka:
		for i := 0; i < chain.OrderCount; i++ {
			// /baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/order[i].[Leagueid].[chainid]orderer
			domain := "orderer" + strconv.Itoa(i) + "." + chain.GetHostDomain(constant.OrdererSuffix)
			fileutil.CreatedDir(filepath.Join(dataPath, domain))
		}

		for i := 0; i < chain.KafkaCount; i++ {
			//baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/order[i].[Leagueid].[chainid]kafka
			kafka := "kafka" + strconv.Itoa(i) + "." + chain.GetHostDomain(constant.KafkaSuffix)
			fileutil.CreatedDir(filepath.Join(dataPath, kafka))
		}
	case constant.OrdererEtcdraft:
		for i := 0; i < chain.OrderCount; i++ {
			//baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/order[i].[Leagueid].[chainid]orderer
			domain := "orderer" + strconv.Itoa(i) + "." + chain.GetHostDomain(constant.OrdererSuffix)
			fileutil.CreatedDir(filepath.Join(dataPath, domain))
		}
	}

	for _, o := range chain.PeersOrgs {
		//改了
		for i := 0; i < chain.PeerOrgConfig[o].Peercount; i++ {
			//baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/peer[i].[Leagueid].[chainid][org]
			domain := "peer" + strconv.Itoa(i) + "." + chain.GetHostDomain(o)
			fileutil.CreatedDir(filepath.Join(dataPath, domain))
			//baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/couchdb.peer[i].[Leagueid].[chainid][org]
			fileutil.CreatedDir(filepath.Join(dataPath, "couchdb."+domain))
		}
	}

	return NewUserChainPath(artifactPath, k8sConfig, dataPath, templatePath)
}

func (p ProjectDir) RemoveProjectDir(chain *model.FabricChain) error {
	//改了
	artifactPath := filepath.Join(p.BaasArtifactsDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	k8sConfig := filepath.Join(p.BaasK8sFabricConfigDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	dataPath := filepath.Join(p.BaasFabricDataDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))

	if fileutil.RemoveDir(artifactPath) && fileutil.RemoveDir(k8sConfig) && fileutil.RemoveDir(dataPath) {
		return nil
	}
	return fmt.Errorf("remove project dir error")

}

func NewProjetc(Config *viper.Viper) ProjectDir {
	// 从feconfig.yaml传入 /baas-fabricEngine/baas-nfsshared
	//todo：这里以后用minio，需要改
	baasNfsSharedDir := filepath.Join(Config.GetString("FabricRootPath"), Config.GetString("FabricNfsShared"))
	baasDataDir := filepath.Join(Config.GetString("FabricDataPath"))
	// /baas-fabricEngine/baas-nfsshared/baas-artifacts
	baasArtifactsDir := filepath.Join(baasNfsSharedDir, constant.BaasArtifacts)
	// /baas-fabricEngine/baas-nfsshared/baas-k8s-config
	BaasK8sFabricConfigDir := filepath.Join(baasNfsSharedDir, constant.BaasK8sConfig)
	//baas-fabricEngine/baas-nfsshared/baas-fabric-data
	BaasFabricDataDir := filepath.Join(baasDataDir, constant.BaasFabricData)
	return ProjectDir{
		BaasArtifactsDir:       baasArtifactsDir,
		BaasK8sFabricConfigDir: BaasK8sFabricConfigDir,
		BaasFabricDataDir:      BaasFabricDataDir,
	}
}
