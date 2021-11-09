package component

import (
	"fmt"
)

type Artifact struct {
	Url     string `yaml:"url" json:"url"`
	Version string `yaml:"version" json:"version"`
	Type    string `yaml:"type" json:"type"`
}

type Component struct {
	Name     string   `yaml:"name" json:"name"`
	Version  string   `yaml:"version" json:"version"`
	Type     string   `yaml:"type" json:"type"`
	Artifact Artifact `yaml:"artifact" json:"artifact"`
}

type Components []Component

func (c *Components) GetComponent(name string) (Component, error) {
	for _, component := range *c {
		if component.Name == name {
			return component, nil
		}
	}

	return Component{}, fmt.Errorf("%s does not exists", name)
}
