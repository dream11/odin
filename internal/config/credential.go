package config

import (
	"path"

	"gopkg.in/yaml.v3"

	"github.com/dream11/odin/api/configuration"
	"github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/file"
)

var logger ui.Logger
var secretText = "<secret text>"

// Get : fetch credentials from all sources
func Get() configuration.Configuration {
	configPath := path.Join(app.WorkDir.Location, app.WorkDir.ConfigFile)
	var configs configuration.Configuration

	rawConfig, _ := file.Read(configPath)
	err := yaml.Unmarshal(rawConfig, &configs)
	if err != nil {
		logger.Error("Unable to parse configuration. " + err.Error())
	}

	if len(configs.AccessToken) == 0 {
		logger.Warn("unable to load Access Token.")
		logger.Debug("Access Token not found at: " + app.WorkDir.ConfigFile)
	}

	// do not expose access key, secret access key & refresh tokens
	configs.Keys = configuration.SecretKeys{
		AccessKey:       secretText,
		SecretAccessKey: secretText,
	}
	configs.RefreshToken = secretText

	return configs
}
