# Adding a command `test` with all its `verbs`

1. Create [test.go](../internal/command/commands/test.go) at `odin/internal/command/commands/` and initiate it,

    ```go
    package commands
    ```

2. Import the necessary packages and initiate the command struct,

    ```go
    import (
        "os"
        "flag"
        "fmt"

        "github.com/dream11/odin/internal/ui" // cli interface
    )

    type Test command
    ```

3. Initiate the required functions for command,

    - `Run(args []string) int {}` - Accepts a list of arguments and return an exit code after processing command based on verbs.

        ```go
        // Run implements the actual functionality of the command
        // and return exit codes based on success/failure of tasks performed
        func (t *Test) Run(args []string) int {
            // Define a custom flagset
            flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)

            // Add required flags to the defined flagset
            testFlag := flagSet.String("test-flag", "default value", "Help text")
            // Positional parse the flags depending upon commands and sub commands
            flagSet.Parse(os.Args[3:])
            // use the parsed flags
            ui.Interface().Info(fmt.Sprintf("-test-flag=%s", *testFlag))

            if t.Create {
                // Perform stuff for record creation of test resource
                ui.Interface().Info(fmt.Sprintf("Test Run(create)! flag value = %s", *testFlag))
                return 0
            }
            if t.Delete {
                // Perform stuff for record deletion of test resource
                ui.Interface().Info(fmt.Sprintf("Test Run(delete)! flag value = %s", *testFlag))
                return 0
            }
            if t.List {
                // Perform stuff to list all test resource
                ui.Interface().Info(fmt.Sprintf("Test Run(list)! flag value = %s", *testFlag))
                return 0
            }
            if t.Describe {
                // Perform stuff to describe a test resource
                ui.Interface().Info(fmt.Sprintf("Test Run(describe)! flag value = %s", *testFlag))
                return 0
            }
            if t.Deploy {
                // Perform stuff to deploy a test resource
                ui.Interface().Info(fmt.Sprintf("Test Run(deploy)! flag value = %s", *testFlag))
                return 0
            }
            if t.Destroy {
                // Perform stuff to destroy a test resource
                ui.Interface().Info(fmt.Sprintf("Test Run(destroy)! flag value = %s", *testFlag))
                return 0
            }

            ui.Interface().Error("Not a valid command")
            return 1
        }
        ```

    - `Help() string {}` - Accepts nothing and returns an explanatory string.

        ```go
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
        ```

    - `Synopsis() string {}` - Accepts nothing and returns a short string.

        ```go
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
        ```

4. Create the commands at [command_catalog.go](../internal/command/command_catalog.go) at `odin/internal/command/command_catalog.go` in the `CommandCatalog()` function,

    For mapping `command` to `verb`, define the command as,

    ```go
    "verb resource": func() (cli.Command, error) {
        return &commands.Resource{Verb: true}, nil
    }
    ```

    Like,

    ```go
    func CommandCatalog() map[string]cli.CommandFactory {
        return map[string]cli.CommandFactory{
            "create test": func() (cli.Command, error) {
                return &commands.Test{Create: true}, nil
            },
            "delete test": func() (cli.Command, error) {
                return &commands.Test{Delete: true}, nil
            },
            "list test": func() (cli.Command, error) {
                return &commands.Test{List: true}, nil
            },
            "describe test": func() (cli.Command, error) {
                return &commands.Test{Describe: true}, nil
            },
            "deploy test": func() (cli.Command, error) {
                return &commands.Test{Deploy: true}, nil
            },
            "destroy test": func() (cli.Command, error) {
                return &commands.Test{Destroy: true}, nil
            },
        }
    }
    ```

> Remove the verbs which are not in scope of that resource.
