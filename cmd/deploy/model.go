package deploy

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

type Model struct {
	ServiceDisplayMeta ServiceDisplayMeta
	ServiceView        ServiceView
}

type ServiceDisplayMeta struct {
	Height               int
	Width                int
	Ready                bool
	Cursor               int
	ElapsedTime          int
	TotalCompletionTime  int
	Progress             progress.Model
	ComponentDisplayMeta []ComponentDisplayMeta
}

type ComponentDisplayMeta struct {
	Height      int
	Width       int
	Toggle      bool
	Spinner     spinner.Model
	LogViewPort viewport.Model
}

type ServiceView struct {
	Name           string
	Status         string
	Action         string
	TraceId        string
	ComponentsView []ComponentView
}

type ComponentView struct {
	Name    string
	Status  string
	Action  string
	Content string
}
