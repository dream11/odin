# Hiding a command from help section

To hide a command and prevent it from appearing in cli's help section,

1. Add a command, by following this [Adding a command doc](./ADD_COMMAND.md)

2. Now, add the commands created in step 4 above as an element to `hidenCommands` variable at [cli.go](../internal/cli/cli.go)

    ```go
    var hidenCommands = []string{
        "create test",
        "delete test",
        "list test",
        "describe test",
        "deploy test",
        "destroy test",
    }
    ```

3. Verify,

    ```shell
    go run main.go <verb> --help
    ```
