package shell

import (
	"bufio"
	"os/exec"

	"github.com/dream11/odin/internal/ui"
)

func Exec(command string) int {
	ui.Interface().Warn("Executing:" + command)

	cmd := exec.Command("bash", "-c", command)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	cmd.Start()

	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
		m := scannerOut.Text()
		ui.Interface().Output(m)
	}

	scannerErr := bufio.NewScanner(stderr)
	for scannerErr.Scan() {
		m := scannerErr.Text()
		ui.Interface().Error(m)
	}

	cmd.Wait()

	return cmd.ProcessState.ExitCode()
}
