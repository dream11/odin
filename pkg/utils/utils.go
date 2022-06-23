package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

func GetProvisioningFileName(env string) string {
	return fmt.Sprintf("provisioning-%s", env)
}

func ParserYmlOrJson(filePath string, in []byte) (interface{}, error) {
	var out interface{}
	if strings.Contains(filePath, ".yaml") || strings.Contains(filePath, ".yml") {
		err := yaml.Unmarshal(in, &out)
		if err != nil {
			return nil, errors.New("unable to parse file: " + filePath + "\n" + err.Error())
		}
		return out, nil
	} else if strings.Contains(filePath, ".json") {
		err := json.Unmarshal(in, &out)
		if err != nil {
			return nil, errors.New("unable to parse file: " + filePath + "\n" + err.Error())
		}
		return out, nil
	}
	return nil, errors.New("unable to parse file. unknown file format found in file: " + filePath)
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
