package commands

import (
	"flag"
	"os"
	"path"

	"github.com/dream11/odin/api/configuration"
	"github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"gopkg.in/yaml.v3"
	"crypto/sha256"
	"encoding/hex"
)

// initiate backend client for auth
var authClient backend.Auth

// Configure : command declaration
type Configure command

const DEFAULT_BACKEND_ADDR = "https://odin-backend.d11dev.com"

// Run : implements the actual functionality of the command
func (c *Configure) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)

	err := flagSet.Parse(args)
	if err != nil {
		c.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	configPath := path.Join(app.WorkDir.Location, app.WorkDir.ConfigFile)

	var config configuration.Configuration

	// get backend address from env variable
	if os.Getenv("ODIN_BACKEND_ADDRESS") != "" {
		config.BackendAddr = os.Getenv("ODIN_BACKEND_ADDRESS")
	} else {
		config.BackendAddr = DEFAULT_BACKEND_ADDR
	}

	// get access key from env variable
	if os.Getenv("ODIN_ACCESS_KEY") != "" {
		config.Keys.AccessKey = os.Getenv("ODIN_ACCESS_KEY")
	} else {
		c.Logger.Error("Environment variable ODIN_ACCESS_KEY is not set. Please set your access key in ODIN_ACCESS_KEY environment variable")
		return 1
	}

	// get secret access key from env variable
	if os.Getenv("ODIN_SECRET_ACCESS_KEY") != "" {
		sha256 := sha256.New()
		sha256.Write([]byte(os.Getenv("ODIN_SECRET_ACCESS_KEY")))
		hashedResult := sha256.Sum(nil)
		config.Keys.SecretAccessKey = hex.EncodeToString(hashedResult)
	} else {
		c.Logger.Error("Environment variable ODIN_SECRET_ACCESS_KEY is not set. Please set your secret access key in ODIN_SECRET_ACCESS_KEY environment variable")
		return 1
	}

	// generate yaml
	configYaml, err := yaml.Marshal(config)
	if err != nil {
		c.Logger.Error(err.Error())
		return 1
	}

	// store configs
	err = file.Write(configPath, string(configYaml), 0755)
	if err != nil {
		c.Logger.Error("Unable to write configuration." + err.Error())
		return 1
	}

	_, err = authClient.Authenticate(config.Keys.AccessKey, config.Keys.SecretAccessKey)
	if err != nil {
		c.Logger.Error("Configuration failed. " + err.Error())
		return 1
	}

	c.Logger.Success("Configured!")
	return 0
}

// Help : returns an explanatory string
func (c *Configure) Help() string {
	return commandHelper("configure", "",
		"Set environment variables ODIN_ACCESS_KEY and ODIN_SECRET_ACCESS_KEY before using this command",
		[]Options{})
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *Configure) Synopsis() string {
	return "configure the cli authentication"
}
