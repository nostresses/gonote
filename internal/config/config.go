package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DataSourceName string
}

// Getting the application configuration
func New() (*Config, error) {
	b, err := os.ReadFile("../../configs/conf.json")

	if err != nil {
		return nil, err
	}

	var conf *Config

	if err := json.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}
