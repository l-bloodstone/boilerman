package cmd

import (
	"boilerman/lib/helpers"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use: "gen",
	Short: "Generate a boilerplate from a base to current directory.",
	RunE: helpers.GenerateBoilerplate,
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringP("input", "i", "", "specify a template to generate")
	genCmd.Flags().StringP("output", "o", "", "output file name")
}
