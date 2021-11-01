# !["Odin"](./docs/odin-logo.jpg "Odin: Internal framework CLI for CRUD operations on environments")

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

## Logging

### Log Levels

To set/change logging levels for the app, provide this variable in the environemnt: `D11_LOG_LEVEL=<log_level>`,

Supported logging levels,

1. `DEBUG` or `10`

2. `INFO` or `20` | default

3. `WARN` or `30`

4. `SUCCESS` or `30`

5. `ERROR` or `40`

> Note: The above logging levels are in decreasing order of their precedence. i.e. `DEBUG` will allow all loggers to log, while `ERROR` wont allow any other than itself.

## How to contribute?

If you made changes in `xyz.go`, then ensure running

1. `go vet xyz.go`

2. `go fmt xyz.go`
