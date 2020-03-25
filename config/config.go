package config

import (
	"os"

	"github.com/spf13/viper"
)

var loaded bool

func Load() {
	if loaded {
		return
	}
	loaded = true
	if os.Getenv("ENVIRONMENT") != "test" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName("test")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	initLogConfig()
	initServerConfig()
	initDatabaseConfig()
}
