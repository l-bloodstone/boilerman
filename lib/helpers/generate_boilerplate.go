package helpers

import (
	"boilerman/lib/parser"
	"boilerman/lib/utils"

	"github.com/spf13/cobra"
)


func GenerateBoilerplate(cmd *cobra.Command, args []string) error {

	input, err := cmd.Flags().GetString("input")
	if err != nil {
		return(err)
	}
	outfile, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	templateFullPath := utils.GetFullFilePath(input)
	dataMap := parser.ParseKeyValueToMap(args)
	ignored, err := parser.ParseTemplate(templateFullPath, outfile, dataMap)
	if err != nil {
		return err
	}
	// print all the ignored keys with warning
	cmd.Println("You have ignored this/these field(s).")
	for k, _ := range ignored {
		cmd.PrintErrln(k)
	}
	
	return nil
}
