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

// CreateDirIfNotExist : create directory if it doesn't exist
func CreateDirIfNotExist(path string) error {
	dirExists, err := Exists(path)
	if err != nil {
		return err
	}
	if dirExists {
		return nil
	}
	if err = Create(path, 0755); err != nil {
		return err
	}
	return nil
}

// CreateFileIfNotExist : create file if it doesn't exist
func CreateFileIfNotExist(path string) error {
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
