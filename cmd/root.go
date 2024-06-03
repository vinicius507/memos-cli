package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinicius507/memos-cli/cmd/create"
	"github.com/vinicius507/memos-cli/cmd/list"
	"github.com/vinicius507/memos-cli/config"
	"github.com/vinicius507/memos-cli/ui/styles"
)

func Execute() {
	cmd := &cobra.Command{
		Use:   "memos",
		Short: "A CLI client for Memos",
	}
	cmd.AddCommand(create.New(), list.New())
	if err := cmd.Execute(); err != nil {
		msg := styles.ErrorMsg.Render(err.Error())
		cmd.Println(msg)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cfg := config.GetConfig()
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
	if err := cfg.Validate(); err != nil {
		os.Stderr.WriteString(styles.ErrorMsg.Render(err.Error()) + "\n")
		os.Exit(1)
	}
	config.SetConfig(cfg)
}
