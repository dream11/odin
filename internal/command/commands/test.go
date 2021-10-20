package commands

import (
	"os"
	"flag"
	"fmt"

	"github.com/brownhash/golog"
)

// --------------------------------------------------------
// Test Command
// --------------------------------------------------------
type Test struct {}

// Run implements the actual functionality of the command
// and return exit codes based on success/failure of tasks performed
func (t *Test) Run(args []string) int {
	// Define a custom flagset
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// flag.ContinueOnError allows execution if flags have errors
	// flag.ExitOnError gracefully stops execution if flags have errors
	// flag.PanicOnError creates a panic if flags have errors

	// Add required flags to the defined flagset
	testFlag := flagSet.String("test-flag", "default value", "Help text")
	// Positional parse the flags depending upon commands and sub commands
	flagSet.Parse(os.Args[2:])
	// use the parsed flags
	golog.Debug(fmt.Sprintf("-test-flag=%s", *testFlag))

	golog.Success("Test Run!")
	return 0
}

// Help should return an explanatory string, 
// that can explain the command
func (t *Test) Help() string {
	return `Usage: d11-cli test [Options]

Options:
	--test-flag="string value"`
}

// Synopsis should return a breif helper text for the command
func (t *Test) Synopsis() string {
	return "command synopsis"
}