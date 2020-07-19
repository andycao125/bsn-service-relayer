package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// MustGetHomeDir gets the user home directory
// Panic if an error occurs
func MustGetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return homeDir
}

// LoadYAMLConfig loads the YAML config file
func LoadYAMLConfig(configPath, configFileName string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(configPath)
	v.SetConfigFile(configFileName)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read the config file: %s", err)
	}

	return v, nil
}

// GetConfigKeyName returns the key name with the given prefix
func GetConfigKeyName(prefix string, key string) string {
	return fmt.Sprintf("%s.%s", prefix, key)
}
