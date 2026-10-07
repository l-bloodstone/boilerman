package helpers

import (
	"boilerman/lib/utils"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)


func EditBoilerplates(cmd *cobra.Command, args []string) error {
	editor, err := cmd.Flags().GetString("editor")
	if err != nil {
		return err
	}
	boilerplateFilePaths := []string{}
	for _, v := range args {
		boilerplateFilePaths = append(boilerplateFilePaths, utils.GetFullFilePath(v))
	}
	command := exec.Command(editor, boilerplateFilePaths...)
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		return err
	}
	return nil
}
