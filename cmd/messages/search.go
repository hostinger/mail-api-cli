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
	SearchCmd.Flags().IntP("page", "", 1, "Page number (1-based).")
	SearchCmd.Flags().IntP("perpage", "", 25, "Items per page (max 100).")
	SearchCmd.Flags().StringP("sort", "", "-uid", "Sort field with optional `-` prefix for descending. Allowed: uid, date, size.")
	SearchCmd.Flags().StringP("before", "", "", "Only messages received before this date (YYYY-MM-DD).")
	SearchCmd.Flags().StringP("body", "", "", "Case-insensitive substring match on the message body only (headers excluded). OR-combined with subject/from/to/cc.")
	SearchCmd.Flags().StringP("cc", "", "", "Case-insensitive substring match on the Cc header. OR-combined with subject/from/to/body.")
	SearchCmd.Flags().StringSliceP("flags", "", nil, "Only messages carrying all of these IMAP flags, e.g. \\Seen, \\Flagged, \\Answered, $forwarded.")
	SearchCmd.Flags().StringP("from", "", "", "Case-insensitive substring match on the From header. OR-combined with subject/to/cc/body.")
	SearchCmd.Flags().StringP("header", "", "", "Match a specific header as Name:value, e.g. X-Custom-Header:value. Value match is a substring.")
	SearchCmd.Flags().IntP("larger", "", 0, "Only messages larger than this size in bytes.")
	SearchCmd.Flags().StringP("since", "", "", "Only messages received on or after this date (YYYY-MM-DD).")
	SearchCmd.Flags().IntP("smaller", "", 0, "Only messages smaller than this size in bytes.")
	SearchCmd.Flags().StringP("subject", "", "", "Case-insensitive substring match on the Subject header. OR-combined with from/to/cc/body.")
	SearchCmd.Flags().StringP("text", "", "", "Case-insensitive substring match across headers and body.")
	SearchCmd.Flags().StringP("to", "", "", "Case-insensitive substring match on the To header. OR-combined with subject/from/cc/body.")
	SearchCmd.Flags().StringP("uid", "", "", "IMAP UID set: single UID, range (1:100), open range (100:*), or comma-separated list.")
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
