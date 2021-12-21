# <img src="./docs/odin-logo.jpg" alt="Odin" title="Odin: Internal framework CLI for CRUD operations on environments." width="30%" height="30%">

Internal framework CLI for CRUD operations on environments.

## Setup

1. [Download](https://golang.org/dl/go1.17.1.darwin-amd64.pkg) and run the Go installer.

2. Verify Go: `go version`

3. Clone this repo: `git clone https://github.com/dream11/odin`

4. Enter the repository: `cd odin`

5. Install dependencies: `go mod download`

6. Verify the cli: `go run main.go --version`

## Build & Install

1. Run `make install`

Or,

1. Build the executable: `go build .`

2. Move to binary to system path: `sudo mv ./odin /usr/local/bin`

3. Verify the cli: `odin --version`

## Contribution guide

### Formatting the code

1. Install Lint tool: `brew install golangci-lint`
2. Upgrade to its latest version: `brew upgrade golangci-lint`
3. Run linter: `make lint`

> All these linting checks are also ensured in pre-commit checks provided below.

### Making commits

1. Install pre-commit: `pip install pre-commit`
2. Setup pre-commit: `cd odin && pre-commit install`
3. Now, make commits.

> Now on every commit that you make, pre-commit hook will validate the `go` code and will suggest changes if any.

### Managing the version

Odin's application version is in the semantic version form i.e. `x.y.z` where, 
`x` is the major version, `y` is the minor version and `z` is the patch version.

The version is maintained in [odin/app/app.go](./app/app.go) inside variable named `App`.

Example: if the current version is `1.0.0`, then in case of 

1. a bug fix or a patch, the patch version is upgraded i.e. `1.0.1`
2. a minor change in some existing feature, the minor version is upgraded i.e. `1.1.0`
3. a major feature addition, the major version is upgraded i.e. `2.0.0`

> Update the version responsibly as this version will be used to create a release against any master/main branch merge.

## Commands

### Structure

All commands are formatted as: odin `<verb>` `<resource>` `<options>`

Here,

1. `verb` - The action to be performed. Supported verbs are -

    - `create` - For creating/uploading a resource record for future access.
    - `delete` - For deleting a created resource record.
    - `update` - For updating 
    - `label` - For attaching a label to a resource.
    - `list` - For listing all the active entities of that resource.
    - `describe` - For describing a particular resource.
    - `status` - For current status of a particular resource
    - `logs` - For execution logs of resource
    - `deploy` - For deploying the resource on actual infrastructure.
    - `destroy` - For destroying a deployed resource.

2. `resource` - The resource on which action will be performed. Example: `env` -

    ```shell
    odin <verb> environment <options>
    ```

3. `options` - Extra properties required to support the commands. Example: `env` -

    ```shell
    odin <verb> environment --name=demo-1234
    ```

### [Adding a command](./docs/ADD_COMMAND.md)

### [Hidden commands](./docs/HIDDEN_COMMAND.md)

## Command Line User Interface

To interact with user via Command Line,

```go
import (
    "github.com/dream11/odin/internal/ui"
)
```

### Logging

```go
var Logger ui.Logger

func main() {
    Logger.Info("string")
    Logger.Success("string")
    Logger.Warn("string")
    Logger.Output("string")
    Logger.Debug("string")
    Logger.Error("string") // This should be followed by an exit call
}
```

> To enable debug logs, set an environment variable `ODIN_DEBUG=yes` & unset to disable if already set.

### Inputs

```go
var Input ui.Logger

func main() {
    text := Input.Ask("Input text")
    secretText := Input.AskSecret("Input secret text")
}
```

> For using these interfaces(logging & inputs) within commands, these are available as part of the declared command struct. Follow step 2 [here](./docs/ADD_COMMAND.md).

## Exit status

Define appropriate exit status for commands based on the success/errors,

| Exit Code | Description |
| --------- | ----------- |
| 0 | Exit status 0 denotes that your command ran successfully |
| 1 | Exit status 1 corresponds to any generic error |
| 2 | Exit status 2 is a permission denied error |
| 126 | Exit status 126 is an interesting permissions error code. |
| 127 | Exit status 127 tells you that one of two things has happened: Either the command doesn't exist, or the command isn't in your path `$PATH` |
| 128 | Exit status 128 is the response received when an out-of-range exit code is used in programming. |
