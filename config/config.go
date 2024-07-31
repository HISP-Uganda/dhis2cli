package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
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

var OutputFormat string
var QueryParams []string
var OutputFile string
var TableMaxStringLength int
var Verbose bool

type GlobalParamsConfig struct {
	Paging   string
	Fields   string
	Order    string
	Query    string
	Filter   []string
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
func GenerateParams(config GlobalParamsConfig, defaultParams, additionalParams map[string]any, excludeKeys []string) map[string]any {
	params := make(map[string]any)

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
	if len(config.Filter) > 0 {
		params["filter"] = config.Filter
	}
	if config.Paging == "true" {
		params["page"] = fmt.Sprintf("%d", config.Page)
		params["pageSize"] = fmt.Sprintf("%d", config.PageSize)
	}
	if config.Paging == "false" {
		delete(params, "pageSize")
		delete(params, "page")
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

// ParamsMap takes a slice of "key=value" strings and returns a map[string]string
func ParamsMap(pairs []string) map[string]any {
	params := make(map[string]any)
	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			params[key] = value
		}
	}
	return params
}

// CombineMaps merges two maps of type map[string]interface{} into a new map.
// Values from the second map will overwrite values from the first map if they have the same key.
func CombineMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	combined := make(map[string]interface{})

	// Add all entries from map1 to the combined map
	for key, value := range map1 {
		combined[key] = value
	}

	// Add all entries from map2 to the combined map
	// This will overwrite any duplicate keys from map1
	for key, value := range map2 {
		combined[key] = value
	}

	return combined
}
