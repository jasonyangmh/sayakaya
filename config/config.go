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

	Schedule string `mapstructure:"SCHEDULE"`
	CodeLen  int    `mapstructure:"CODE_LEN"`

	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPort  string `mapstructure:"SMTP_Port"`
	SMTPEmail string `mapstructure:"SMTP_EMAIL"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
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
