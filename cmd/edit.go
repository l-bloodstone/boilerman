package cmd

import (
	"boilerman/lib/helpers"

	"github.com/spf13/cobra"
	"os"
)


var editCmd = &cobra.Command{
	Use: "edit",
	Short: "Open and edit a file using a preferred editor.",
	RunE: helpers.EditBoilerplates,
}


func init() {
	rootCmd.AddCommand(editCmd)
	editorEnv := os.Getenv("EDITOR")
	if editorEnv == "" {
		createCmd.Println("An editor was not specified. Either use `-e` or set `EDITOR` environment variable, or edit the file manually.")
	}
	editCmd.Flags().StringP("editor", "e", editorEnv, "set preferred editor")
}
