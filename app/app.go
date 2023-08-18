package app

type application struct {
	Name    string
	Version string
}

// App (Application) interface
var App application = application{
	Name:    "odin",
	Version: "2.0.0",
}
