package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppKey string
}

func GetAppConfig() AppConfig {
	csViper := viper.Sub("app")
	csViper.SetDefault("app_key", "fake")

	return AppConfig{
		AppKey: csViper.GetString("app_key"),
	}
}
