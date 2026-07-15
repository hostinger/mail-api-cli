package account

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "account",
	Short: "Account commands",
}

func init() {
	GroupCmd.AddCommand(CurrentCmd)
}
