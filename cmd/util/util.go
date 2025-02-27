package util

import (
	"fmt"

	"github.com/dream11/odin/internal/ui"
	log "github.com/sirupsen/logrus"
)

// AskForConfirmation asks for confirmation before proceeding with the operation
func AskForConfirmation(expectedValue, consentMessage string) {
	inputHandler := ui.Input{}
	val, err := inputHandler.Ask(consentMessage)
	if err != nil {
		log.Fatal(err.Error())
	}
	if val != expectedValue {
		log.Fatal(fmt.Errorf("aborting the operation"))
	}
}
