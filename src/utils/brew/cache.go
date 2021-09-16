package brew

import (
	"brew-updates/utils"
	"os"
)

func GetUpdatablePackagesCacheManager() utils.CountFileManager {
	HOME_DIRECTORY, _ := os.UserHomeDir()
	return utils.CountFileManager{
		FILE_PATH: HOME_DIRECTORY + "/.local/share/brew_updates/updatable_packages_count",
	}
}
