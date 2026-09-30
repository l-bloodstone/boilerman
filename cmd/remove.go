package cmd

import (
	"boilerman/lib/helpers"

	"github.com/spf13/cobra"
)


var removeCmd = &cobra.Command{
	Use: "remove",
	Short: "Remove a boilerplate from stored directory",
	RunE: helpers.RemoveFiles,
}


func init() {
	rootCmd.AddCommand(removeCmd)
}
