package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	PrivateKey      string
	WalletAddress  string
	GasLimit uint64
	Networks []Network
}

type Network struct {
	Name string
	Url string
	Type string
	Contracts map[string]string
}

func Load() (Config, error) {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return c, nil
}