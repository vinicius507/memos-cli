package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinicius507/memos-cli/cmd/create"
)

var cfgFile string

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memos",
		Short: "A CLI client for Memos",
	}
	cmd.AddCommand(create.New())
	return cmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		xdgConfigHome, err := os.UserConfigDir()
		if err != nil {
			cobra.CheckErr(err)
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath(xdgConfigHome)
		viper.SetConfigName(".memos-cli")
	}
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}
