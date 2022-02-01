package commands

import (
	"flag"
	"path"

	"github.com/dream11/odin/api/configuration"
	"github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"gopkg.in/yaml.v3"
)

// initiate backend client for auth
var authClient backend.Auth

// Configure : command declaration
type Configure command

// Run : implements the actual functionality of the command
func (c *Configure) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	refresh := flagSet.Bool("refresh", false, "refresh token using existing tokens")
	hardRefresh := flagSet.Bool("hard", false, "hard refresh token using existing keys")

	err := flagSet.Parse(args)
	if err != nil {
		c.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	configPath := path.Join(app.WorkDir.Location, app.WorkDir.ConfigFile)

	var config configuration.Configuration

	// run only when --refresh is not applied
	if *refresh {
		// read secret keys from workdir
		preConfigYaml, err := file.Read(configPath)
		if err != nil {
			c.Logger.Error("Unable to read configured keys. " + err.Error())
			return 1
		}

		// parse secret keys
		err = yaml.Unmarshal(preConfigYaml, &config)
		if err != nil {
			c.Logger.Error("Unable to parse secret keys. " + err.Error())
			return 1
		}

		if *hardRefresh {
			authResponse, err := authClient.GetToken(config.Keys.AccessKey, config.Keys.SecretAccessKey)
			if err != nil {
				c.Logger.Error("Unable to hard refresh the tokens. " + err.Error())
				return 1
			}

			config.AccessToken = authResponse.AccessToken
			config.RefreshToken = authResponse.RefreshToken
		} else {
			authResponse, err := authClient.RefreshToken(config.RefreshToken)
			if err != nil {
				c.Logger.Error("Unable to refresh the tokens. " + err.Error())
				return 1
			}

			config.AccessToken = authResponse.AccessToken
			config.RefreshToken = authResponse.RefreshToken
		}
	} else {
		// get access key from user
		config.BackendAddr, err = c.Input.Ask("Enter Backend Address:")
		if err != nil {
			c.Logger.Error(err.Error())
			return 1
		}

		// get access key from user
		config.Keys.AccessKey, err = c.Input.Ask("Enter Access Key:")
		if err != nil {
			c.Logger.Error(err.Error())
			return 1
		}

		// get secret access key from user
		config.Keys.SecretAccessKey, err = c.Input.AskSecret("Enter Secret Access Key:")
		if err != nil {
			c.Logger.Error(err.Error())
			return 1
		}

		// generate yaml
		configYaml, err := yaml.Marshal(config)
		if err != nil {
			c.Logger.Error(err.Error())
			return 1
		}

		// store pre configs
		err = file.Write(configPath, string(configYaml), 0755)
		if err != nil {
			c.Logger.Error("Unable to write configuration." + err.Error())
			return 1
		}

		authResponse, err := authClient.GetToken(config.Keys.AccessKey, config.Keys.SecretAccessKey)
		if err != nil {
			c.Logger.Error("Unable to refresh the tokens. " + err.Error())
			return 1
		}

		config.AccessToken = authResponse.AccessToken
		config.RefreshToken = authResponse.RefreshToken
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

	c.Logger.Success("Configured!")
	return 0
}

// Help : returns an explanatory string
func (c *Configure) Help() string {
	return commandHelper("configure", "", []string{
		"--refresh (to enable only token refresh using pre fetched refresh tokens)",
		"--hard (to enable token refresh using pre entered access keys | works with --refresh)",
	})
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *Configure) Synopsis() string {
	return "configure the cli authentication"
}
