package main

import (
	"fmt"

	"github.com/spf13/cobra"

	cfg "relayer/config"
	"relayer/hub"
)

var (
	HubCmd = &cobra.Command{
		Use:   "hub",
		Short: "Irita-Hub commands",
	}

	KeysCmd = &cobra.Command{
		Use:   "keys",
		Short: "Key management commands",
	}
)

// KeysAddCmd implements the keys add command
func KeysAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name] [passphrase] [config-file-path]",
		Short: "Generate a new key",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 2 {
				configFileName = cfg.DefaultConfigFileName
			} else {
				configFileName = args[2]
			}

			config, err := cfg.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			hubChain := hub.MakeIritaHubChain(hub.NewConfig(config))

			addr, mnemonic, err := hubChain.AddKey(args[0], args[1])
			if err != nil {
				return err
			}

			fmt.Printf("key generated successfully: \n\nname: %s\naddress: %s\nmnemonic: %s\n\n", args[0], addr, mnemonic)

			return nil
		},
	}

	return cmd
}

// KeysShowCmd implements the keys show command
func KeysShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [name] [passphrase] [config-file-path]",
		Short: "Show the key infomation by name",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 2 {
				configFileName = cfg.DefaultConfigFileName
			} else {
				configFileName = args[2]
			}

			config, err := cfg.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			hubChain := hub.MakeIritaHubChain(hub.NewConfig(config))

			addr, err := hubChain.ShowKey(args[0], args[1])
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", addr)

			return nil
		},
	}

	return cmd
}

// KeysImportCmd implements the keys import command
func KeysImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import [name] [passphrase] [key-armor] [config-file-path]",
		Short: "Import a key from the private key armor",
		Args:  cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 3 {
				configFileName = cfg.DefaultConfigFileName
			} else {
				configFileName = args[2]
			}

			config, err := cfg.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			hubChain := hub.MakeIritaHubChain(hub.NewConfig(config))

			addr, err := hubChain.ImportKey(args[0], args[1], args[2])
			if err != nil {
				return err
			}

			fmt.Printf("key imported successfully: %s\n", addr)

			return nil
		},
	}

	return cmd
}

func init() {
	KeysCmd.AddCommand(
		KeysAddCmd(),
		KeysShowCmd(),
		KeysImportCmd(),
	)

	HubCmd.AddCommand(KeysCmd)
}
