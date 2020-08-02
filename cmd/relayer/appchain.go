package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"relayer/appchains"
	cmn "relayer/common"
)

var (
	AppChainCmd = &cobra.Command{
		Use:   "appchain",
		Short: "appchain commands",
	}
)

// AddServiceBindingCmd implements the appchain add-svc-binding command
func AddServiceBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-svc-binding [svc-name] [schemas] [provider] [svc-fee] [qos] [config-file-path]",
		Short: "Add the specified service binding",
		Args:  cobra.RangeArgs(5, 6),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 5 {
				configFileName = cmn.DefaultConfigFileName
			} else {
				configFileName = args[5]
			}

			qos, err := strconv.ParseUint(args[4], 10, 64)
			if err != nil {
				return err
			}

			config, err := cmn.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			appChain, err := appchains.NewAppChainFactory(config).Make(config.GetString(cmn.ConfigKeyAppChainName))
			if err != nil {
				return err
			}

			err = appChain.AddServiceBinding(args[0], args[1], args[2], args[3], qos)
			if err != nil {
				return err
			}

			fmt.Printf("AddServiceBinding transaction minted\n")

			return nil
		},
	}

	return cmd
}

// UpdateServiceBindingCmd implements the appchain update-svc-binding command
func UpdateServiceBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-svc-binding [svc-name] [provider] [svc-fee] [qos] [config-file-path]",
		Short: "Update the specified service binding",
		Args:  cobra.RangeArgs(4, 5),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 4 {
				configFileName = cmn.DefaultConfigFileName
			} else {
				configFileName = args[4]
			}

			qos, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			config, err := cmn.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			appChain, err := appchains.NewAppChainFactory(config).Make(config.GetString(cmn.ConfigKeyAppChainName))
			if err != nil {
				return err
			}

			err = appChain.UpdateServiceBinding(args[0], args[1], args[2], qos)
			if err != nil {
				return err
			}

			fmt.Printf("UpdateServiceBinding transaction minted\n")

			return nil
		},
	}

	return cmd
}

func init() {
	AppChainCmd.AddCommand(
		AddServiceBindingCmd(),
		UpdateServiceBindingCmd(),
	)
}
