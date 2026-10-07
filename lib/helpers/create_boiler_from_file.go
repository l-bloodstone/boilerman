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

	if len(inputFileStr) > 2 {
		err := utils.CopyFileFromString(inputFileStr, outputFileStr)
		if err != nil {
			return err
		}
		
	} else {
		err := utils.CreateFileFromStdin(os.Stdin, args[0])
		if err != nil {
			return err
		}
	}

	command := exec.Command(editor, outputFileStr)
	command.Dir = utils.GetBoilerDataDir()
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		cmd.PrintErr("File created but couldn't find specified editor. use -e flag or manually edit using `boilerman edit -e`")
	}

	return nil
}
