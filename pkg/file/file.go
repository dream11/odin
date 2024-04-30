package file

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"strings"

	yamlPorvider "gopkg.in/yaml.v3"
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

// FindAndReadAllAllowedFormat : takes in file path and allowed format list and returns data path and error
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

// ParseFile parse json or yaml file given path to json and return as interface
func ParseFile(filePath string) (interface{}, error) {
	if len(filePath) != 0 {
		var parsedDefinition interface{}

		fileDefinition, err := Read(filePath)
		if err != nil {
			return parsedDefinition, errors.New("file does not exist")
		}

		if strings.Contains(filePath, ".yaml") || strings.Contains(filePath, ".yml") {
			err = yamlPorvider.Unmarshal(fileDefinition, &parsedDefinition)
			if err != nil {
				return parsedDefinition, errors.New("Unable to parse YAML. " + err.Error())
			}
		} else if strings.Contains(filePath, ".json") {
			err = json.Unmarshal(fileDefinition, &parsedDefinition)
			if err != nil {
				return parsedDefinition, errors.New("Unable to parse JSON. " + err.Error())
			}
		} else {
			return parsedDefinition, errors.New("unrecognized file format")
		}
		return parsedDefinition, nil
	}
	return errors.New("filepath cannot be empty"), nil
}
