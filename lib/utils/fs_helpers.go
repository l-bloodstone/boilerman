package utils

import (
	"fmt"
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

func CreateBoilerDirs(relativeFileName []string) {
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
