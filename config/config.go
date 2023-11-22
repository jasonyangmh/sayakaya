package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASSWORD"`
	DBName string `mapstructure:"DB_NAME"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
}

func Load() *Config {
	cfg := &Config{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to find the config file: %v", err)
		return nil
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("unable to load the environment: %v", err)
		return nil
	}

	return cfg
}
