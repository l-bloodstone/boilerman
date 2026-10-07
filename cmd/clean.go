package cmd

import (
	"github.com/spf13/cobra"
	"boilerman/lib/helpers"
)


var cleanCmd = &cobra.Command{
	Use: "clean",
	Short: "Clean empty boiler groups",
	RunE: helpers.CleanEmptyBoilerDirs,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
