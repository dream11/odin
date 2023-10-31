package config

import (
	"errors"
	"strings"
	"sync"

	"github.com/dream11/odin/api/configuration"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var onceGet sync.Once
var onceSet sync.Once
var appConfig *configuration.Configuration
var err error

func readConfig() (*configuration.Configuration, error) {
	configuration := configuration.Configuration{}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.odin")
	viper.SetEnvPrefix("ODIN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Error("Not configured odin yet? Run `odin configure`")
		}
		return nil, err
	}
	profile := viper.GetString("profile")
	err := viper.UnmarshalKey(profile, &configuration)
	if err != nil {
		log.Fatal("Configuration can't be loaded: ", err)
		return nil, err
	}
	return &configuration, nil
}

func writeConfig(config *configuration.Configuration) error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.odin")
	viper.SetEnvPrefix("ODIN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
	profile := viper.GetString("profile")
	viper.Set(profile, config)
	if err := viper.WriteConfig(); err!= nil {
		return err
	}
	return nil
}

// GetConfig returns the reference of viper config
func GetConfig() *configuration.Configuration {
	onceGet.Do(func() {
		appConfig, err = readConfig()
	})
	if err != nil {
		log.Fatal("Error while reading config ", err)
	}
	return appConfig
}

func SetConfig(config *configuration.Configuration) error {
	onceSet.Do(func() {
		err = writeConfig(config)
	})
	return err
}
