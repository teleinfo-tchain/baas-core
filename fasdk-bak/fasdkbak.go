package fasdk_bak

import (
	"github.com/bif/telbaas/baas-core/common/log"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
)

const (
	Admin = "Admin"
	User  = "User1"
)

var logger = log.GetLogger("fasdk", log.ERROR)

type FabricClient struct {
	ConnectionFile []byte
	OrdererDomain  string
	Orgs           []string
	OrgAdmin       string
	UserName       string
	ChannelId      string
	GoPath         string

	resmgmtClients []*resmgmt.Client
	sdk            *fabsdk.FabricSDK
	retry          resmgmt.RequestOption
	orderer        resmgmt.RequestOption
}

func NewFabricClient(connectionFile []byte, channelId string, orgs []string, orderer string) *FabricClient {
	fabric := &FabricClient{
		ConnectionFile: connectionFile,
		ChannelId:      channelId,
		OrdererDomain:  orderer,
		Orgs:           orgs,
		OrgAdmin:       Admin,
		UserName:       User,
		GoPath:         os.Getenv("GOPATH"),
	}

	return fabric

}

func (f *FabricClient) Close() {
	if f.sdk != nil {
		f.sdk.Close()
	}
}

func (f *FabricClient) Setup() error {
	sdk, err := fabsdk.New(config.FromRaw(f.ConnectionFile, "yaml"))
	if err != nil {
		logger.Error("failed to create SDK")
		return err
	}
	f.sdk = sdk

	resmgmtClients := make([]*resmgmt.Client, 0)
	for _, v := range f.Orgs {
		resmgmtClient, err := resmgmt.New(sdk.Context(fabsdk.WithUser(f.OrgAdmin), fabsdk.WithOrg(v)))
		if err != nil {
			logger.Errorf("Failed to create channel management client: %s", err)
		}
		resmgmtClients = append(resmgmtClients, resmgmtClient)
	}
	f.resmgmtClients = resmgmtClients

	f.retry = resmgmt.WithRetry(retry.DefaultResMgmtOpts)
	f.orderer = resmgmt.WithOrdererEndpoint(f.OrdererDomain)

	return nil
}

func (f *FabricClient) QueryLedger() (*fab.BlockchainInfoResponse, error) {

	ledger, err := ledger.New(f.sdk.ChannelContext(f.ChannelId, fabsdk.WithUser(f.UserName), fabsdk.WithOrg(f.Orgs[0])))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	bci, err := ledger.QueryInfo()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return bci, nil
}
