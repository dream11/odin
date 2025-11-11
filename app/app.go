package app

type application struct {
	Name    string
	Version string
}

// App (Application) interface
var App = application{
	Name:    "odin",
	Version: "2.1.2",
}
