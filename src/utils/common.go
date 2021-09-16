package utils

import (
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) string {
	splittedCommand := strings.Split(command, " ")
	c, b := exec.Command(splittedCommand[0], splittedCommand[1:]...), new(strings.Builder)
	c.Stdout = b
	c.Run()
	return b.String()
}
