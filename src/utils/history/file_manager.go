package history

import (
	"brew-updates/utils"
	"fmt"
	"io/ioutil"
)

type HistoryFileManager struct {
	FILE_PATH string
}

func GetHistoryFileManager() HistoryFileManager {
	return HistoryFileManager{FILE_PATH: "/var/tmp/shell_history"}
}

func (history HistoryFileManager) UpdateHistory() {
	command := fmt.Sprintf("history --show-time > %s", history.FILE_PATH)
	utils.ExecuteCommand(command)
}

func (history HistoryFileManager) GetHistory() (string, error) {
	content, err := ioutil.ReadFile(history.FILE_PATH)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
