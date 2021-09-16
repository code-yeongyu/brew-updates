package main

import (
	"brew-updates/utils"
	"brew-updates/utils/brew"
	"brew-updates/utils/history"
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

func getRecordedBrewUpgradeCount() int {
	countFileManager := history.GetCountFileManager()
	recordedBrewUpgradeCount, _ := countFileManager.GetCount()
	return recordedBrewUpgradeCount
}

func getCurrentBrewUpgradeCount() int {
	historyFileManager := history.GetHistoryFileManager()
	historyFileManager.UpdateHistory()
	commandHistory, _ := historyFileManager.GetHistory()
	return history.CountBrewUpgrade(commandHistory)
}

func isBrewUpgradeCountHasDiffer(currentBrewUpgradeCount int) bool {
	recordedBrewUpgradeCount := getRecordedBrewUpgradeCount()
	return recordedBrewUpgradeCount != currentBrewUpgradeCount
}

func getUpgradablePackagesCount(brewManager brew.BrewManager) int {
	brewManager.UpdateInfo()
	upgradablePackages := brewManager.GetUpgradablePackages()
	return len(upgradablePackages)
}

func main() {
	brewManager := brew.BrewManager{}
	checkExecutable(brewManager)

	currentBrewUpgradeCount := getCurrentBrewUpgradeCount()
	updatablePackageCacheManager := brew.GetUpdatablePackagesCacheManager()

	if isBrewUpgradeCountHasDiffer(currentBrewUpgradeCount) {
		historyCountManager := history.GetCountFileManager()
		if err := historyCountManager.SaveCount(currentBrewUpgradeCount); err != nil {
			panic(err)
		}

		upgradablePackagesCount := getUpgradablePackagesCount(brewManager)
		if err := updatablePackageCacheManager.SaveCount(upgradablePackagesCount); err != nil {
			panic(err)
		}
	}

	upgradablePackagesCnt, _ := updatablePackageCacheManager.GetCount()

	fmt.Printf("%d updates can be applied immediately.", upgradablePackagesCnt)
	if upgradablePackagesCnt > 10 {
		fmt.Printf(
			"To see these additional updates run: brew outdated\nTo upgrade all of your outdated packages run: brew upgrade",
		)
	}
}
