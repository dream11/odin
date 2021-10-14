package dir

import (
    "os"
	"io/ioutil"
)

func SubDirs(dirPath string) ([]string, error) {
	subDirs := []string{}

	files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        return subDirs, err
    }
 
    for _, f := range files {
        subDirs = append(subDirs, f.Name())
    }

	return subDirs, err
}

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