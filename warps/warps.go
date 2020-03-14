package warps

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Entry struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Warps struct {
	Warps map[string]Entry `json:"warps"`
}

func checkExist(path string) bool {
	stats, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !stats.IsDir()
}

func createNewDatabaseFile(path string) error {
	database := Warps{}
	serialized, err := json.Marshal(database)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, serialized, 0755)
	return err
}

func Load(path string) (*Warps, error) {
	var fileContents []byte
	var err error
	if !checkExist(path) {
		err = createNewDatabaseFile(path)
		if err != nil {
			return nil, err
		}
	}

	fileContents, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	database := Warps{}

	err = json.Unmarshal(fileContents, &database)
	if err != nil {
		return nil, err
	}

	return &database, nil
}