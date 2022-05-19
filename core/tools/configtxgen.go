/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bif/telbaas/baas-core/common/log"
	"github.com/bif/telbaas/baas-core/core/tools/fabric/common/configtxgen/encoder"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/common/channelconfig"
	//"github.com/hyperledger/fabric/common/tools/configtxgen/encoder"
	genesisconfig "github.com/bif/telbaas/baas-core/core/tools/fabric/common/configtxgen/localconfig"
	//genesisconfig "github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	"github.com/bif/telbaas/baas-core/core/tools/fabric/common/configtxgen/metadata"
	//"github.com/hyperledger/fabric/common/tools/configtxgen/metadata"
	"github.com/hyperledger/fabric/common/tools/protolator"
	cb "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
)

var exitCode = 0

var logger = log.GetLogger("configtxgen", log.INFO)

type Configtxgen struct {
	outputBlock                string
	outputChannelCreateTx      string
	channelCreateTxBaseProfile string
	profile                    string
	configPath                 string
	channelID                  string
	inspectBlock               string
	inspectChannelCreateTx     string
	outputAnchorPeersUpdate    string
	asOrg                      string
	printOrg                   string
	outorg                     string
}

func NewConfigtxgen() *Configtxgen {
	return &Configtxgen{}
}

func (c *Configtxgen) SetOutputBlock(outputBlock string) {
	c.outputBlock = outputBlock
}
func (c *Configtxgen) SetOutputChannelCreateTx(outputChannelCreateTx string) {
	c.outputChannelCreateTx = outputChannelCreateTx
}
func (c *Configtxgen) SetChannelCreateTxBaseProfile(channelCreateTxBaseProfile string) {
	c.channelCreateTxBaseProfile = channelCreateTxBaseProfile
}
func (c *Configtxgen) SetProfile(profile string) {
	c.profile = profile
}
func (c *Configtxgen) SetConfigPath(configPath string) {
	c.configPath = configPath
}
func (c *Configtxgen) SetChannelID(channelID string) {
	c.channelID = channelID
}
func (c *Configtxgen) SetInspectBlock(inspectBlock string) {
	c.inspectBlock = inspectBlock
}
func (c *Configtxgen) SetInspectChannelCreateTx(inspectChannelCreateTx string) {
	c.inspectChannelCreateTx = inspectChannelCreateTx
}
func (c *Configtxgen) SetOutputAnchorPeersUpdate(outputAnchorPeersUpdate string) {
	c.outputAnchorPeersUpdate = outputAnchorPeersUpdate
}
func (c *Configtxgen) SetAsOrg(asOrg string) {
	c.asOrg = asOrg
}
func (c *Configtxgen) SetPrintOrg(printOrg string) {
	c.printOrg = printOrg
}
func (c *Configtxgen) SetOutOrg(outorg string) {
	c.outorg = outorg
}

func (c *Configtxgen) doOutputBlock(config *genesisconfig.Profile, channelID string, outputBlock string) error {
	pgen := encoder.New(config)
	logger.Info("Generating genesis block")
	if config.Orderer == nil {
		return errors.Errorf("refusing to generate block which is missing orderer section")
	}
	if config.Consortiums == nil {
		logger.Warning("Genesis block does not contain a consortiums group definition.  This block cannot be used for orderer bootstrap.")
	}
	genesisBlock := pgen.GenesisBlockForChannel(channelID)
	logger.Info("Writing genesis block")
	err := ioutil.WriteFile(outputBlock, utils.MarshalOrPanic(genesisBlock), 0644)
	if err != nil {
		return fmt.Errorf("Error writing genesis block: %s", err)
	}
	return nil
}

