package send

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var EmailCmd = &cobra.Command{
	Use:               "email <mailbox-resource-id>",
	Short:             "Send email",
	Long:              "Send a message from the managed mailbox. Saves a copy to INBOX.Sent.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --attachments\n    [\n      {\n        \"cid\": \"logo\",\n        \"content\": \"JVBERi0xLjQK...\",\n        \"contentType\": \"application/pdf\",\n        \"encoding\": \"base64\",\n        \"filename\": \"invoice.pdf\"\n      }\n    ]\n\n  --forwardof\n    {\n      \"folder\": \"INBOX\",\n      \"uid\": 123\n    }\n\n  --inreplyto\n    {\n      \"folder\": \"INBOX\",\n      \"uid\": 123\n    }",
	Example:           "  hostinger-mail send email my-mailbox-resource-id",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := emailBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().SendEmailWithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(EmailCmd, utils.StringArg("mailbox-resource-id"))
	EmailCmd.Flags().StringP("attachments", "", "", "Files to attach. Inline images set cid; regular attachments omit it. (JSON, @file or -; shape in the help above)")
	EmailCmd.Flags().StringSliceP("bcc", "", nil, "Blind-carbon-copy recipient email addresses. Not visible to other recipients.")
	EmailCmd.Flags().StringSliceP("cc", "", nil, "Carbon-copy recipient email addresses.")
	EmailCmd.Flags().StringP("displayname", "", "", "Sender display name shown in the From header alongside the mailbox address.")
	EmailCmd.Flags().StringP("forwardof", "", "", "Source message this forwards. Copies its Message-Id/References into In-Reply-To/References and flags it $forwarded. Mutually exclusive with inReplyTo. (JSON, @file or -; shape in the help above)")
	EmailCmd.Flags().StringP("html", "", "", "HTML body. Optional; if both text and html are omitted the message is sent without a body. Inline images are referenced via cid: URLs matching attachment cid values.")
	EmailCmd.Flags().StringP("inreplyto", "", "", "Source message this is a reply to. Copies its Message-Id/References into In-Reply-To/References and flags it \\Answered. Mutually exclusive with forwardOf. (JSON, @file or -; shape in the help above)")
	EmailCmd.Flags().StringP("subject", "", "", "Message subject line.")
	EmailCmd.Flags().StringP("text", "", "", "Plain-text body. Optional; if both text and html are omitted the message is sent without a body.")
	EmailCmd.Flags().StringSliceP("to", "", nil, "Primary recipient email addresses.")
	EmailCmd.RegisterFlagCompletionFunc("attachments", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	EmailCmd.RegisterFlagCompletionFunc("bcc", cobra.NoFileCompletions)
	EmailCmd.RegisterFlagCompletionFunc("cc", cobra.NoFileCompletions)
	EmailCmd.RegisterFlagCompletionFunc("displayname", cobra.NoFileCompletions)
	EmailCmd.RegisterFlagCompletionFunc("forwardof", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	EmailCmd.RegisterFlagCompletionFunc("html", cobra.NoFileCompletions)
	EmailCmd.RegisterFlagCompletionFunc("inreplyto", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	EmailCmd.RegisterFlagCompletionFunc("subject", cobra.NoFileCompletions)
	EmailCmd.RegisterFlagCompletionFunc("text", cobra.NoFileCompletions)
	EmailCmd.RegisterFlagCompletionFunc("to", cobra.NoFileCompletions)
}

func emailBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("attachments") {
		attachmentsVal, _ := cmd.Flags().GetString("attachments")
		attachmentsValJSON, err := utils.JSONValue(attachmentsVal, "attachments")
		if err != nil {
			return nil, err
		}
		body["attachments"] = attachmentsValJSON
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
		forwardOfVal, _ := cmd.Flags().GetString("forwardof")
		forwardOfValJSON, err := utils.JSONValue(forwardOfVal, "forwardof")
		if err != nil {
			return nil, err
		}
		body["forwardOf"] = forwardOfValJSON
	}
	if cmd.Flags().Changed("html") {
		v, _ := cmd.Flags().GetString("html")
		body["html"] = v
	}
	if cmd.Flags().Changed("inreplyto") {
		inReplyToVal, _ := cmd.Flags().GetString("inreplyto")
		inReplyToValJSON, err := utils.JSONValue(inReplyToVal, "inreplyto")
		if err != nil {
			return nil, err
		}
		body["inReplyTo"] = inReplyToValJSON
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
	return body, nil
}
