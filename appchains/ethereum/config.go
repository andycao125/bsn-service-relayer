package ethereum

import (
	"github.com/spf13/viper"

	"relayer/common"
)

const (
	Prefix                 = "eth"
	ChainID                = "chain_id"
	NodeRPCAddr            = "node_rpc_addr"
	IServiceCoreContract   = "iservice_core_contract"
	IServiceMarketContract = "iservice_market_contract"
	Key                    = "key"
	Passphrase             = "passphrase"
)

// Config represents the Ethereum chain config
type Config struct {
	ChainID                string `yaml:"chain_id"`
	NodeRPCAddr            string `yaml:"node_rpc_addr"`
	IServiceCoreContract   string `yaml:"iservice_core_contract"`
	IServiceMarketContract string `yaml:"iservice_market_contract"`
	Key                    string `yaml:"key"`
	Passphrase             string `yaml:"passphrase"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChainID:                v.GetString(common.GetConfigKeyName(Prefix, ChainID)),
		NodeRPCAddr:            v.GetString(common.GetConfigKeyName(Prefix, NodeRPCAddr)),
		IServiceCoreContract:   v.GetString(common.GetConfigKeyName(Prefix, IServiceCoreContract)),
		IServiceMarketContract: v.GetString(common.GetConfigKeyName(Prefix, IServiceMarketContract)),
		Key:                    v.GetString(common.GetConfigKeyName(Prefix, Key)),
		Passphrase:             v.GetString(common.GetConfigKeyName(Prefix, Passphrase)),
	}
}
