package cmd

import (
	"github.com/spf13/cobra"
	"boilerman/lib/helpers"
)


var removeGroupCmd = &cobra.Command{
	Use: "group",
	Short: "Remove a group of boilerplates",
	RunE: helpers.RemoveGroupFunc,
}


func init() {
	removeCmd.AddCommand(removeGroupCmd)
}
