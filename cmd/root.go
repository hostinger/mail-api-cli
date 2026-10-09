package cmd

import (
	"fmt"
	"github.com/hostinger/mail-api-cli/cmd/account"
	"github.com/hostinger/mail-api-cli/cmd/feedback"
	"github.com/hostinger/mail-api-cli/cmd/folders"
	"github.com/hostinger/mail-api-cli/cmd/messages"
	"github.com/hostinger/mail-api-cli/cmd/quota"
	"github.com/hostinger/mail-api-cli/cmd/send"
	"github.com/hostinger/mail-api-cli/cmd/webhooks"
	"os"

	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	OutputFormat string
	validFormats = []string{"json", "table", "tree"}
)

var RootCmd = &cobra.Command{
	Use:   "hostinger-mail",
	Short: "Hostinger Mail API Command Line Interface",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.DisableAutoGenTag = true

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.hostinger-mail.yaml)")
	RootCmd.PersistentFlags().StringVar(&OutputFormat, "format", "", "Output format type (json|table|tree), default: table")

	RootCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return validFormats, cobra.ShellCompDirectiveNoFileComp
	})

	RootCmd.AddCommand(account.GroupCmd)
	RootCmd.AddCommand(feedback.GroupCmd)
	RootCmd.AddCommand(folders.GroupCmd)
	RootCmd.AddCommand(messages.GroupCmd)
	RootCmd.AddCommand(quota.GroupCmd)
	RootCmd.AddCommand(send.GroupCmd)
	RootCmd.AddCommand(webhooks.GroupCmd)
	RootCmd.AddCommand(VersionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		path, _ := utils.MigrateLegacyConfig(home, os.Stderr)
		viper.SetConfigFile(path)
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("hostinger_mail")
	viper.AutomaticEnv() // read in environment variables that match
	utils.BindLegacyEnv(os.Stderr)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
