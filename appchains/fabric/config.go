package fabric

import (
	"github.com/spf13/viper"

	"relayer/common"
)

const (
	Prefix         = "fabric"
	ChannelID      = "channel_id"
	PeerRPCAddrs   = "peer_rpc_addrs"
	OrdererRPCAddr = "orderer_rpc_addr"
	Key            = "key"
	Passphrase     = "passphrase"
)

// Config represents the Fabric chain config
type Config struct {
	ChannelID      string   `yaml:"channel_id"`
	PeerRPCAddrs   []string `yaml:"peer_rpc_addrs"`
	OrdererRPCAddr string   `yaml:"orderer_rpc_addr"`
	Key            string   `yaml:"key"`
	Passphrase     string   `yaml:"passphrase"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChannelID:      v.GetString(common.GetConfigKeyName(Prefix, ChannelID)),
		PeerRPCAddrs:   v.GetStringSlice(common.GetConfigKeyName(Prefix, PeerRPCAddrs)),
		OrdererRPCAddr: v.GetString(common.GetConfigKeyName(Prefix, OrdererRPCAddr)),
		Key:            v.GetString(common.GetConfigKeyName(Prefix, Key)),
		Passphrase:     v.GetString(common.GetConfigKeyName(Prefix, Passphrase)),
	}
}
