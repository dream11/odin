package app

type application struct {
	Name    string
	Version string
}

// App (Application) structure
var App application = application{
	Name:    "odin",
	Version: "1.0.0-beta",
}
