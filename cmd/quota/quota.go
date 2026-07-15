package quota

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "quota",
	Short: "Quota commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
}
