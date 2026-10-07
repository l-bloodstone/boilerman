package helpers

import (
	"boilerman/lib/utils"
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
)


func ViewTemplatesToStdout(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return errors.New("At least one argument is needed")
	}

	for _, arg := range args {
		fullpath := utils.GetFullFilePath(arg)
		f, err := os.Open(fullpath)
		if err != nil {
			return err
		}
		defer f.Close()

		os.Stdout.WriteString("\n------" + arg + ":------\n\n")
		n, err := f.WriteTo(os.Stdout)
		if err != nil && err != io.EOF {
			return err
		}
		if n < 1 {
			return errors.New("The file is empty")
		}
		os.Stdout.WriteString("\n------------\n")
	}

	return nil
}
