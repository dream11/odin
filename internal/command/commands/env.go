package commands

// Env : command declaration
type Env command

// Run : implements the actual functionality of the command
func (e *Env) Run(args []string) int {
	if e.List {
		e.Logger.Info("Listing all envs")
		// TODO: call PG api that fetches a list of valid envs
		// Example: dev, stag, rev, load, uat, prod
		// GET /env

		return 0
	}

	e.Logger.Error("Not a valid command")
	return 1
}

// Help : returns an explanatory string
func (e *Env) Help() string {
	if e.List {
		return commandHelper("list", "env", []string{})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (e *Env) Synopsis() string {
	if e.List {
		return "list all valid envs"
	}

	return defaultHelper()
}
