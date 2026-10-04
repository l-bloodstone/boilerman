package helpers

import (
	"github.com/spf13/cobra"
	"errors"
	"os"
	"os/exec"

	"boilerman/lib/utils"

)


func CreateBoilerFile(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Specify boilerplate name(s)")
	}
	editor, _:= cmd.Flags().GetString("editor")

	// creates all the parent directories of the arguments.
	err := utils.CreateBoilerDirsByRelativePath(args)
	if err != nil {
		return err
	}

	if editor == "" {
		editors := []string{"nvim", "kate"}
		for _, editor := range editors {
			editCommand := exec.Command(editor, utils.GetFullFilePathSlice(args)...)
			editCommand.Dir = utils.GetBoilerDataDir()
			editCommand.Stdout = os.Stdout
			if err := editCommand.Run(); err != nil {
				continue
			}
		}
	} else {
		editCommand := exec.Command(editor, utils.GetFullFilePathSlice(args)...)
		editCommand.Dir = utils.GetBoilerDataDir()
		editCommand.Stdout = os.Stdout
		if err := editCommand.Run(); err != nil {
			return err
		}
	}
	return nil
}
