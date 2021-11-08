package service

import (
	"fmt"
)

type Component struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}

type Service struct {
	Name       string      `yaml:"name" json:"name"`
	Version    string      `yaml:"version" json:"version"`
	Components []Component `yaml:"components" json:"components"`
}

type Services []Service

func (s *Services) GetService(name string) (Service, error) {
	for _, service := range *s {
		if service.Name == name {
			return service, nil
		}
	}

	return Service{}, fmt.Errorf("%s does not exists", name)
}
