package messages

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var PatchCmd = &cobra.Command{
	Use:   "patch <mailbox-resource-id> <folder> <uid>",
	Short: "Update message flags",
	Long:  "Add and/or remove flags on a single message. Returns the updated message.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(patchBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().PatchMessageWithBodyWithResponse(context.TODO(), args[0], args[1], utils.StringToInt(args[2]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	PatchCmd.Flags().StringSliceP("addflags", "", nil, "IMAP flags to set on the message, e.g. \\Seen, \\Flagged, \\Answered, $forwarded.")
	PatchCmd.Flags().StringSliceP("removeflags", "", nil, "IMAP flags to clear from the message.")
}

func patchBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("addflags") {
		v, _ := cmd.Flags().GetStringSlice("addflags")
		body["addFlags"] = v
	}
	if cmd.Flags().Changed("removeflags") {
		v, _ := cmd.Flags().GetStringSlice("removeflags")
		body["removeFlags"] = v
	}
	return body
}
