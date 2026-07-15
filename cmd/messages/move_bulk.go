package messages

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/spf13/cobra"
)

var MoveBulkCmd = &cobra.Command{
	Use:   "move-bulk <mailbox-resource-id> <folder>",
	Short: "Move messages",
	Long:  "Move multiple messages from a source folder to a target folder.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(moveBulkBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MoveMessagesWithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	MoveBulkCmd.Flags().StringP("targetfolder", "", "", "Destination folder path.")
	MoveBulkCmd.Flags().IntSliceP("uids", "", nil, "Message UIDs to move. 1-100 entries, each > 0.")
	MoveBulkCmd.MarkFlagRequired("targetfolder")
	MoveBulkCmd.MarkFlagRequired("uids")
}

func moveBulkBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	targetFolderVal, _ := cmd.Flags().GetString("targetfolder")
	body["targetFolder"] = targetFolderVal
	uidsVal, _ := cmd.Flags().GetIntSlice("uids")
	body["uids"] = uidsVal
	return body
}
