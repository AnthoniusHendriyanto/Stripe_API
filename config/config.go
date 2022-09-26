package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database  Database        `mapstructure:",squash"`
	StripeAPI StripeAPIConfig `mapstructure:",squash"`
}

type Database struct {
	Host   string
	User   string
	DBName string
}

type StripeAPIConfig struct {
	Address string
	Token   string
}

var cfg *Config

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: %v", err)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		fmt.Printf("Error on unmarshalling config: %v", err)
	}

	return cfg, nil
}
