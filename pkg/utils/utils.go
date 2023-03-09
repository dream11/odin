package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/dream11/odin/api/service"
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
	match := SearchString(string(data), `(?:envName: \S+)`)
	if match != "nil" {
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

func SearchString(stringMeta string, stringToSearch string) string {
	r, _ := regexp.Compile(stringToSearch)
	match := r.FindString(stringMeta)
	if match != "" {
		return match
	}
	return "nil"
}

func FetchKey(keyName string) string {
	var appConfig = config.Get()
	r := reflect.ValueOf(appConfig)
	f := reflect.Indirect(r).FieldByName(keyName)
	return f.String()
}

func maxWidth(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetColumnWidth(services []service.Service) []int {
	a := []int{4, 7, 15}
	for _, service := range services {
		a[0] = maxWidth(a[0], len(service.Name))
		a[1] = maxWidth(a[1], len(service.Version))
	}
	return a
}
