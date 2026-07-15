package send

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "send",
	Short: "Send commands",
}

func init() {
	GroupCmd.AddCommand(EmailCmd)
}
