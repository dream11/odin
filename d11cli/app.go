package d11cli

type application struct {
	Name       string
	Version    string
}

var App application = application{
	Name: "d11-cli",
	Version: "1.0.0-beta",
}