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
type Test command

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
	flagSet.Parse(os.Args[3:])
	// use the parsed flags
	golog.Debug(fmt.Sprintf("-test-flag=%s", *testFlag))

	if t.Create {
		// Perform stuff for record creation of test resource
		golog.Success(fmt.Sprintf("Test Run(create)! flag value = %s", *testFlag))
		return 0
	}
	if t.Delete {
		// Perform stuff for record deletion of test resource
		golog.Success(fmt.Sprintf("Test Run(delete)! flag value = %s", *testFlag))
		return 0
	}
	if t.List {
		// Perform stuff to list all test resource
		golog.Success(fmt.Sprintf("Test Run(list)! flag value = %s", *testFlag))
		return 0
	}
	if t.Describe {
		// Perform stuff to describe a test resource
		golog.Success(fmt.Sprintf("Test Run(describe)! flag value = %s", *testFlag))
		return 0
	}
	if t.Deploy {
		// Perform stuff to deploy a test resource
		golog.Success(fmt.Sprintf("Test Run(deploy)! flag value = %s", *testFlag))
		return 0
	}
	if t.Destroy {
		// Perform stuff to destroy a test resource
		golog.Success(fmt.Sprintf("Test Run(destroy)! flag value = %s", *testFlag))
		return 0
	}

	golog.Error("Not a valid command")

	return 1
}

// Help should return an explanatory string, 
// that can explain the command's verbs
func (t *Test) Help() string {
	if t.Create {
		return commandHelper("create", "test", []string{
			"--test-flag=required value",
		})
	}
	if t.Delete {
		return commandHelper("delete", "test", []string{
			"--test-flag=required value",
		})
	}
	if t.List {
		return commandHelper("list", "test", []string{
			"--test-flag=required value",
		})
	}
	if t.Describe {
		return commandHelper("describe", "test", []string{
			"--test-flag=required value",
		})
	}
	if t.Deploy {
		return commandHelper("deploy", "test", []string{
			"--test-flag=required value",
		})
	}
	if t.Destroy {
		return commandHelper("destroy", "test", []string{
			"--test-flag=required value",
		})
	}

	return defaultHelper()
}

// Synopsis should return a breif helper text for the command's verbs
func (t *Test) Synopsis() string {
	if t.Create {
		return "create a test resource"
	}
	if t.Delete {
		return "delete a test resource"
	}
	if t.List {
		return "list all test resources"
	}
	if t.Describe {
		return "describe a test resource"
	}
	if t.Deploy {
		return "deploy a test resource"
	}
	if t.Destroy {
		return "destroy a test resource"
	}

	return defaultHelper()
}