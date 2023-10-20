package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	PasswordEncryptionKey string
	TokenEncryptionKey    string
	PaginationSize        int
}

func GetAppConfig() AppConfig {
	csViper := viper.Sub("app")
	csViper.SetDefault("password_encryption_key", "password_encrypt")
	csViper.SetDefault("token_encryption_key", "token_encryption")
	csViper.SetDefault("pagination_size", 10)

	return AppConfig{
		PasswordEncryptionKey: csViper.GetString("password_encryption_key"),
		TokenEncryptionKey:    csViper.GetString("token_encryption_key"),
		PaginationSize:        csViper.GetInt("pagination_size"),
	}
}
