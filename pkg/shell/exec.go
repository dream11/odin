package shell

import (
	"bufio"
	"log"
	"os/exec"
)

func Exec(command string) int {
	log.Println("Executing:", command)
    
    cmd := exec.Command("bash", "-c", command)
	stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()

    cmd.Start()

	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
        m := scannerOut.Text()
        log.Println(m)
    }

	scannerErr := bufio.NewScanner(stderr)
    for scannerErr.Scan() {
        m := scannerErr.Text()
        log.Println(m)
    }

    cmd.Wait()
	
	return cmd.ProcessState.ExitCode()
}