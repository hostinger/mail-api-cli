package feedback

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Feedback commands",
}

func init() {
	GroupCmd.AddCommand(SubmitCmd)
}
