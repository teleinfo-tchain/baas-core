package path

import (
	"github.com/spf13/viper"
)

func GetTelChainProjectPath(leagueID, chainID string, config *viper.Viper) string {
	chainDataPath := config.GetString("TelChainFilePath") + "/" +
		leagueID + "/" + chainID + "/chainData/"
	return chainDataPath
}
