package dir

import (
	"os"
)

// Create : create a directory
func Create(dirPath string, permission os.FileMode) error {
	err := os.Mkdir(dirPath, permission)
	if err != nil {
		return err
	}

	return nil
}

// SubDirs : get a list of all subdirectories
func SubDirs(dirPath string) ([]string, error) {
	subDirs := []string{}

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return subDirs, err
	}

	for _, f := range files {
		subDirs = append(subDirs, f.Name())
	}

	return subDirs, err
}

// IsDir : check if given path is dir or not
func IsDir(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

// Exists : check if directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateDirIfNotExist(path string) (error) {
	wExists, err := Exists(path)
	if err != nil {
		return err
	}
	if wExists {
		return nil
	}
	err = Create(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CreateFileIfNotExist(path string) (error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}
	return nil
}
