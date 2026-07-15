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

var DeleteBulkCmd = &cobra.Command{
	Use:   "delete-bulk <mailbox-resource-id> <folder>",
	Short: "Delete messages",
	Long:  "Permanently delete multiple messages from a folder.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deleteBulkBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DeleteMessagesWithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteBulkCmd.Flags().IntSliceP("uids", "", nil, "Message UIDs to delete. 1-100 entries, each > 0.")
	DeleteBulkCmd.MarkFlagRequired("uids")
}

func deleteBulkBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	uidsVal, _ := cmd.Flags().GetIntSlice("uids")
	body["uids"] = uidsVal
	return body
}
