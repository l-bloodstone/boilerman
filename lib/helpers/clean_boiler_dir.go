package helpers

import (
	"github.com/spf13/cobra"
	"boilerman/lib/utils"
)


func CleanEmptyBoilerDirs(cmd *cobra.Command, args []string) error {

	err := utils.CleanBoilerDirs(utils.GetBoilerDataDir(), "")
	if err != nil {
		return err
	}

	return nil
}
