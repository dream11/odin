package artifact

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type version struct {
	Maven    string    `yaml:"maven"`
	Java     string    `yaml:"java"`
}

type flavour struct {
	Name       string     `yaml:"name"`
	Version    version    `yaml:"version"`
}

type steps struct {
	Pre      []string    `yaml:"pre"`
	Build    []string    `yaml:"build"`
	Run      []string    `yaml:"run"`
}

type Artifact struct {
	Flavour         flavour    `yaml:"flavour"`
	PropertyFile    string     `yaml:"propertyFile"`
	Port            string     `yaml:"port"`
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