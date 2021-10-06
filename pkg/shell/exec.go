package shell

import (
    "fmt"
	"bufio"
	"os/exec"

    "github.com/brownhash/golog"
)

func Exec(command string) int {
    golog.Debug(fmt.Sprintf("Executing: %s", command))
    
    cmd := exec.Command("bash", "-c", command)
	stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()

    cmd.Start()

	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
        m := scannerOut.Text()
        golog.Println(m)
    }

	scannerErr := bufio.NewScanner(stderr)
    for scannerErr.Scan() {
        m := scannerErr.Text()
        golog.Println(m)
    }

    cmd.Wait()
	
	return cmd.ProcessState.ExitCode()
}