package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinicius507/memoscli/cmd/memo"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memos",
		Short: "A CLI client for Memos",
	}
	cmd.AddCommand(memo.NewMemoCmd())
	return cmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetEnvPrefix("memos")
	err := viper.BindEnv("api_url")
	cobra.CheckErr(err)
	err = viper.BindEnv("api_token")
	cobra.CheckErr(err)
}
