package utils

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

// INFO:  just a dummy function to run commands
func RunCommand(command string, args ...string) ([]byte, error) {
	defer timeTrack(time.Now(), "COMMAND RUNNER")
	myCmd := exec.Command(command, args...)
	output, err := myCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	fmt.Printf(" == Running %s with args %v == \n", command, args)

	return output, nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
