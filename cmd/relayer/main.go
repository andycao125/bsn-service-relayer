package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	fabriccmd "relayer/appchains/fabric/cmd"
	"relayer/common"
)

// rootCmd is the entry point
var (
	rootCmd = &cobra.Command{
		Use:   "relayer",
		Short: "Relayer command line interface",
	}
)

func main() {
	cobra.EnableCommandSorting = false

	// Add appchain-specific subcommands
	appchainCmds := []*cobra.Command{fabriccmd.FabricCmd()}
	rootCmd.AddCommand(appchainCmds...)

	// Add common flags
	common.AddCommonFlags(appchainCmds...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
