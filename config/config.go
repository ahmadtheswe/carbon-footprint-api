package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBStringConnection string `mapstructure:"DB_STRING_CONNECTION"`
	Port               string `mapstructure:"PORT"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file:", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatal("Error unmarshaling config:", err)
	}
}
