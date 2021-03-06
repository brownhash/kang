package shell

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"

	"github.com/brownhash/golog"
)

// bash command executioner
// return exit code and logs the outputs
func Exec(command string) int {
    golog.Debug(fmt.Sprintf("Executing: %s", command))
    
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