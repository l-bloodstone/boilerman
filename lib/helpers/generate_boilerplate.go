package helpers

import (
	"boilerman/lib/parser"
	"boilerman/lib/utils"
	"errors"

	"github.com/spf13/cobra"
)


func GenerateBoilerplate(cmd *cobra.Command, args []string) error {

	input, err := cmd.Flags().GetString("input")
	if err != nil {
		return(err)
	}
	if input == "" {
		return errors.New("Input was not specified")
	}
	outfile, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	if outfile == "" {
		return errors.New("Output file was not specified")
	}
	isForced, err := cmd.Flags().GetBool("force-ignore")
	
	templateFullPath := utils.GetFullFilePath(input)
	dataMap := parser.ParseKeyValueToMap(args)
	ignored, err := parser.ParseTemplate(templateFullPath, outfile, dataMap, isForced)
	if err != nil {
		cmd.Println("\nAvailable fields are...")
		for k, _ := range ignored {
			cmd.PrintErrln("`" + k + "`")
		}
		return err
	}
	// print all the ignored keys with warning
	if len(ignored) > 0 {
		cmd.Println("\nYou have ignored field(s).")
		for k, _ := range ignored {
			cmd.PrintErrln("  `" + k + "`")
		}
	}
	
	return nil
}
