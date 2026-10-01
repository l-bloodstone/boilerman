package helpers

import (
	"boilerman/lib/utils"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RemoveGroupFunc(cmd *cobra.Command, args []string) error {
	for _, dir := range args {
		fmt.Print("Do you want to delete this boilerplate group? (y/n): ")
		buf := make([]byte, 2)
		os.Stdin.Read(buf)
		switch buf[0] {
		case 'y':
			err := utils.RemoveBoilerplateGroup(dir)
			if err != nil {
				return err
			}
		case 'n':
			return errors.New("Deletion Aborted!")

		default:
			return errors.New("only y/n supported")
		}
	}
	return nil
}
