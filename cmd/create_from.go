package cmd


import (
	"github.com/spf13/cobra"

	"boilerman/lib/helpers"
)


var createFromCmd = &cobra.Command{
	Use: "from",
	Short: "Create from a file as a base of the boilerplate. Simple copy-paste in backbround for convenience.",
	RunE: helpers.CreateBoilerFromFile,
}


func init() {
	createCmd.AddCommand(createFromCmd)
	createFromCmd.Flags().StringP("input-file", "i", "", "input file for boilerplate as a base")
}