func (c *Configtxgen) doOutputChannelCreateTx(conf, baseProfile *genesisconfig.Profile, channelID string, outputChannelCreateTx string) error {
	logger.Info("Generating new channel configtx")

	var configtx *cb.Envelope
	var err error
	if baseProfile == nil {
		configtx, err = encoder.MakeChannelCreationTransaction(channelID, nil, conf)
	} else {
		configtx, err = encoder.MakeChannelCreationTransactionWithSystemChannelContext(channelID, nil, conf, baseProfile)
	}
	if err != nil {
		return err
	}

	logger.Info("Writing new channel tx")
	logger.Info("configtx: ", configtx)
	err = ioutil.WriteFile(outputChannelCreateTx, utils.MarshalOrPanic(configtx), 0644)
	if err != nil {
		return fmt.Errorf("Error writing channel create tx: %s", err)
	}
	return nil
}

func (c *Configtxgen) doOutputAnchorPeersUpdate(conf *genesisconfig.Profile, channelID string, outputAnchorPeersUpdate string, asOrg string) error {
	logger.Info("Generating anchor peer update")
	if asOrg == "" {
		return fmt.Errorf("Must specify an organization to update the anchor peer for")
	}

	if conf.Application == nil {
		return fmt.Errorf("Cannot update anchor peers without an application section")
	}

	var org *genesisconfig.Organization
	for _, iorg := range conf.Application.Organizations {
		if iorg.Name == asOrg {
			org = iorg
		}
	}

	if org == nil {
		return fmt.Errorf("No organization name matching: %s", asOrg)
	}

	anchorPeers := make([]*pb.AnchorPeer, len(org.AnchorPeers))
	for i, anchorPeer := range org.AnchorPeers {
		anchorPeers[i] = &pb.AnchorPeer{
			Host: anchorPeer.Host,
			Port: int32(anchorPeer.Port),
		}
	}

	configUpdate := &cb.ConfigUpdate{
		ChannelId: channelID,
		WriteSet:  cb.NewConfigGroup(),
		ReadSet:   cb.NewConfigGroup(),
	}

	// Add all the existing kubeconfig to the readset
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey] = cb.NewConfigGroup()
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey].Version = 1
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey].ModPolicy = channelconfig.AdminsPolicyKey
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name] = cb.NewConfigGroup()
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Values[channelconfig.MSPKey] = &cb.ConfigValue{}
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Policies[channelconfig.ReadersPolicyKey] = &cb.ConfigPolicy{}
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Policies[channelconfig.WritersPolicyKey] = &cb.ConfigPolicy{}
	configUpdate.ReadSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Policies[channelconfig.AdminsPolicyKey] = &cb.ConfigPolicy{}

	// Add all the existing at the same versions to the writeset
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey] = cb.NewConfigGroup()
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Version = 1
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].ModPolicy = channelconfig.AdminsPolicyKey
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name] = cb.NewConfigGroup()
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Version = 1
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].ModPolicy = channelconfig.AdminsPolicyKey
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Values[channelconfig.MSPKey] = &cb.ConfigValue{}
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Policies[channelconfig.ReadersPolicyKey] = &cb.ConfigPolicy{}
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Policies[channelconfig.WritersPolicyKey] = &cb.ConfigPolicy{}
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Policies[channelconfig.AdminsPolicyKey] = &cb.ConfigPolicy{}
	configUpdate.WriteSet.Groups[channelconfig.ApplicationGroupKey].Groups[org.Name].Values[channelconfig.AnchorPeersKey] = &cb.ConfigValue{
		Value:     utils.MarshalOrPanic(channelconfig.AnchorPeersValue(anchorPeers).Value()),
		ModPolicy: channelconfig.AdminsPolicyKey,
	}

	configUpdateEnvelope := &cb.ConfigUpdateEnvelope{
		ConfigUpdate: utils.MarshalOrPanic(configUpdate),
	}

	update := &cb.Envelope{
		Payload: utils.MarshalOrPanic(&cb.Payload{
			Header: &cb.Header{
				ChannelHeader: utils.MarshalOrPanic(&cb.ChannelHeader{
					ChannelId: channelID,
					Type:      int32(cb.HeaderType_CONFIG_UPDATE),
				}),
			},
			Data: utils.MarshalOrPanic(configUpdateEnvelope),
		}),
	}

	logger.Info("Writing anchor peer update")
	err := ioutil.WriteFile(outputAnchorPeersUpdate, utils.MarshalOrPanic(update), 0644)
	if err != nil {
		return fmt.Errorf("Error writing channel anchor peer update: %s", err)
	}
	return nil
}

