package odin

type application struct {
	Name    string
	Version string
}

var App application = application{
	Name:    "odin",
	Version: "1.0.0-beta",
}
