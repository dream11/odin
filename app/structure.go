package app

import (
	"github.com/dream11/odin/pkg/dir"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

type workdir struct {
	Location     string
	ConfigFile   string
	EnvVarPrefix string
}

// Create : Creates the required working directory
func (w *workdir) Create() error {
	wExists, err := dir.Exists(w.Location)
	if err != nil {
		return err
	}

	if wExists {
		return nil
	}

	return dir.Create(w.Location, 0755)
}

// WorkDir interface
var WorkDir = workdir{
	Location:     path.Join(os.Getenv("HOME"), "."+App.Name),
	ConfigFile:   "config",
	EnvVarPrefix: strings.ToUpper(App.Name) + "_",
}

// initiate dir structure on app initialization
func init() {
	err := WorkDir.Create()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	secretCredentialsExist, err := dir.Exists(path.Join(WorkDir.Location, WorkDir.ConfigFile))
	if err != nil {
		log.Error(err.Error())
	}

	if !secretCredentialsExist {
		log.Warn("Run, `odin configure` to configure odin")
	}
}
