# <img src="./docs/odin-logo.jpg" alt="Odin" title="Odin: Internal framework CLI for CRUD operations on environments." width="30%" height="30%">

Internal framework CLI for CRUD operations on environments.

## Setup

1. [Download](https://golang.org/dl/go1.17.1.darwin-amd64.pkg) and run Go setup

2. Verify Go: `go version`

3. Clone this repo: `git clone https://github.com/dream11/odin`

4. Enter the repository: `cd odin`

5. Install dependencies: `go mod download`

6. Verify the cli: `go run main.go --version`

## Build & Install

1. Build the executable: `go build .`

2. Move to binary to system path: `sudo mv ./odin /usr/local/bin`

3. Verify the cli: `odin --version`

## Commands

### Structure

All commands are formated as: odin `<verb>` `<resource>` `<options>`

Here,

1. `verb` - The action to be performed. Suported verbs are -

    - `create` - For creating/uploading a resource record for future access.
    - `delete` - For deleting a created resource record.
    - `list` - For listing all the active entities of that resource.
    - `describe` - For describing a particular resource.
    - `deploy` - For deploying the resource on actual infrastructure.
    - `destroy` - For destroying a deployed resource.

2. `resource` - The resource on which action will be performed. Example: `env` -

    ```shell
    odin <verb> env <options>
    ```

3. `options` - Extra properties required to support the commands. Example: `env` -

    ```shell
    odin <verb> env --name=demo-1234
    ```

### [Adding a command](./docs/ADD_COMMAND.md)

### [Hidden commands](./docs/HIDDEN_COMMAND.md)

## Command Line User Interface

To interacte with user via Command Line,

```go
import (
    "github.com/dream11/odin/internal/commandline"
)
```

### Logging

```go
func main() {
    commandline.Interface.Info("string")
    commandline.Interface.Warn("string")
    commandline.Interface.Error("string") // This should be followed by an exit call
}
```

### Inputs

```go
func main() {
    text := commandline.Interface.Ask("Input text")
    secretText := commandline.Interface.AskSecret("Input secret text")
}
```

## How to contribute?

If you made changes in `xyz.go`, then ensure running

1. `go vet xyz.go`

2. `go fmt xyz.go`
