/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/dream11/odin/cmd"
	_ "github.com/dream11/odin/cmd/configure"
	_ "github.com/dream11/odin/cmd/create"
	_ "github.com/dream11/odin/cmd/delete"
	_ "github.com/dream11/odin/cmd/deploy"
	_ "github.com/dream11/odin/cmd/describe"
	_ "github.com/dream11/odin/cmd/list"
	_ "github.com/dream11/odin/cmd/operate"
	_ "github.com/dream11/odin/cmd/set"
	_ "github.com/dream11/odin/cmd/status"
	_ "github.com/dream11/odin/cmd/undeploy"
	_ "github.com/dream11/odin/internal/ui"
)

func main() {
	cmd.Execute()
}
