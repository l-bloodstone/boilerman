package utils

import (
	"errors"
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

func CleanBoilerDirs(current string, prev string) error {
	baseDir := path.Join(current, prev)
	dirEntry, err := os.ReadDir(baseDir)
	if err != nil {
		return err
	}

	for _, dir := range dirEntry {
		if dir.IsDir() {
			fmt.Println(dir.Name())
			openDir, err := os.Open(path.Join(baseDir, dir.Name()))
			if err != nil {
				return err
			}
			defer openDir.Close()
			
			dirnames, err := openDir.Readdirnames(1)
			if err != nil && err != io.EOF {
				return err
			}
			if len(dirnames) == 0 {
				os.Remove(path.Join(baseDir, dir.Name()))
			} else {
				
				CleanBoilerDirs(baseDir, dir.Name())
				
				dirnames, err = openDir.Readdirnames(1)
				if err != nil && err != io.EOF {
					return err
				}
				if len(dirnames) == 0 {
					os.Remove(path.Join(baseDir, dir.Name()))
				}
			}
		}
	}

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

func CopyFileFromString(sourceFullPath string, destinationFullPath string) error {

	srcFile, err := os.Open(sourceFullPath)
	defer srcFile.Close()
	if err != nil {
		return err
	}

	// creating all the parent directories
	err = os.MkdirAll(filepath.Dir(destinationFullPath), 0777)
	if err != nil {
		return err
	}

	destFile, err := os.Create(destinationFullPath)
	defer destFile.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

func CreateFileFromStdin(stdin *os.File, dest string) error {

	destFullPath := GetFullFilePath(dest)
	f, err := os.Create(destFullPath)
	defer f.Close()
	if err != nil {
		return err
	}

	n, err := f.ReadFrom(stdin)
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("Zero bytes read from stdin")
	}

	return nil
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
	reg := regexp.MustCompile(`.*/boilerman/(.*)`)
	for _, dir := range dirs {
		if dir.IsDir() {
			dirFullPath := path.Join(n, dir.Name())
			group := reg.FindStringSubmatch(dirFullPath)[1]
			fmt.Println(group)
			ReadDirAndPrintRecur(n, dir.Name())
		}
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
