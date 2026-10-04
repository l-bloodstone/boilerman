package helpers

import (
	"boilerman/lib/utils"
	"path"

	"github.com/spf13/cobra"
)


func ListBoilerplates(cmd *cobra.Command, args []string) error {
	err := utils.ReadDirAndPrintFilesRecursive(utils.GetBoilerDataDir(), "")
	if err != nil {
		return err
	}
	return nil
}

func ListGroupsOfBoilerplates(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		for _, group := range args {
			err := utils.ReadSingleDirAndPrint(path.Join(utils.GetBoilerDataDir() + "/" + group))
			if err != nil {
				return err
			}
		}
		return nil
	}
	
	err := utils.ReadDirAndPrintRecur(utils.GetBoilerDataDir(), "")
	if err != nil {
		return err
	}
	return nil
}
