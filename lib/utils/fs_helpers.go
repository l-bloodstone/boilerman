package utils

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)


var boilerDataDir string
func GetBoilerDataDir() string {
	if boilerDataDir != "" {
		return boilerDataDir
	}
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	boilerDataDir = path.Join(userConfigDir, "boilerman")
	return  boilerDataDir
}

func CreateBoilerDirsByRelativePath(relativeFileName []string) {
	for _, fileName := range relativeFileName {
		fullFillPath := GetFullFilePath(fileName)
		fmt.Println("File Created: " + fullFillPath)
		if err := os.MkdirAll(filepath.Dir(fullFillPath), 0666); err != nil {
			panic(err)
		}
	}
}

func CleanBoilerDirs() error {
	// TODO: not implemented!
	return nil
}

func GetFullFilePath(relativeFileName string) string {
	return path.Join(GetBoilerDataDir(), relativeFileName + ".boil")
}

func GetFullFilePathSlice(relativeFileName []string) []string {
	var s []string
	for _, filename := range relativeFileName {
		s = append(s, GetFullFilePath(filename))
	}
	return s
}

func CopyFileFromString(sourceFullPath string, destinationFullPath string) {
	
	srcFile, err := os.Open(sourceFullPath)
	if err != nil {
		fmt.Println("Source file cannot be copied!")
		panic(err)
	}
	
	// creating all the parent directories
	err = os.MkdirAll(filepath.Dir(destinationFullPath), 0777)
	if err != nil {
		panic(err)
	}
	
	destFile, err := os.Create(destinationFullPath)
	if err != nil {
		fmt.Println("Couldn't create boilerplate file!")
		panic(err)
	}

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}
}
