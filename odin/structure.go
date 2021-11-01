package odin

import (
	"os"
	"path"

	"github.com/brownhash/golog"
	"github.com/dream11/odin/pkg/dir"
)

type workdir struct {
	Location string
}

func (w *workdir) Create() error {
	wExists, err := dir.Exists(w.Location)
	if err != nil {
		return err
	}

	if wExists {
		return nil
	}

	return dir.Create(w.Location, 755)
}

var WorkDir workdir = workdir{
	Location: path.Join(os.Getenv("HOME"), "."+App.Name),
}

// initiate dir structure on app initialization
func init() {
	err := WorkDir.Create()
	if err != nil {
		golog.Error(err)
	}
}
