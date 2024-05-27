package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PrivateKey              string `mapstructure:"PRIVATE_KEY"`
	ResearchContractAddress string `mapstructure:"RESEARCH_CONTRACT_ADDRESS"`
	GoogleAPIKEY            string `mapstructure:"GOOGLE_API_KEY"`
	NetworkName             string `mapstructure:"NETWORK"`
	NetCfg                  *NetworkConfig
}

func InitConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../shared/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return &Config{}, fmt.Errorf("failed to read config file, %v", err)
	}

	cfg.NetCfg = GetNetworkConfig(cfg.NetworkName)

	return &cfg, nil
}
