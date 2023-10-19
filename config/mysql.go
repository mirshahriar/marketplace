package config

import (
	"time"

	"github.com/spf13/viper"
)

// DBConfig holds the database credential & other configuration
// nolint:tagalign
type DBConfig struct {
	Host        string        `json:"host,omitempty" yml:"host"`
	Port        int           `json:"port,omitempty" yml:"port"`
	Name        string        `json:"name,omitempty" yml:"name"`
	User        string        `json:"user,omitempty" yml:"user"`
	Password    string        `json:"password,omitempty" yml:"password"`
	MaxIdleConn int           `json:"max_idle_conn,omitempty" yml:"max_idle_conn"`
	MaxOpenConn int           `json:"max_open_conn,omitempty" yml:"max_open_conn"`
	LogDB       bool          `json:"log_db,omitempty" yml:"log_db"`
	MaxConnTime time.Duration `json:"max_conn_time" yml:"max_conn_time"`
}

// GetDBConfig returns database configuration
func GetDBConfig() DBConfig {
	msViper := viper.Sub("mysql")
	msViper.SetDefault("port", 3306)
	msViper.SetDefault("max_idle_conn", 10)
	msViper.SetDefault("max_open_conn", 10)
	msViper.SetDefault("log_db", true)
	msViper.SetDefault("max_conn_time", 5*time.Second)

	return DBConfig{
		Host:        msViper.GetString("host"),
		Port:        msViper.GetInt("port"),
		Password:    msViper.GetString("password"),
		User:        msViper.GetString("user"),
		Name:        msViper.GetString("name"),
		MaxIdleConn: msViper.GetInt("max_idle_conn"),
		MaxOpenConn: msViper.GetInt("max_open_conn"),
		MaxConnTime: msViper.GetDuration("max_conn_time"),
		LogDB:       msViper.GetBool("log_db"),
	}
}
