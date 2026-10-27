package cmd

import (
	"context"
	"fmt"
	"github.com/hostinger/mail-api-cli/cmd/account"
	"github.com/hostinger/mail-api-cli/cmd/feedback"
	"github.com/hostinger/mail-api-cli/cmd/folders"
	"github.com/hostinger/mail-api-cli/cmd/messages"
	"github.com/hostinger/mail-api-cli/cmd/quota"
	"github.com/hostinger/mail-api-cli/cmd/send"
	"github.com/hostinger/mail-api-cli/cmd/webhooks"
	"os"
	"os/signal"

	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Command groups of the root help.
const (
	groupAPI    = "api"
	groupDeploy = "deploy"
	groupSetup  = "setup"
)

var RootCmd = &cobra.Command{
	Use:   "hostinger-mail",
	Short: "Hostinger Mail API Command Line Interface",
	Long:  "Hostinger Mail API Command Line Interface.\n\nGet started:\n  1. Sign in:        hostinger-mail config set api_token   (or set HOSTINGER_MAIL_API_TOKEN)\n\nOutput is a table on a terminal and JSON when piped; pick one with --format\n(table|json|yaml|tree) or filter it with --jq. Missing arguments and required\nflags are prompted for in a terminal (disable with --no-input). Commands that\ndelete or stop things ask for confirmation (skip with --yes), and --dry-run\nprints the request as curl without sending it.\n\nDocumentation: https://github.com/hostinger/mail-api-cli/tree/main/docs",
	// Errors are printed once, by Execute, with the exit code they map to.
	SilenceErrors:     true,
	SilenceUsage:      true,
	PersistentPreRunE: preRun,
}

// preRun rejects an invalid output format before an API request is sent, then
// prompts for missing arguments and required flags in a terminal.
func preRun(cmd *cobra.Command, args []string) error {
	if isAPICommand(cmd) {
		if _, err := output.SelectedFormat(cmd); err != nil {
			return err
		}
	}
	return utils.PromptMissing(cmd, args)
}

// isAPICommand reports whether cmd sits under one of the API command groups
// (setup commands such as config must keep working with a bad format setting).
func isAPICommand(cmd *cobra.Command) bool {
	for c := cmd; c.HasParent(); c = c.Parent() {
		if !c.Parent().HasParent() {
			return c.GroupID == groupAPI
		}
	}
	return false
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// The first Ctrl-C cancels the command; a second one kills the process.
	context.AfterFunc(ctx, stop)
	cmd, err := RootCmd.ExecuteContextC(ctx)
	stop()
	if cmd == nil {
		cmd = RootCmd
	}
	if code := utils.HandleError(os.Stderr, cmd, err); code != utils.ExitOK {
		os.Exit(code)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.DisableAutoGenTag = true
	RootCmd.SetUsageTemplate(utils.UsageTemplate)
	RootCmd.AddGroup(
		&cobra.Group{ID: groupAPI, Title: "API Commands:"},
		&cobra.Group{ID: groupSetup, Title: "Setup Commands:"},
	)
	RootCmd.SetHelpCommandGroupID(groupSetup)
	RootCmd.SetCompletionCommandGroupID(groupSetup)

	flags := RootCmd.PersistentFlags()
	flags.StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.hostinger-mail.yaml)")
	flags.StringP("format", "o", "", "Output format: table|json|yaml|tree (default: table on a terminal, json when piped)")
	flags.String("jq", "", "Filter the JSON response with a jq expression, e.g. '.data[].id'")
	flags.Bool("dry-run", false, "Print the request as a curl command instead of sending it")
	flags.Bool("verbose", false, "Print the config file used and each HTTP request to stderr")
	flags.Bool("no-input", false, "Never prompt; fail when an argument, flag or confirmation is missing")
	for key, flag := range map[string]string{"dry_run": "dry-run", "verbose": "verbose", "no_input": "no-input"} {
		cobra.CheckErr(viper.BindPFlag(key, flags.Lookup(flag)))
	}

	RootCmd.RegisterFlagCompletionFunc("format", cobra.FixedCompletions([]cobra.Completion{"table", "json", "yaml", "tree"}, cobra.ShellCompDirectiveNoFileComp))
	RootCmd.RegisterFlagCompletionFunc("jq", cobra.NoFileCompletions)

	account.GroupCmd.GroupID = groupAPI
	RootCmd.AddCommand(account.GroupCmd)
	feedback.GroupCmd.GroupID = groupAPI
	RootCmd.AddCommand(feedback.GroupCmd)
	folders.GroupCmd.GroupID = groupAPI
	RootCmd.AddCommand(folders.GroupCmd)
	messages.GroupCmd.GroupID = groupAPI
	RootCmd.AddCommand(messages.GroupCmd)
	quota.GroupCmd.GroupID = groupAPI
	RootCmd.AddCommand(quota.GroupCmd)
	send.GroupCmd.GroupID = groupAPI
	RootCmd.AddCommand(send.GroupCmd)
	webhooks.GroupCmd.GroupID = groupAPI
	RootCmd.AddCommand(webhooks.GroupCmd)
	VersionCmd.GroupID = groupSetup
	RootCmd.AddCommand(AuthCmd, ConfigCmd, VersionCmd)
	registerCompletionInstall(RootCmd)
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
	if err := viper.ReadInConfig(); err == nil && viper.GetBool("verbose") {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
