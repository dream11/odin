package deploy

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

type ServiceSetView struct {
	Ready       bool
	Cursor      int
	Header      Header
	Status      string
	ServiceView []ServiceView
}

type ServiceView struct {
	Ready          bool
	Cursor         int
	Header         Header
	Status         string
	ComponentsView []ComponentView
}

type ComponentView struct {
	Toggle  bool
	Header  Header
	Status  string
	LogView LogView
}

type Header struct {
	Toggle  bool
	Text    string
	Spinner spinner.Model
}

type LogView struct {
	Toggle      bool
	Content     string
	LogViewPort viewport.Model
}
