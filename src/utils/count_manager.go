package utils

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type CountFileManager struct {
	FILE_PATH string
}

func (manager *CountFileManager) createDirectoryIfNotExists() error {
	DIRECTORY_PATH := manager.FILE_PATH[:strings.LastIndex(manager.FILE_PATH, "/")]
	if _, err := os.Stat(DIRECTORY_PATH); os.IsNotExist(err) {
		if err := os.MkdirAll(DIRECTORY_PATH, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (manager *CountFileManager) SaveCount(count int) error {
	if err := manager.createDirectoryIfNotExists(); err != nil {
		return err
	}

	file, err := os.OpenFile(manager.FILE_PATH, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(count))
	if err != nil {
		return err
	}

	return nil
}

func (manager *CountFileManager) GetCount() (int, error) {
	content, err := ioutil.ReadFile(manager.FILE_PATH)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(content))
}
