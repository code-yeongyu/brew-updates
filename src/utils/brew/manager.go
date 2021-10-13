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

func (b BrewManager) GetCurrentUpgradablePackages() []string {
	outdatedPackagesString := utils.ExecuteCommand("brew outdated")
	if outdatedPackagesString == "" {
		return []string{}
	}
	outdatedPackages := strings.Split(outdatedPackagesString, "\n")
	return outdatedPackages
}

func refreshUpdatableListManager(b BrewManager, updatableListManager *UpdatableListManager) {
	b.UpdateInfo()
	updatableListManager.Packages = b.GetCurrentUpgradablePackages()

	if err := updatableListManager.Save(); err != nil {
		panic(err)
	}
}

func (b BrewManager) GetUpgradablePackages() []string {
	updatableListManager := GetUpdatableListManager()

	// load the cached packages, or create a new one if it doesn't exist
	if err := updatableListManager.Load(); err != nil {
		refreshUpdatableListManager(b, updatableListManager)
		if err := updatableListManager.Load(); err != nil {
			panic(err)
		}
	}

	if !updatableListManager.IsAvailable() {
		refreshUpdatableListManager(b, updatableListManager)
	}

	return updatableListManager.Packages
}
