package utils

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
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

func CreateBoilerDirsByRelativePath(relativeFileName []string) error {
	for _, fileName := range relativeFileName {
		fullFillPath := GetFullFilePath(fileName)
		if err := os.MkdirAll(filepath.Dir(fullFillPath), 0777); err != nil {
			return err
		}
		fmt.Println("File Created: " + fullFillPath)
	}
	return nil
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
	defer srcFile.Close()
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
	defer destFile.Close()
	if err != nil {
		fmt.Println("Couldn't create boilerplate file!")
		panic(err)
	}

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}
}

func RemoveBoilerplate(relativePaths []string) error {
	for _, file := range relativePaths {
		fileFullPath := GetFullFilePath(file)
		err := os.Remove(fileFullPath)
		if err != nil {
			return err
		}
		fmt.Println("File Removed: " + fileFullPath)
	}
	return nil
}

func RemoveBoilerplateGroup(relativePath string) error {
	dirFullPath := path.Join(GetBoilerDataDir(), relativePath)
	_, err := os.ReadDir(dirFullPath)
	if err != nil {
		fmt.Println("This group is not exists or Not a directory at all.")
		return err
	}
	os.RemoveAll(dirFullPath)
	return nil

}

func ReadDirAndPrintRecur(prevPath string, newpath string) error {
	n := path.Join(prevPath, newpath)
	dirs, err := os.ReadDir(n)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			ReadDirAndPrintRecur(n, dir.Name())
			continue
		}
		reg := regexp.MustCompile(`boilerman\/(.*)$`)
		file := reg.FindStringSubmatch(path.Dir(path.Join(n, dir.Name())))[1]
		fmt.Println(file)
	}
	return nil
}

func ReadSingleDirAndPrint(dirname string) error {
	dir, err := os.Open(dirname)
	defer dir.Close()
	if err != nil {
		return err
	}
	files, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, file := range files {
		_, err := os.ReadDir(path.Join(dir.Name(), file))
		if err == nil {
			continue
		}
		reg := regexp.MustCompile(`boilerman\/(.*).boil$`)
		f := reg.FindStringSubmatch(path.Join(dir.Name(), file))[1]
		fmt.Println(f)
	}
	return nil
}

func ReadDirAndPrintFilesRecursive(prevPath string, newpath string) error {
	n := path.Join(prevPath, newpath)
	dirs, err := os.ReadDir(n)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			ReadDirAndPrintFilesRecursive(n, dir.Name())
			continue
		}
		reg := regexp.MustCompile(`boilerman\/(.*).boil$`)
		file := reg.FindStringSubmatch(path.Join(n, dir.Name()))[1]
		fmt.Println(file)
	}
	return nil
}
