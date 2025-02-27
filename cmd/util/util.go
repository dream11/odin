package util

import (
	"fmt"

	"github.com/dream11/odin/internal/ui"
	log "github.com/sirupsen/logrus"
)

func AskForConfirmation(env string) {
	consentMessage := fmt.Sprintf("\nYou are executing the above command on a restricted environment. Are you sure? Enter \033[1m%s\033[0m to continue:", env)
	inputHandler := ui.Input{}
	val, err := inputHandler.Ask(consentMessage)
	if err != nil {
		log.Fatal(err.Error())
	}
	if val != env {
		log.Fatal(fmt.Errorf("aborting the operation"))
	}
}
