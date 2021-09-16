package history

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

type HistoryFileManager struct {
	FILE_PATH string
}

func GetHistoryFileManager() HistoryFileManager {
	return HistoryFileManager{FILE_PATH: "/var/tmp/shell_history"}
}

func (history HistoryFileManager) UpdateHistory() {
	command := fmt.Sprintf("history --show-time > %s", history.FILE_PATH)
	exec.Command(command).Run()
}

func (history HistoryFileManager) GetHistory() (string, error) {
	content, err := ioutil.ReadFile(history.FILE_PATH)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
