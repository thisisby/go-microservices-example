package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	DSN  string `mapstructure:"DSN"`
}

func LoadConfig() Config {
	var config Config

	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}

	return config
}
