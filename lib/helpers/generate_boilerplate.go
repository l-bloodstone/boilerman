package helpers

import (
	"boilerman/lib/parser"
	"boilerman/lib/utils"
	"os"

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

	isForced, err := cmd.Flags().GetBool("force-ignore")

	templateFullPath := utils.GetFullFilePath(input)
	dataMap := parser.ParseKeyValueToMap(args)

	var srcFile *os.File
	if input == "" {
		srcFile = os.Stdin
	} else {
		srcFile, err = os.Open(templateFullPath)
		if err != nil {
			return err
		}
	}
	defer srcFile.Close()

	var destFile *os.File
	if outfile == "" {
		destFile = os.Stdout
	} else {
		destFile, err = os.Create(outfile)
		if err != nil {
			return err
		}
	}
	defer destFile.Close()
	
	ignored, err := parser.ParseTemplate(srcFile, destFile, dataMap, isForced)
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
