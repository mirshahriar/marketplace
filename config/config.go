// Package config is used to load configuration from consul.
package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Init initiates of config load
func Init() {
	// Setting the prefix for environment variables.
	viper.SetEnvPrefix("app")
	_ = viper.BindEnv("env")
	_ = viper.BindEnv("consul_url")
	_ = viper.BindEnv("consul_path")

	// Getting the value of consul_url from the config file.
	consulURL := viper.GetString("consul_url")
	// Getting the value of consul_path from the config file.
	consulPath := viper.GetString("consul_path")

	if consulURL == "" || consulPath == "" {
		viper.SetConfigType("json")
		configReader := strings.NewReader(defaultConfig)

		if err := viper.ReadConfig(configReader); err != nil {
			log.Fatal(err.Error())
		}
	} else {
		// Adding a remote provider to viper.
		_ = viper.AddRemoteProvider("consul", consulURL, consulPath)
		// Setting the type of the config file.
		viper.SetConfigType("yml")
		if err := viper.ReadRemoteConfig(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

var defaultConfig = `
{
   "app":{
      "app_key":"fake",
   },
   "mysql":{
      "name":"product",
      "host":"localhost",
      "port":"3306",
      "user":"root",
      "password":"test"
   }
}
`
