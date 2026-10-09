package feedback

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/spf13/cobra"
)

var SubmitCmd = &cobra.Command{
	Use:   "submit <mailbox-resource-id>",
	Short: "Submit feedback",
	Long:  "Send the user's feedback about the Hostinger Email API to the Hostinger mail team.\n\nOnly call this when the user explicitly asks to send feedback, report a problem, or request a feature. Send the user's own words; never include tokens, passwords or email contents. The message is capped at 2000 characters.\n\nA `429` (`ERR_FEEDBACK_RATE_LIMIT`) means feedback for this customer was submitted less than ten seconds ago; wait and retry.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(submitBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().SubmitFeedbackWithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SubmitCmd.Flags().StringP("message", "", "", "The user's feedback in their own words. Never include tokens, passwords or email contents.")
	SubmitCmd.Flags().IntP("score", "", 0, "The user's rating of the Hostinger Email API: 1 (poor) to 10 (excellent).")
	SubmitCmd.MarkFlagRequired("message")
	SubmitCmd.MarkFlagRequired("score")
}

func submitBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	messageVal, _ := cmd.Flags().GetString("message")
	body["message"] = messageVal
	scoreVal, _ := cmd.Flags().GetInt("score")
	body["score"] = scoreVal
	return body
}
