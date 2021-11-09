package shell

import (
	"bufio"
	"os/exec"

	"github.com/dream11/odin/internal/commandline"
)

func Exec(command string) int {
	commandline.Interface.Warn("Executing:" + command)

	cmd := exec.Command("bash", "-c", command)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	cmd.Start()

	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
		m := scannerOut.Text()
		commandline.Interface.Output(m)
	}

	scannerErr := bufio.NewScanner(stderr)
	for scannerErr.Scan() {
		m := scannerErr.Text()
		commandline.Interface.Error(m)
	}

	cmd.Wait()

	return cmd.ProcessState.ExitCode()
}
