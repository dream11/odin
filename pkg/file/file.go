package file

import (
	"io/fs"
	"os"
)

func Write(filePath, data string, permission fs.FileMode) error {
	byteData := []byte(data)
	err := os.WriteFile(filePath, byteData, permission)

	return err
}
