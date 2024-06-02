package config

import (
	"fmt"
)

var cfg Config

type Config struct {
	Api struct {
		AccessToken string `mapstructure:"token"`
		Url         string `mapstructure:"url"`
	} `mapstructure:"api"`
}

func GetConfig() Config {
	return cfg
}

func SetConfig(c Config) {
	cfg = c
}

func (c *Config) Validate() error {
	if c.Api.Url == "" {
		return fmt.Errorf("invalid configuration: api url is required")
	}
	if c.Api.AccessToken == "" {
		return fmt.Errorf("invalid configuration: api token is required")
	}
	return nil
}
