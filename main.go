package main

import (
	"encoding/json"
	"fmt"
	"os"

	odin "github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/cli"
	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/request"
)

var logger ui.Logger

const GITHUB_TAGS_URL = "https://api.github.com/repos/dream11/odin/tags"

func getLatestVersion() string {
	req := request.Request{
		Method: "GET",
		URL:    GITHUB_TAGS_URL,
	}
	res := req.Make()
	if res.Error != nil {
		logger.Debug("Error making http req to fetch latest version: " + res.Error.Error())
		return ""
	}
	var jsonResponse []map[string]interface{}
	err := json.Unmarshal(res.Body, &jsonResponse)
	if err != nil {
		logger.Debug("Unable to unmarshal latest version response : " + err.Error())
		return ""
	}

	// return the latest tag
	return jsonResponse[0]["name"].(string)
}

func isLatestVersion(currentVersion string, latestVersion string) bool {
	return currentVersion == latestVersion
}

func main() {
	c := cli.Cli(odin.App.Name, odin.App.Version)
	exitStatus, err := c.Run()

	latestVersion := getLatestVersion()
	if latestVersion != "" && !isLatestVersion(odin.App.Version, latestVersion) {
		logger.Info(fmt.Sprintf("You are using odin version %s; however, version %s is available", odin.App.Version, latestVersion))
		logger.Info("Upgrade to the latest version via command 'brew install dream11/tools/odin'")
	}

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	os.Exit(exitStatus)
}

// TODO: https://github.com/mitchellh/go-glint
// TODO: https://github.com/charmbracelet/bubbletea for advanced interactions with user
