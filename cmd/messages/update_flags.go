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

var UpdateFlagsCmd = &cobra.Command{
	Use:   "update-flags <mailbox-resource-id> <folder>",
	Short: "Update message flags",
	Long:  "Add and/or remove flags on multiple messages. Returns 200 when all UIDs succeed, 207 with per-UID outcome when some fail.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateFlagsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().UpdateMessageFlagsWithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateFlagsCmd.Flags().StringSliceP("addflags", "", nil, "IMAP flags to set on every listed message, e.g. \\Seen, \\Flagged, \\Answered, $forwarded.")
	UpdateFlagsCmd.Flags().StringSliceP("removeflags", "", nil, "IMAP flags to clear from every listed message.")
	UpdateFlagsCmd.Flags().IntSliceP("uids", "", nil, "Message UIDs to update. 1-100 entries, each > 0.")
	UpdateFlagsCmd.MarkFlagRequired("uids")
}

func updateFlagsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("addflags") {
		v, _ := cmd.Flags().GetStringSlice("addflags")
		body["addFlags"] = v
	}
	if cmd.Flags().Changed("removeflags") {
		v, _ := cmd.Flags().GetStringSlice("removeflags")
		body["removeFlags"] = v
	}
	uidsVal, _ := cmd.Flags().GetIntSlice("uids")
	body["uids"] = uidsVal
	return body
}
