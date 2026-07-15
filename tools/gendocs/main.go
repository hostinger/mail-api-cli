package main

import (
	"os"

	"github.com/hostinger/mail-api-cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	if err := doc.GenMarkdownTree(cmd.RootCmd, "./docs"); err != nil {
		os.Exit(1)
	}
}
