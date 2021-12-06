package shell

import (
	"bufio"
	"os/exec"

	"github.com/dream11/odin/internal/ui"
)

var logger ui.Logger

// Exec : execute given command
func Exec(command string) int {
	logger.Warn("Executing:" + command)

	cmd := exec.Command("bash", "-c", command)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		logger.Error("Unable to start cmd execution. " + err.Error())
		return 1
	}

	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
		m := scannerOut.Text()
		logger.Output(m)
	}

	scannerErr := bufio.NewScanner(stderr)
	for scannerErr.Scan() {
		m := scannerErr.Text()
		logger.Error(m)
	}

	err = cmd.Wait()
	if err != nil {
		logger.Error(err.Error())
		return 1
	}

	return cmd.ProcessState.ExitCode()
}
