package commands

import (
	"log"
)

type Test struct {}

func (t *Test) Run(args []string) int {
	log.Println("test run")
	return 0
}

func (t *Test) Help() string {
	return "command help"
}

func (t *Test) Synopsis() string {
	return "command synopsis"
}