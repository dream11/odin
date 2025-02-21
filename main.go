package main

import (
	"encoding/json"
	"fmt"
	odin "github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/cli"
	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/request"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
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
		logger.Debug("Error making http request to fetch latest version: " + res.Error.Error())
		return ""
	}
	if res.StatusCode != 200 {
		logger.Debug("Invalid status code while checking latest version of Odin: " + fmt.Sprint(res.StatusCode))
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
	// Check if the new odin cli binary is loaded in /usr/local/bin/odin-*
	// If it is not, then run the upgrade script
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := filepath.Join(dir, ".odin/odin-*")
	files, err := filepath.Glob(path)
	if len(files) == 0 {
		// Run bash command to download script
		logger.Info("Upgrading odin")
		cmd := exec.Command("bash", "-c", "curl --silent https://artifactory.dream11.com/migrarts/odin-artifact/migrate-script.sh | bash")
		// Execute the command and display the output
		// Get pipes for stdout and stderr
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		// Start the command
		if err := cmd.Start(); err != nil {
			panic(err)
		}

		// Pipe stdout and stderr to Go’s stdout and stderr
		go io.Copy(os.Stdout, stdout)
		go io.Copy(os.Stderr, stderr)

		// Wait for the command to finish
		if err := cmd.Wait(); err != nil {
			panic(err)
		}

		logger.Warn("Please restart the terminal to use the latest version of odin or run 'source ~/.zshrc' or 'source ~/.bashrc'")
	}

	c := cli.Cli(odin.App.Name, odin.App.Version)
	exitStatus, err := c.Run()

	latestVersion := getLatestVersion()
	if latestVersion != "" && !isLatestVersion(odin.App.Version, latestVersion) {
		logger.Info(fmt.Sprintf("\nYou are using odin version %s; however, version %s is available", odin.App.Version, latestVersion))
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
