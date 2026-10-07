package cmd

import (
	"github.com/spf13/cobra"
	"boilerman/lib/helpers"
)


var viewCmd = &cobra.Command{
	Use: "view",
	Short: "print out template to stdout",
	RunE: helpers.ViewTemplatesToStdout,
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
