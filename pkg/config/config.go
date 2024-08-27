package config

import (
	"errors"
	"strings"
	"sync"

	"github.com/dream11/odin/api/configuration"
	"github.com/dream11/odin/app"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var once sync.Once
var appConfig *configuration.Configuration
var err error

func readConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/." + app.App.Name)
	viper.SetEnvPrefix("ODIN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Fatal("Not configured odin yet? Run `odin configure`")
		}
		log.Fatal("Error while reading config file: ", err)
	}
}

func getConfigForProfile(profile string) (*configuration.Configuration, error) {
	config := configuration.Configuration{}
	if err := viper.UnmarshalKey(profile, &config); err != nil {
		log.Fatal("Configuration can't be loaded: ", err)
	}
	return &config, nil
}

func readConfig() (*configuration.Configuration, error) {
	readConfigFile()
	profile := viper.GetString("profile")
	return getConfigForProfile(profile)
}

// GetConfig returns the reference of viper config
func GetConfig() *configuration.Configuration {
	once.Do(func() {
		appConfig, err = readConfig()
	})
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}
	return appConfig
}

// WriteConfig writes the given config to the config file
func WriteConfig(config *configuration.Configuration) {
	profile := viper.GetString("profile")
	viper.Set(profile, config)
	if err := viper.WriteConfig(); err != nil {
		log.Fatal("Unable to write configuration: ", err)
	}
}

// SetProfile sets the profile in the config file
func SetProfile(profileName string) {
	readConfigFile()

	config, err := getConfigForProfile(profileName)
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}
	if *config == (configuration.Configuration{}) {
		log.Fatal("Configuration for profile [", profileName, "] not found!")
	}

	viper.Set("profile", profileName)
	if err := viper.WriteConfig(); err != nil {
		log.Fatal("Unable to write configuration: ", err)
	}
}

// UpdateEnvName updates the EnvName in the configuration for the given profile
func UpdateEnvName(envName string) {
	readConfigFile()
	profile := viper.GetString("profile")

	// Retrieve the configuration for the specified profile
	config, err := getConfigForProfile(profile)
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}

	// Update the EnvName field
	config.EnvName = envName

	// Write the updated configuration back to the file
	viper.Set(profile, config)
	if err := viper.WriteConfig(); err != nil {
		log.Fatal("Unable to write configuration: ", err)
	}
	log.Infof("EnvName updated to [%s] successfully in profile [%s]", envName, profile)
}

// GetActiveProfileEnvName returns the EnvName for the active profile
func GetActiveProfileEnvName() string {
	readConfigFile()
	profile := viper.GetString("profile")
	config, err := getConfigForProfile(profile)
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}
	return config.EnvName
}

// EnsureEnvPresent checks if default env is present in config else asks user for env via --env param
func EnsureEnvPresent(inputEnv string) string {
	if inputEnv != "" {
		return inputEnv
	}
	env := GetActiveProfileEnvName()
	if env == "" {
		log.Fatal("Please provide the environment name using --env, or set the default environment using `odin set environment <env-name>`")
	}
	return env
}
