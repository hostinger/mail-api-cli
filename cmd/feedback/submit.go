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
	Long:  "Report a problem or suggestion about this API or the MCP server to the Hostinger mail team.\n\nReport when a call returned 4xx/5xx or unexpected data, was too slow, when documentation was missing or unclear, or when a capability you needed does not exist. Mention the failing operation and the status code you received so the team can find the request. Never include tokens, passwords or email contents: the message is scrubbed of secrets and capped at 2000 characters. Send one report per distinct issue.\n\nA `429` (`ERR_FEEDBACK_RATE_LIMIT`) means feedback for this customer was submitted less than ten seconds ago; wait and retry.",
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
	SubmitCmd.Flags().StringP("message", "", "", "What happened and what was expected, including the operation and status code involved. Never include tokens, passwords or email contents.")
	SubmitCmd.Flags().IntP("score", "", 0, "How well the API served the task: 1 (poor) to 10 (excellent).")
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
