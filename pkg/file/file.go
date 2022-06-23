package file

import (
	"errors"
	"io/fs"
	"os"
)

// Read : read data from file
func Read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// Write : write data to file
func Write(filePath, data string, permission fs.FileMode) error {
	byteData := []byte(data)
	err := os.WriteFile(filePath, byteData, permission)

	return err
}

func FindAndReadAllAllowedFormat(path string, allowedFormats []string) ([]byte, string, error) {
	for _, allowedFormat := range allowedFormats {
		filepath := path + allowedFormat
		data, err := Read(filepath)
		if err == nil {
			return data, filepath, nil
		}
	}
	return nil, "", errors.New("unable to read file: " + path)
}
