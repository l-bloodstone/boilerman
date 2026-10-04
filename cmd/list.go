package cmd

import (
	"github.com/spf13/cobra"
	"boilerman/lib/helpers"
)


var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all the available boilerplates",
	RunE: helpers.ListBoilerplates,
}

var listGroupCmd = &cobra.Command{
	Use: "group",
	Short: "List only the gorups of boilerplates",
	RunE: helpers.ListGroupsOfBoilerplates,
}

func init() {
	rootCmd.AddCommand(listCmd)
	
	listCmd.AddCommand(listGroupCmd)
}
