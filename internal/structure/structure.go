package structure

import (
	"os"
	"path"

	"github.com/dream11/d11-cli/pkg/dir"
)

type workdir struct {
	Location    string
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
	Location: path.Join(os.Getenv("HOME"), ".d11-cli"),
}

