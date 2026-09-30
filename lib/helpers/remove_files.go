package helpers

import (
	"boilerman/lib/utils"

	"github.com/spf13/cobra"
)


func RemoveFiles(cmd *cobra.Command, args []string) error {
	err := utils.RemoveBoilerplate(args)
	if err != nil {
		cmd.PrintErrln("File Doesn't exists or You do not have the right permission.")
		return err
	}
	return nil
}
