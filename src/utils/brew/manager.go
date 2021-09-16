package brew

import (
	"brew-updates/utils"
	"strings"
)

type BrewManager struct {
}

func (b BrewManager) IsInstalled() bool {
	result := utils.ExecuteCommand("brew --version")
	return result != ""
}

func (b BrewManager) UpdateInfo() {
	utils.ExecuteCommand("brew update")
}

func (b BrewManager) GetUpgradablePackages() []string {
	outdatedPackagesString := utils.ExecuteCommand("brew outdated")
	if outdatedPackagesString == "" {
		return []string{}
	}
	outdatedPackages := strings.Split(outdatedPackagesString, "\n")
	return outdatedPackages
}
