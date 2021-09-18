package main

import (
	"brew-updates/utils"
	"brew-updates/utils/brew"
	"fmt"
)

func checkExecutable(brewManager brew.BrewManager) {
	if !utils.IsMacOS() {
		panic("Only macOS is supported.")
	}
	if !brewManager.IsInstalled() {
		panic("Brew is not installed.")
	}
}

func getUpgradablePackagesCount(brewManager brew.BrewManager) int {
	brewManager.UpdateInfo()
	upgradablePackages := brewManager.GetUpgradablePackages()
	return len(upgradablePackages)
}

func main() {
	brewManager := brew.BrewManager{}
	checkExecutable(brewManager)
	upgradablePackagesCnt := getUpgradablePackagesCount(brewManager)

	fmt.Printf("%d updates can be applied immediately.\n", upgradablePackagesCnt)
	if upgradablePackagesCnt > 10 {
		fmt.Printf(
			"To see these additional updates run: brew outdated\nTo upgrade all of your outdated packages run: brew upgrade",
		)
	}
}
