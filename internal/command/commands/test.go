package commands

import (
	"flag"
	"fmt"
)

// Test : Sample command declaration
type Test command

// Run implements the actual functionality of the command
// and return exit codes based on success/failure of tasks performed
func (t *Test) Run(args []string) int {
	// Define a custom flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// flag.ContinueOnError allows execution if flags have errors
	// flag.ExitOnError gracefully stops execution if flags have errors
	// flag.PanicOnError creates a panic if flags have errors

	// Add required flags to the defined flag set
	testFlag := flagSet.String("test-flag", "default value", "Help text")
	// Positional parse the flags depending upon commands and sub commands
	err := flagSet.Parse(args)
	if err != nil {
		t.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	// use the parsed flags
	t.Logger.Info(fmt.Sprintf("-test-flag=%s", *testFlag))

	if t.Create {
		// Perform stuff for record creation of test resource
		t.Logger.Info(fmt.Sprintf("Test Run(create)! flag value = %s", *testFlag))
		return 0
	}

	if t.Delete {
		// Perform stuff for record deletion of test resource
		t.Logger.Info(fmt.Sprintf("Test Run(delete)! flag value = %s", *testFlag))
		return 0
	}

	if t.List {
		// Perform stuff to list all test resource
		t.Logger.Info(fmt.Sprintf("Test Run(list)! flag value = %s", *testFlag))
		return 0
	}

	if t.Describe {
		// Perform stuff to describe a test resource
		t.Logger.Info(fmt.Sprintf("Test Run(describe)! flag value = %s", *testFlag))
		return 0
	}

	if t.Label {
		// Perform stuff to describe a test resource
		t.Logger.Info(fmt.Sprintf("Test Run(label)! flag value = %s", *testFlag))
		return 0
	}

	if t.Status {
		// Perform stuff to describe a test resource
		t.Logger.Info(fmt.Sprintf("Test Run(status)! flag value = %s", *testFlag))
		return 0
	}

	if t.Logs {
		// Perform stuff to describe a test resource
		t.Logger.Info(fmt.Sprintf("Test Run(logs)! flag value = %s", *testFlag))
		return 0
	}

	if t.Deploy {
		// Perform stuff to deploy a test resource
		t.Logger.Info(fmt.Sprintf("Test Run(deploy)! flag value = %s", *testFlag))
		return 0
	}

	if t.Destroy {
		// Perform stuff to destroy a test resource
		t.Logger.Info(fmt.Sprintf("Test Run(destroy)! flag value = %s", *testFlag))
		return 0
	}

	t.Logger.Error("Not a valid command")
	return 127
}

// Help should return an explanatory string,
// that can explain the command's verbs
func (t *Test) Help() string {
	if t.Create {
		return commandHelper("create", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	if t.Delete {
		return commandHelper("delete", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	if t.List {
		return commandHelper("list", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	if t.Describe {
		return commandHelper("describe", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	if t.Label {
		return commandHelper("lable", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	if t.Logs {
		return commandHelper("logs", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	if t.Deploy {
		return commandHelper("deploy", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	if t.Destroy {
		return commandHelper("destroy", "test", "", []Options{
			{Flag: "--test-flag", Description: "required value"},
		})
	}

	return defaultHelper()
}

// Synopsis should return a brief helper text for the command's verbs
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

	if t.Label {
		return "label a test resource"
	}

	if t.Status {
		return "current status of test resource"
	}

	if t.Logs {
		return "execution logs of test resource"
	}

	if t.Deploy {
		return "deploy a test resource"
	}

	if t.Destroy {
		return "destroy a test resource"
	}

	return defaultHelper()
}
