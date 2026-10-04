package cmd

import (

	"boilerman/lib/helpers"

	"github.com/spf13/cobra"
	"os"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new boilerplate",
	RunE: helpers.CreateBoilerFile,
}

func init() {
	rootCmd.AddCommand(createCmd)

	editorEnv := os.Getenv("EDITOR")
	if editorEnv == "" {
		createCmd.Println("An editor was not specified. Either use `-e` or set `EDITOR` environment variable.")
	}
	createCmd.PersistentFlags().StringP("editor", "e", editorEnv, "specify an editor")
	createCmd.PersistentFlags().BoolP("no-editor", "n", false, "set it true if editor should not open")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
