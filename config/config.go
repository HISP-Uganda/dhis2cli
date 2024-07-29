package config

import (
	"fmt"
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

type GlobalParamsConfig struct {
	Paging   string
	Fields   string
	Order    string
	Query    string
	Filter   string
	Page     int
	PageSize int
}

var GlobalParams GlobalParamsConfig

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

// GenerateParams generates a map of parameters from the global config,
// applies defaultParams, merges in any additionalParams, and excludes any keys in excludeKeys.
func GenerateParams(config GlobalParamsConfig, defaultParams, additionalParams map[string]string, excludeKeys []string) map[string]string {
	params := make(map[string]string)

	// Start with defaultParams
	for key, value := range defaultParams {
		params[key] = value
	}

	// Overwrite with values from global config
	if config.Paging != "" {
		params["paging"] = config.Paging
	}
	if config.Fields != "" {
		params["fields"] = config.Fields
	}
	if config.Order != "" {
		params["order"] = config.Order
	}
	if config.Query != "" {
		params["query"] = config.Query
	}
	if config.Filter != "" {
		params["filter"] = config.Filter
	}
	if config.Paging == "true" {
		params["page"] = fmt.Sprintf("%d", config.Page)
		params["pageSize"] = fmt.Sprintf("%d", config.PageSize)
	}

	// Merge in additionalParams, which should also overwrite any existing values
	for key, value := range additionalParams {
		params[key] = value
	}

	// Remove excluded keys
	for _, key := range excludeKeys {
		delete(params, key)
	}

	return params
}
