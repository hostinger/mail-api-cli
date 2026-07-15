package messages

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "messages",
	Short: "Messages commands",
}

func init() {
	GroupCmd.AddCommand(AttachmentCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(DeleteAllCmd)
	GroupCmd.AddCommand(DeleteBulkCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(MoveCmd)
	GroupCmd.AddCommand(MoveBulkCmd)
	GroupCmd.AddCommand(PatchCmd)
	GroupCmd.AddCommand(SearchCmd)
	GroupCmd.AddCommand(SourceCmd)
	GroupCmd.AddCommand(TextCmd)
	GroupCmd.AddCommand(UpdateFlagsCmd)
}
