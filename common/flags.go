package common

import (
	"github.com/spf13/cobra"
)

// common flags
const (
	FlagAppChainID = "app-chain-id"
	FlagHubChainID = "hub-chain-id"
	FlagAppNodeRPC = "app-rpc"
	FlagHubNodeRPC = "hub-rpc"
)

// AddCommonFlags appends common flags for given commands
func AddCommonFlags(cmds ...*cobra.Command) {
	for _, c := range cmds {
		c.Flags().String(FlagAppChainID, "fabric", "chain ID of the application chain")
		c.Flags().String(FlagAppNodeRPC, "grpc://localhost:7051", "RPC address of the application chain node")
		c.Flags().String(FlagHubChainID, "irita-hub", "chain ID of the Hub chain")
		c.Flags().String(FlagHubNodeRPC, "http://localhost:26657", "RPC address of the Hub chain node")
	}
}