func (c *Configtxgen) doInspectBlock(inspectBlock string) error {
	logger.Info("Inspecting block")
	data, err := ioutil.ReadFile(inspectBlock)
	if err != nil {
		return fmt.Errorf("Could not read block %s", inspectBlock)
	}

	logger.Info("Parsing genesis block")
	block, err := utils.UnmarshalBlock(data)
	if err != nil {
		return fmt.Errorf("error unmarshaling to block: %s", err)
	}
	err = protolator.DeepMarshalJSON(os.Stdout, block)
	if err != nil {
		return fmt.Errorf("malformed block contents: %s", err)
	}
	return nil
}

func (c *Configtxgen) doInspectChannelCreateTx(inspectChannelCreateTx string) error {
	logger.Info("Inspecting transaction")
	data, err := ioutil.ReadFile(inspectChannelCreateTx)
	if err != nil {
		return fmt.Errorf("could not read channel create tx: %s", err)
	}

	logger.Info("Parsing transaction")
	env, err := utils.UnmarshalEnvelope(data)
	if err != nil {
		return fmt.Errorf("Error unmarshaling envelope: %s", err)
	}

	err = protolator.DeepMarshalJSON(os.Stdout, env)
	if err != nil {
		return fmt.Errorf("malformed transaction contents: %s", err)
	}

	return nil
}

func (c *Configtxgen) doPrintOrg(t *genesisconfig.TopLevel, printOrg, out string) error {
	// 新增组织可用
	for _, org := range t.Profiles["OrdererGenesis"].Consortiums["SampleConsortium"].Organizations {
		if org.Name == printOrg {
			og, err := encoder.NewOrdererOrgGroup(org)
			if err != nil {
				return errors.Wrapf(err, "bad org definition for org %s", org.Name)
			}
			fp, err := os.Create(out)
			if err := protolator.DeepMarshalJSON(fp, &cb.DynamicConsortiumOrgGroup{ConfigGroup: og}); err != nil {
				return errors.Wrapf(err, "malformed org definition for org: %s", org.Name)
			}
			return nil
		}
	}

	//for _, org := range t.Organizations {
	//	if org.Name == printOrg {
	//		og, err := encoder.NewOrdererOrgGroup(org)
	//		if err != nil {
	//			return errors.Wrapf(err, "bad org definition for org %s", org.Name)
	//		}
	//		fp, err := os.Create(out)
	//		if err := protolator.DeepMarshalJSON(fp, &cb.DynamicConsortiumOrgGroup{ConfigGroup: og}); err != nil {
	//			return errors.Wrapf(err, "malformed org definition for org: %s", org.Name)
	//		}
	//		return nil
	//	}
	//}

	return errors.Errorf("organization %s not found", printOrg)

}

