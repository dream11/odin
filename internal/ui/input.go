package ui

import (
	"fmt"
)

// Input : interface declaration
type Input struct{}

// Ask : asks for a generic text ui
func (i *Input) Ask(description string) (string, error) {
	return userInterface.Ask(description)
}

// AskWithConstraints : asks for a generic text ui which are part of the constraints provided
func (i *Input) AskWithConstraints(description string, constraints map[string]struct{}) (string, error) {
	for {
		val, err := i.Ask(description)

		if err != nil {
			return "", err
		}

		_, ok := constraints[val]

		if ok {
			return val, nil
		}

		fmt.Println("Invalid ui, retry")
	}
}

// AskSecret : asks for a secret text ui
func (i *Input) AskSecret(description string) (string, error) {
	return userInterface.AskSecret(description)
}
