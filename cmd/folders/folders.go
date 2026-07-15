package folders

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "folders",
	Short: "Folders commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
