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

var listFieldsCmd = &cobra.Command{
	Use: "field",
	Short: "List all the available fields from a boilerplate",
	RunE: helpers.ListBoilerplateFields,
}

func init() {
	rootCmd.AddCommand(listCmd)
	
	listCmd.AddCommand(listGroupCmd)

	listCmd.AddCommand(listFieldsCmd)
}
