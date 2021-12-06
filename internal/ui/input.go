package ui

// Input : interface declaration
type Input struct{}

// Ask : asks for a generic text input
func (i *Input) Ask(description string) (string, error) {
	return userInterface.Ask(description)
}

// AskSecret : asks for a secret text input
func (i *Input) AskSecret(description string) (string, error) {
	return userInterface.AskSecret(description)
}
