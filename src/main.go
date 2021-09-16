package main

import (
	"fmt"
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

type BrewManager struct {
}

func (b BrewManager) IsInstalled() bool {
	result := ExecuteCommand("brew --version")
	return result != ""
}

func (b BrewManager) UpdateInfo() {
	ExecuteCommand("brew update")
}

func (b BrewManager) GetUpgradablePackages() []string {
	outdatedPackagesString := ExecuteCommand("brew outdated")
	if outdatedPackagesString == "" {
		return []string{}
	}
	outdatedPackages := strings.Split(outdatedPackagesString, "\n")
	return outdatedPackages
}

func main() {

}
