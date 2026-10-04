package helpers

import (
	"boilerman/lib/utils"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)


func CreateBoilerFromFile(cmd *cobra.Command, args []string) error {
	editor, err := cmd.Flags().GetString("editor")
	if err != nil {
		panic(err)
	}
	
	inputFileStr, err := cmd.Flags().GetString("input-file")
	if err != nil {
		panic(err)
	}

	outputFileStr := utils.GetFullFilePath(args[0])

	utils.CopyFileFromString(inputFileStr, outputFileStr)

	command := exec.Command(editor, outputFileStr)
	command.Dir = utils.GetBoilerDataDir()
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		panic(err)
	}

	return nil
}
