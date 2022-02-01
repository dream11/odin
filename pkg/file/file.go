package file

import (
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
