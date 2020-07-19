package main

import (
	"github.com/spf13/cobra"

	"relayer/appchains"
	"relayer/common"
	"relayer/core"
	"relayer/hub"
)

// StartCmd implements the start command
func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [config-file-path]",
		Short: "Start the relayer process with a config file path",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := ""

			if len(args) == 0 {
				configPath = "./config/"
			} else {
				configPath = args[0]
			}

			config, err := common.LoadYAMLConfig(configPath, "config")
			if err != nil {
				return err
			}

			appChain, err := appchains.NewAppChainFactory(config).Make(config.GetString("app_chain_name"))
			if err != nil {
				return err
			}

			hubChain := hub.MakeIritaHubChain(hub.NewConfig(config))

			relayerInstance := core.NewRelayer(hubChain, appChain, common.Log)
			relayerInstance.Start()

			return nil
		},
	}

	return cmd
}
