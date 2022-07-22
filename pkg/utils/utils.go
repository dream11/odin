package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/config"
	"github.com/dream11/odin/pkg/file"
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

func SetEnv(envName string) error {
	configPath := path.Join(app.WorkDir.Location, app.WorkDir.ConfigFile)
	data, err := file.Read(configPath)
	if err != nil {
		return err
	}
	result := ""
	r, _ := regexp.Compile(`(?:envName: [a-zA-Z]+-\w+)`)
	match := r.FindString(string(data))
	if match != "" {
		result = strings.Replace(string(data), match, fmt.Sprintf("envName: %s", envName), 1)
	} else {
		result = string(data) + fmt.Sprintf("envName: %s\n", envName)
	}
	err = file.Write(configPath, result, 0755)
	if err != nil {
		return err
	}
	return err
}

func FetchEnv(envName string) string {
	if envName != "" {
		return envName
	} else {
		var appConfig = config.Get()
		return appConfig.EnvName
	}
}
