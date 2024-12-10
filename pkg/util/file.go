package util

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	yamlPorvider "gopkg.in/yaml.v3"
)

// Read : read data from file
func Read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
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
