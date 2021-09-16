package history

import (
	"brew-updates/utils"
	"os"
)

func GetCountFileManager() utils.CountFileManager {
	HOME_DIRECTORY, _ := os.UserHomeDir()
	return utils.CountFileManager{
		FILE_PATH: HOME_DIRECTORY + "/.local/share/brew_updates/upgrade_count",
	}
}
