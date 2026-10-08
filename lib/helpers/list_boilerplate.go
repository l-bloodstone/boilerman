package helpers

import (
	"boilerman/lib/parser"
	"boilerman/lib/utils"
	"errors"
	"os"
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

func ListBoilerplateFields(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Needs at least 1 argument")
	}
	for _, boilerplate := range args {
		text, err := os.ReadFile(path.Join(utils.GetBoilerDataDir(), boilerplate) + ".boil")
		if err != nil {
			return err
		}
		
		fields, err := parser.ParseFieldsFromFile(string(text))
		if err != nil {
			return err
		}
		cmd.Println("\n " + boilerplate + ":")
		for key, _ := range fields {
			cmd.Println("    `" + key + "`")
		}
	}
	return nil
}

func ListGroupsOfBoilerplates(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		for _, group := range args {
			err := utils.ReadSingleDirAndPrint(path.Join(utils.GetBoilerDataDir(), group))
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
