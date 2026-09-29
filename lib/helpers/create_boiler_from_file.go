package helpers

import (
	"github.com/spf13/cobra"
)


func CreateBoilerFromFile(cmd *cobra.Command, args []string) error {
	// TODO: needs implementation
	editor, err := cmd.Flags().GetString("editor")
	if err != nil {
		panic(err)
	}
	cmd.Println(editor)
	return nil
}
