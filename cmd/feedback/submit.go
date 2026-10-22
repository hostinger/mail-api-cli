package feedback

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var SubmitCmd = &cobra.Command{
	Use:               "submit <mailbox-resource-id>",
	Short:             "Submit feedback",
	Long:              "Send the user's feedback about the Hostinger Email API to the Hostinger mail team.\n\nOnly call this when the user explicitly asks to send feedback, report a problem, or request a feature. Send the user's own words; never include tokens, passwords or email contents. The message is capped at 2000 characters.\n\nA `429` (`ERR_FEEDBACK_RATE_LIMIT`) means feedback for this customer was submitted less than ten seconds ago; wait and retry.",
	Example:           "  hostinger-mail feedback submit my-mailbox-resource-id --message 'Sending a message with a 10 MB attachment fails with a timeout.' --score 3",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := submitBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().SubmitFeedbackWithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(SubmitCmd, utils.StringArg("mailbox-resource-id"))
	SubmitCmd.Flags().StringP("message", "", "", "The user's feedback in their own words. Never include tokens, passwords or email contents.")
	SubmitCmd.Flags().IntP("score", "", 0, "The user's rating of the Hostinger Email API: 1 (poor) to 10 (excellent).")
	SubmitCmd.MarkFlagRequired("message")
	SubmitCmd.MarkFlagRequired("score")
	SubmitCmd.RegisterFlagCompletionFunc("message", cobra.NoFileCompletions)
	SubmitCmd.RegisterFlagCompletionFunc("score", cobra.NoFileCompletions)
}

func submitBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	messageVal, _ := cmd.Flags().GetString("message")
	body["message"] = messageVal
	scoreVal, _ := cmd.Flags().GetInt("score")
	body["score"] = scoreVal
	return body, nil
}
