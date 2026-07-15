package messages

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search <mailbox-resource-id> <folder>",
	Short: "Search messages",
	Long:  "Search messages in a folder. Filters in body; pagination and sort via query (`page`, `perPage`, `sort`).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(searchBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().SearchMessagesWithBodyWithResponse(context.TODO(), args[0], args[1], searchParams(cmd), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SearchCmd.Flags().IntP("page", "", 1, "")
	SearchCmd.Flags().IntP("perpage", "", 25, "")
	SearchCmd.Flags().StringP("sort", "", "-uid", "")
	SearchCmd.Flags().StringP("before", "", "", "")
	SearchCmd.Flags().StringP("body", "", "", "")
	SearchCmd.Flags().StringP("cc", "", "", "")
	SearchCmd.Flags().StringSliceP("flags", "", nil, "")
	SearchCmd.Flags().StringP("from", "", "", "")
	SearchCmd.Flags().StringP("header", "", "", "")
	SearchCmd.Flags().IntP("larger", "", 0, "")
	SearchCmd.Flags().StringP("since", "", "", "")
	SearchCmd.Flags().IntP("smaller", "", 0, "")
	SearchCmd.Flags().StringP("subject", "", "", "")
	SearchCmd.Flags().StringP("text", "", "", "")
	SearchCmd.Flags().StringP("to", "", "", "")
	SearchCmd.Flags().StringP("uid", "", "", "")
}

func searchParams(cmd *cobra.Command) *client.SearchMessagesParams {
	params := &client.SearchMessagesParams{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("perpage") {
		v, _ := cmd.Flags().GetInt("perpage")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("sort") {
		v, _ := cmd.Flags().GetString("sort")
		params.Sort = &v
	}
	return params
}

func searchBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("before") {
		v, _ := cmd.Flags().GetString("before")
		body["before"] = v
	}
	if cmd.Flags().Changed("body") {
		v, _ := cmd.Flags().GetString("body")
		body["body"] = v
	}
	if cmd.Flags().Changed("cc") {
		v, _ := cmd.Flags().GetString("cc")
		body["cc"] = v
	}
	if cmd.Flags().Changed("flags") {
		v, _ := cmd.Flags().GetStringSlice("flags")
		body["flags"] = v
	}
	if cmd.Flags().Changed("from") {
		v, _ := cmd.Flags().GetString("from")
		body["from"] = v
	}
	if cmd.Flags().Changed("header") {
		v, _ := cmd.Flags().GetString("header")
		body["header"] = v
	}
	if cmd.Flags().Changed("larger") {
		v, _ := cmd.Flags().GetInt("larger")
		body["larger"] = v
	}
	if cmd.Flags().Changed("since") {
		v, _ := cmd.Flags().GetString("since")
		body["since"] = v
	}
	if cmd.Flags().Changed("smaller") {
		v, _ := cmd.Flags().GetInt("smaller")
		body["smaller"] = v
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
		v, _ := cmd.Flags().GetString("to")
		body["to"] = v
	}
	if cmd.Flags().Changed("uid") {
		v, _ := cmd.Flags().GetString("uid")
		body["uid"] = v
	}
	return body
}
