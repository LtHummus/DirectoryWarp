package warps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/logrusorgru/aurora"
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
	database := Warps{make(map[string]Entry)}
	serialized, err := json.Marshal(database)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, serialized, 0755)
	fmt.Printf("Wrote new warps database file at %s\n", Yellow(path))
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

func (warps *Warps) Add(name string, path string) bool {
	var overwrite bool
	_, exists := warps.Warps[name]
	if exists {
		overwrite = true
	}
	warps.Warps[name] = Entry{
		Name: name,
		Path: path,
	}

	return overwrite
}

func (warps *Warps) Write(path string) error {
	serialized, err := json.Marshal(warps)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, serialized, 0755)
	return err
}