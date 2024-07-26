package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		DatabaseURI string `mapstructure:"database_uri"`
		BaseUrl     string `mapstructure:"base_url"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		AuthToken   string `mapstructure:"auth_token"`
		AuthMethod  string `mapstructure:"auth_method"`
	} `yaml:"server"`
}

var Paging string
var Fields string
var Filter string
var Query string
var Order string
var Page int
var PageSize int
var OutputFormat string
var OutputFile string
var Verbose bool

var Cfg Config

func LoadConfig(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)
	}
}
