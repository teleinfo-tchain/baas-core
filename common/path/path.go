package path

import (
	"fmt"
	"github.com/bif/telbaas/baas-core/common/config"
	"github.com/bif/telbaas/baas-core/common/constant"
	"github.com/bif/telbaas/baas-core/core/model"
	"path/filepath"
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

func (p ProjectDir) GetProjectDir(chain *model.FabricChain) UserChainPath {
	artifactPath := filepath.Join(p.BaasArtifactsDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-k8s-config/[leagueid]/[chainid]/
	k8sConfig := filepath.Join(p.BaasK8sFabricConfigDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	// /baas-fabricEngine/baas-nfsshared/baas-fabric-data/[leagueid]/[chainid]/
	dataPath := filepath.Join(p.BaasFabricDataDir, fmt.Sprintf("%d", chain.LeagueId), fmt.Sprintf("%d", chain.ChainId))
	//模板
	// /baas-fabricEngine/baas-template
	templatePath := filepath.Join(config.Config.GetString("BaasRootPath"), config.Config.GetString("BaasTemplate"))
	return NewUserChainPath(artifactPath, k8sConfig, dataPath, templatePath)
}

func NewProjetc() ProjectDir {
	// 从feconfig.yaml传入 /baas-fabricEngine/baas-nfsshared
	//todo：这里以后用minio，需要改
	baasNfsSharedDir := filepath.Join(config.Config.GetString("BaasRootPath"), config.Config.GetString("BaasNfsShared"))
	// /baas-fabricEngine/baas-nfsshared/baas-artifacts
	baasArtifactsDir := filepath.Join(baasNfsSharedDir, constant.BaasArtifacts)
	// /baas-fabricEngine/baas-nfsshared/baas-k8s-config
	BaasK8sFabricConfigDir := filepath.Join(baasNfsSharedDir, constant.BaasK8sConfig)
	//baas-fabricEngine/baas-nfsshared/baas-fabric-data
	BaasFabricDataDir := filepath.Join(baasNfsSharedDir, constant.BaasFabricData)
	return ProjectDir{
		BaasArtifactsDir:       baasArtifactsDir,
		BaasK8sFabricConfigDir: BaasK8sFabricConfigDir,
		BaasFabricDataDir:      BaasFabricDataDir,
	}
}






