package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"relayer/appchains/fabric"
	"relayer/common"
	"relayer/core"
	"relayer/hub"
)

// Fabric-specific flags
const (
	FlagChannelID      = "channel-id"
	FlagOrdererNodeRPC = "orderer-rpc"
)

var (
	FsFabric = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsFabric.String(FlagChannelID, "", "Fabric channel ID")
	FsFabric.String(FlagOrdererNodeRPC, "grpc://localhost:7052", "RPC address of the Fabric orderer node")
}

// FabricCmd implements the Fabric subcommand
func FabricCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fabric",
		Short: "Start the Fabric relayer process",
		RunE: func(cmd *cobra.Command, args []string) error {
			appChainID := viper.GetString(common.FlagAppChainID)
			appNodeRPC := viper.GetString(common.FlagAppNodeRPC)
			channelID := viper.GetString(FlagChannelID)
			ordererNodeRPC := viper.GetString(FlagOrdererNodeRPC)

			hubChainID := viper.GetString(common.FlagHubChainID)
			hubNodeRPC := viper.GetString(common.FlagHubNodeRPC)

			appChain := fabric.NewFabricChain(appChainID, channelID, []string{appNodeRPC}, ordererNodeRPC, nil, nil)
			hubChain := hub.NewIritaHubChain(hubChainID, hubNodeRPC, nil, nil)

			logger := log.New(os.Stdout, "", log.LstdFlags)

			relayerInstance := core.NewRelayer(hubChain, appChain, logger)
			relayerInstance.Start()

			return nil
		},
	}

	cmd.Flags().AddFlagSet(FsFabric)

	return cmd
}
