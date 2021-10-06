package artifact

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type flavour struct {
	Name       string               `yaml:"name"`
	Version    map[string]string    `yaml:"version"`
}

type steps struct {
	Pre      []string    `yaml:"pre"`
	Build    []string    `yaml:"build"`
	Post     []string    `yaml:"post"`
}

type Artifact struct {
	Flavour         flavour    `yaml:"flavour"`
	PropertyFile    string     `yaml:"propertyFile"`
	Steps           steps      `yaml:"steps"`
	ArtifactPath    string     `yaml:"artifactPath"`
}

func ParseFile(filePath string) (Artifact, error) {
	var artifact Artifact

	yFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return artifact, err
	}

	err = yaml.Unmarshal(yFile, &artifact)

	return artifact, err
}