package send

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

var EmailCmd = &cobra.Command{
	Use:   "email <mailbox-resource-id>",
	Short: "Send email",
	Long:  "Send a message from the managed mailbox. Saves a copy to INBOX.Sent.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(emailBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().SendEmailWithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	EmailCmd.Flags().StringP("attachments", "", "", " (JSON)")
	EmailCmd.Flags().StringSliceP("bcc", "", nil, "")
	EmailCmd.Flags().StringSliceP("cc", "", nil, "")
	EmailCmd.Flags().StringP("displayname", "", "", "")
	EmailCmd.Flags().StringP("forwardof", "", "", "Source message this forwards. Copies its Message-Id/References into In-Reply-To/References and flags it $forwarded. Mutually exclusive with inReplyTo. (JSON)")
	EmailCmd.Flags().StringP("html", "", "", "")
	EmailCmd.Flags().StringP("inreplyto", "", "", "Source message this is a reply to. Copies its Message-Id/References into In-Reply-To/References and flags it \\Answered. Mutually exclusive with forwardOf. (JSON)")
	EmailCmd.Flags().StringP("subject", "", "", "")
	EmailCmd.Flags().StringP("text", "", "", "")
	EmailCmd.Flags().StringSliceP("to", "", nil, "")
}

func emailBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("attachments") {
		v, _ := cmd.Flags().GetString("attachments")
		body["attachments"] = utils.JSONValue(v, "attachments")
	}
	if cmd.Flags().Changed("bcc") {
		v, _ := cmd.Flags().GetStringSlice("bcc")
		body["bcc"] = v
	}
	if cmd.Flags().Changed("cc") {
		v, _ := cmd.Flags().GetStringSlice("cc")
		body["cc"] = v
	}
	if cmd.Flags().Changed("displayname") {
		v, _ := cmd.Flags().GetString("displayname")
		body["displayName"] = v
	}
	if cmd.Flags().Changed("forwardof") {
		v, _ := cmd.Flags().GetString("forwardof")
		body["forwardOf"] = utils.JSONValue(v, "forwardof")
	}
	if cmd.Flags().Changed("html") {
		v, _ := cmd.Flags().GetString("html")
		body["html"] = v
	}
	if cmd.Flags().Changed("inreplyto") {
		v, _ := cmd.Flags().GetString("inreplyto")
		body["inReplyTo"] = utils.JSONValue(v, "inreplyto")
	}
	if cmd.Flags().Changed("subject") {
		v, _ := cmd.Flags().GetString("subject")
		body["subject"] = v
	}
	if cmd.Flags().Changed("text") {
		v, _ := cmd.Flags().GetString("text")
		body["text"] = v
	}
	if cmd.Flags().Changed("to") {
		v, _ := cmd.Flags().GetStringSlice("to")
		body["to"] = v
	}
	return body
}
