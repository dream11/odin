package commands

import (
	"log"
)

// --------------------------------------------------------
// Test Command
// --------------------------------------------------------
type Test struct {}

// Run implements the actual functionality of the command
// and return exit codes based on success/failure of tasks performed
func (t *Test) Run(args []string) int {
	log.Println("test run")
	return 0
}

// Help should return an explanatory string, 
// that can explain the command
func (t *Test) Help() string {
	return "command help"
}

// Synopsis should return a breif helper text for the command
func (t *Test) Synopsis() string {
	return "command synopsis"
}