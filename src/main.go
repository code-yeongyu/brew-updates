package main

import (
	"brew-updates/utils/brew"
	"fmt"
)

func main() {
	brewManager := brew.BrewManager{}

	if !brewManager.IsInstalled() {
		fmt.Println("Brew is not installed.")
		return
	}

	brewManager.UpdateInfo()
	upgradablePackages := brewManager.GetUpgradablePackages()
	upgradablePackagesCnt := len(upgradablePackages)

	fmt.Printf("%d updates can be applied immediately.", upgradablePackagesCnt)
	if upgradablePackagesCnt > 10 {
		fmt.Printf(
			"To see these additional updates run: brew outdated\nTo upgrade all of your outdated packages run: brew upgrade",
		)
	}
}
