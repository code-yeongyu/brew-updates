package brew

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type UpdatableListManager struct {
	FILE_PATH          string   `json:_`
	CreatedAtTimeStamp int64    `json:createdAt`
	Packages           []string `json:packages`
}

var instance *UpdatableListManager
var once sync.Once

func GetUpdatableListManager() *UpdatableListManager {
	HOME_DIRECTORY, _ := os.UserHomeDir()
	once.Do(func() {
		instance = &UpdatableListManager{
			FILE_PATH: HOME_DIRECTORY + "/.local/share/brew_updates/updatable_packages_log.json",
		}
	})

	return instance
}

func (manager *UpdatableListManager) Save() error {
	now := time.Now()
	timestamp := now.Unix()
	manager.CreatedAtTimeStamp = timestamp

	jsonString, err := json.Marshal(manager)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(manager.FILE_PATH, jsonString, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (manager *UpdatableListManager) Load() error {
	file, err := os.Open(manager.FILE_PATH)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonString, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonString, &manager)
	if err != nil {
		return err
	}

	return nil
}

func (manager UpdatableListManager) IsAvailable() bool {
	if manager.CreatedAtTimeStamp == 0 {
		return false
	}
	createdAt := time.Unix(manager.CreatedAtTimeStamp, 0)
	expiry := createdAt.Add(time.Hour * 24)
	today := time.Now()

	return today.Before(expiry)
}
