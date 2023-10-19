package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppKey         string
	PaginationSize int
}

func GetAppConfig() AppConfig {
	csViper := viper.Sub("app")
	csViper.SetDefault("app_key", "fake")
	csViper.SetDefault("pagination_size", 10)

	return AppConfig{
		AppKey:         csViper.GetString("app_key"),
		PaginationSize: csViper.GetInt("pagination_size"),
	}
}