func (c *Configtxgen) doPrintOrgTest(t *genesisconfig.TopLevel, printOrg, out string, isAddOrg bool) error {
	if isAddOrg {
		for _, org := range t.Profiles["OrdererGenesis"].Consortiums["SampleConsortium"].Organizations {
			if org.Name == printOrg {
				og, err := encoder.NewOrdererOrgGroup(org)
				if err != nil {
					return errors.Wrapf(err, "bad org definition for org %s", org.Name)
				}
				fp, err := os.Create(out)
				if err := protolator.DeepMarshalJSON(fp, &cb.DynamicConsortiumOrgGroup{ConfigGroup: og}); err != nil {
					return errors.Wrapf(err, "malformed org definition for org: %s", org.Name)
				}
				return nil
			}
		}
	} else {
		for _, org := range t.Organizations {
			if org.Name == printOrg {
				og, err := encoder.NewOrdererOrgGroup(org)
				if err != nil {
					return errors.Wrapf(err, "bad org definition for org %s", org.Name)
				}
				fp, err := os.Create(out)
				if err := protolator.DeepMarshalJSON(fp, &cb.DynamicConsortiumOrgGroup{ConfigGroup: og}); err != nil {
					return errors.Wrapf(err, "malformed org definition for org: %s", org.Name)
				}
				return nil
			}
		}
	}

	return errors.Errorf("organization %s not found", printOrg)

}

func (c *Configtxgen) Exec() error {

	if c.profile == "" {
		c.profile = genesisconfig.SampleInsecureSoloProfile
	}

	version := false

	if c.channelID == "" && (c.outputBlock != "" || c.outputChannelCreateTx != "" || c.outputAnchorPeersUpdate != "") {
		c.channelID = genesisconfig.TestChainID
		logger.Warningf("Omitting the channel ID for configtxgen for output operations is deprecated.  Explicitly passing the channel ID will be required in the future, defaulting to '%s'.", c.channelID)
	}

	// show version
	if version {
		c.printVersion()
		os.Exit(exitCode)
	}

	// don't need to panic when running via command line
	defer func() {
		if err := recover(); err != nil {
			if strings.Contains(fmt.Sprint(err), "Error reading configuration: Unsupported Config Type") {
				logger.Error("Could not find configtx.yaml. " +
					"Please make sure that FABRIC_CFG_PATH or -configPath is set to a path " +
					"which contains configtx.yaml")
				os.Exit(1)
			}
			if strings.Contains(fmt.Sprint(err), "Could not find profile") {
				logger.Error(fmt.Sprint(err) + ". " +
					"Please make sure that FABRIC_CFG_PATH or -configPath is set to a path " +
					"which contains configtx.yaml with the specified profile")
				err = errors.New(fmt.Sprint(err) + ". " +
					"Please make sure that FABRIC_CFG_PATH or -configPath is set to a path " +
					"which contains configtx.yaml with the specified profile")
				os.Exit(1)
			}
			logger.Panic(err)
		}
	}()

	logger.Info("Loading configuration")
	factory.InitFactories(nil)
	var profileConfig *genesisconfig.Profile
	var err error
	if c.outputBlock != "" || c.outputChannelCreateTx != "" || c.outputAnchorPeersUpdate != "" {
		if c.configPath != "" {
			logger.Info("Catch point01")
			profileConfig, err = genesisconfig.Load(c.profile, c.configPath)
			if err != nil {
				return err
			}
		} else {
			profileConfig, err = genesisconfig.Load(c.profile)
			if err != nil {
				return err
			}
		}
	}
	var topLevelConfig *genesisconfig.TopLevel
	if c.configPath != "" {
		topLevelConfig, err = genesisconfig.LoadTopLevel(c.configPath)
		if err != nil {
			return err
		}
	} else {
		topLevelConfig, err = genesisconfig.LoadTopLevel()
		if err != nil {
			return err
		}
	}

	var baseProfile *genesisconfig.Profile
	logger.Info("catch c.channelCreateTxBaseProfile, channelCreateTxBaseProfile: ", c.channelCreateTxBaseProfile)
	if c.channelCreateTxBaseProfile != "" {
		logger.Info("catch c.channelCreateTxBaseProfile")
		if c.outputChannelCreateTx == "" {
			logger.Warning("Specified 'channelCreateTxBaseProfile', but did not specify 'outputChannelCreateTx', 'channelCreateTxBaseProfile' will not affect output.")
		}
		if c.configPath != "" {
			baseProfile, err = genesisconfig.Load(c.channelCreateTxBaseProfile, c.configPath)
			logger.Info("catch genesisconfig.Load, baseProfile: ", baseProfile)
			if err != nil {
				return err
			}
		} else {
			baseProfile, err = genesisconfig.Load(c.channelCreateTxBaseProfile)
			if err != nil {
				return err
			}
		}
	}

	if c.outputBlock != "" {
		if err := c.doOutputBlock(profileConfig, c.channelID, c.outputBlock); err != nil {
			logger.Error("Error on outputBlock: %s", err)
			err = errors.New("Error on outputBlock: " + err.Error())
			return err
		}
	}

	logger.Info("catch c.outputChannelCreateTx, baseProfile: ", baseProfile)
	if c.outputChannelCreateTx != "" {
		if err := c.doOutputChannelCreateTx(profileConfig, baseProfile, c.channelID, c.outputChannelCreateTx); err != nil {
			logger.Error("Error on outputChannelCreateTx: %s", err)
			err = errors.New("Error on outputChannelCreateTx: " + err.Error())
			return err
		}
	}

	if c.inspectBlock != "" {
		if err := c.doInspectBlock(c.inspectBlock); err != nil {
			logger.Error("Error on inspectBlock: %s", err)
			err = errors.New("Error on inspectBlock: " + err.Error())
			return err
		}
	}

	if c.inspectChannelCreateTx != "" {
		if err := c.doInspectChannelCreateTx(c.inspectChannelCreateTx); err != nil {
			logger.Error("Error on inspectChannelCreateTx: %s", err)
			err = errors.New("Error on inspectChannelCreateTx: " + err.Error())
			return err
		}
	}

	if c.outputAnchorPeersUpdate != "" {
		if err := c.doOutputAnchorPeersUpdate(profileConfig, c.channelID, c.outputAnchorPeersUpdate, c.asOrg); err != nil {
			logger.Error("Error on inspectChannelCreateTx: %s", err)
			err = errors.New("Error on inspectChannelCreateTx: " + err.Error())

			return err
		}
	}

	if c.printOrg != "" {
		logger.Info("Catch printOrg&outorg:", c.printOrg, "--", c.outorg)
		logger.Info(len(topLevelConfig.Organizations))
		if err := c.doPrintOrg(topLevelConfig, c.printOrg, c.outorg); err != nil {
			logger.Error("Error on printOrg: %s", err)
			err = errors.New("Error on printOrg: " + err.Error())
			return err
		}
	}
	return nil
}

