package util

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
	var parsedContent interface{}

	if len(filePath) == 0 {
		return parsedContent, errors.New("filepath cannot be empty")
	}

	fileContent, err := Read(filePath)
	if err != nil {
		return parsedContent, errors.New("Error reading file: " + err.Error())
	}

	if strings.Contains(filePath, ".yaml") || strings.Contains(filePath, ".yml") {
		err = yamlPorvider.Unmarshal(fileContent, &parsedContent)
		if err != nil {
			return parsedContent, errors.New("Unable to parse YAML file: " + err.Error())
		}
	} else if strings.Contains(filePath, ".json") {
		err = json.Unmarshal(fileContent, &parsedContent)
		if err != nil {
			return parsedContent, errors.New("Unable to parse JSON file: " + err.Error())
		}
	} else {
		return parsedContent, errors.New("unrecognized file format")
	}
	return parsedContent, nil
}
