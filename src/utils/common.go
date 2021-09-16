package utils

import (
	"os/exec"
	"runtime"
	"strings"
)

func ExecuteCommand(command string) string {
	splittedCommand := strings.Split(command, " ")
	c, b := exec.Command(splittedCommand[0], splittedCommand[1:]...), new(strings.Builder)
	c.Stdout = b
	c.Run()
	return b.String()
}

func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}
