package hub

import (
	iritasdk "github.com/bianjieai/irita-ent-sdk-go/types"
	"github.com/spf13/viper"

	"relayer/common"
)

// default config variables
var (
	defaultChainID     = "irita-hub"
	defaultNodeRPCAddr = "http://127.0.0.1:26657"
	defaultMode        = iritasdk.Commit
	defaultDBRoot      = common.MustGetHomeDir()
)

const (
	Prefix      = "hub"
	ChainID     = "chain_id"
	NodeRPCAddr = "node_rpc_addr"
	Key         = "key"
	Passphrase  = "passphrase"
	DBRoot      = "db_root"
)

// Config is a config struct for IRITA-HUB
type Config struct {
	ChainID     string `yaml:"chain_id"`
	NodeRPCAddr string `yaml:"node_rpc_addr"`
	Key         string `yaml:"key"`
	Passphrase  string `yaml:"passphrase"`
	DBRoot      string `yaml:"db_root"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChainID:     v.GetString(common.GetConfigKeyName(Prefix, ChainID)),
		NodeRPCAddr: v.GetString(common.GetConfigKeyName(Prefix, NodeRPCAddr)),
		Key:         v.GetString(common.GetConfigKeyName(Prefix, Key)),
		Passphrase:  v.GetString(common.GetConfigKeyName(Prefix, Passphrase)),
		DBRoot:      v.GetString(common.GetConfigKeyName(Prefix, DBRoot)),
	}
}