func (c *Configtxgen) ExecTest(isAddOrg bool) error {

	if c.profile == "" {
		c.profile = genesisconfig.SampleInsecureSoloProfile
	}

	version := false

	if c.channelID == "" && (c.outputBlock != "" || c.outputChannelCreateTx != "" || c.outputAnchorPeersUpdate != "") {
		c.channelID = genesisconfig.TestChainID
		logger.Warningf("Omitting the channel ID for configtxgen for output operations is deprecated.  Explicitly passing the channel ID will be required in the future, defaulting to '%s'.", c.channelID)
	}

	// show version
	if version {
		c.printVersion()
		os.Exit(exitCode)
	}

	// don't need to panic when running via command line
	defer func() {
		if err := recover(); err != nil {
			if strings.Contains(fmt.Sprint(err), "Error reading configuration: Unsupported Config Type") {
				logger.Error("Could not find configtx.yaml. " +
					"Please make sure that FABRIC_CFG_PATH or -configPath is set to a path " +
					"which contains configtx.yaml")
				os.Exit(1)
			}
			if strings.Contains(fmt.Sprint(err), "Could not find profile") {
				logger.Error(fmt.Sprint(err) + ". " +
					"Please make sure that FABRIC_CFG_PATH or -configPath is set to a path " +
					"which contains configtx.yaml with the specified profile")
				err = errors.New(fmt.Sprint(err) + ". " +
					"Please make sure that FABRIC_CFG_PATH or -configPath is set to a path " +
					"which contains configtx.yaml with the specified profile")
				os.Exit(1)
			}
			logger.Panic(err)
		}
	}()

	logger.Info("Loading configuration")
	factory.InitFactories(nil)
	var profileConfig *genesisconfig.Profile
	var err error
	if c.outputBlock != "" || c.outputChannelCreateTx != "" || c.outputAnchorPeersUpdate != "" {
		if c.configPath != "" {
			profileConfig, err = genesisconfig.Load(c.profile, c.configPath)
			if err != nil {
				return err
			}
		} else {
			profileConfig, err = genesisconfig.Load(c.profile)
			if err != nil {
				return err
			}
		}
	}
	var topLevelConfig *genesisconfig.TopLevel
	if c.configPath != "" {
		topLevelConfig, err = genesisconfig.LoadTopLevel(c.configPath)
		if err != nil {
			return err
		}
	} else {
		topLevelConfig, err = genesisconfig.LoadTopLevel()
		if err != nil {
			return err
		}
	}

	var baseProfile *genesisconfig.Profile
	if c.channelCreateTxBaseProfile != "" {
		if c.outputChannelCreateTx == "" {
			logger.Warning("Specified 'channelCreateTxBaseProfile', but did not specify 'outputChannelCreateTx', 'channelCreateTxBaseProfile' will not affect output.")
		}
		if c.configPath != "" {
			baseProfile, err = genesisconfig.Load(c.channelCreateTxBaseProfile, c.configPath)
			if err != nil {
				return err
			}
		} else {
			baseProfile, err = genesisconfig.Load(c.channelCreateTxBaseProfile)
			if err != nil {
				return err
			}
		}
	}

	if c.outputBlock != "" {
		if err := c.doOutputBlock(profileConfig, c.channelID, c.outputBlock); err != nil {
			logger.Error("Error on outputBlock: %s", err)
			err = errors.New("Error on outputBlock: " + err.Error())
			return err
		}
	}

	if c.outputChannelCreateTx != "" {
		if err := c.doOutputChannelCreateTx(profileConfig, baseProfile, c.channelID, c.outputChannelCreateTx); err != nil {
			logger.Error("Error on outputChannelCreateTx: %s", err)
			err = errors.New("Error on outputChannelCreateTx: " + err.Error())
			return err
		}
	}

	if c.inspectBlock != "" {
		if err := c.doInspectBlock(c.inspectBlock); err != nil {
			logger.Error("Error on inspectBlock: %s", err)
			err = errors.New("Error on inspectBlock: " + err.Error())
			return err
		}
	}

	if c.inspectChannelCreateTx != "" {
		if err := c.doInspectChannelCreateTx(c.inspectChannelCreateTx); err != nil {
			logger.Error("Error on inspectChannelCreateTx: %s", err)
			err = errors.New("Error on inspectChannelCreateTx: " + err.Error())
			return err
		}
	}

	if c.outputAnchorPeersUpdate != "" {
		if err := c.doOutputAnchorPeersUpdate(profileConfig, c.channelID, c.outputAnchorPeersUpdate, c.asOrg); err != nil {
			logger.Error("Error on inspectChannelCreateTx: %s", err)
			err = errors.New("Error on inspectChannelCreateTx: " + err.Error())

			return err
		}
	}

	if c.printOrg != "" {
		logger.Info("Catch printOrg&outorg:", c.printOrg, "--", c.outorg)
		logger.Info(len(topLevelConfig.Organizations))
		if err := c.doPrintOrgTest(topLevelConfig, c.printOrg, c.outorg, isAddOrg); err != nil {
			logger.Error("Error on printOrg: %s", err)
			err = errors.New("Error on printOrg: " + err.Error())
			return err
		}
	}
	return nil
}

func (c *Configtxgen) printVersion() {
	fmt.Println(metadata.GetVersionInfo())
}
