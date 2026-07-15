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

var MoveCmd = &cobra.Command{
	Use:   "move <mailbox-resource-id> <folder> <uid>",
	Short: "Move message",
	Long:  "Move a single message to a target folder.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(moveBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MoveMessageWithBodyWithResponse(context.TODO(), args[0], args[1], utils.StringToInt(args[2]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	MoveCmd.Flags().StringP("targetfolder", "", "", "Destination folder path.")
	MoveCmd.MarkFlagRequired("targetfolder")
}

func moveBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	targetFolderVal, _ := cmd.Flags().GetString("targetfolder")
	body["targetFolder"] = targetFolderVal
	return body
}
