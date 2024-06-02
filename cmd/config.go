package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinicius507/memos-cli/ui/styles"
)

var cfg Config

type Config struct {
	Api struct {
		AccessToken string `mapstructure:"token"`
		Url         string `mapstructure:"url"`
	} `mapstructure:"api"`
}

func initConfig() {
	xdgConfigHome, err := os.UserConfigDir()
	if err != nil {
		cobra.CheckErr(err)
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath(xdgConfigHome)
	viper.SetConfigName(".memos-cli")
	viper.SetEnvPrefix("memos")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		os.Stderr.WriteString(styles.ErrorMsg.Render(err.Error()) + "\n")
		os.Exit(1)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		os.Stderr.WriteString(styles.ErrorMsg.Render(err.Error()) + "\n")
		os.Exit(1)
	}
	if err := validateConfig(cfg); err != nil {
		os.Stderr.WriteString(styles.ErrorMsg.Render(err.Error()) + "\n")
		os.Exit(1)
	}
}

func validateConfig(config Config) error {
	if config.Api.Url == "" {
		return fmt.Errorf("api url is required")
	}
	if config.Api.AccessToken == "" {
		return fmt.Errorf("api token is required")
	}
	return nil
}
