package warps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

func emptyWarps() *Warps {
	return &Warps{make(map[string]Entry)}
}

func createNewDatabaseFile(path string) error {
	database := emptyWarps()
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


func (warps *Warps) Write(path string) error {
	serialized, err := json.Marshal(warps)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, serialized, 0755)
	return err
}

func (warps *Warps) ListEntries() {
	for name, entry := range warps.Warps {
		fmt.Printf("%s -> %s\n", Cyan(name), Green(entry.Path))
	}
}

func (warps *Warps) SetEntry(name string, path string) {
	cleanedPath, err := filepath.Abs(path)
	if err != nil {
		msg := fmt.Sprintf("Can not clean path: %v", err)
		fmt.Printf("%s\n", Red(msg))
	}
	var overwritten bool
	_, exists := warps.Warps[name]
	if exists {
		overwritten = true
	}

	warps.Warps[name] = Entry{
		Name: name,
		Path: cleanedPath,
	}

	if overwritten {
		fmt.Printf("%s\n", Yellow("Overwriting old entry"))
	}

	fmt.Printf("Set %s to %s\n", Cyan(name), Green(cleanedPath))
}

func (warps *Warps) DeleteEntry(name string) bool {
	_, exists := warps.Warps[name]
	if !exists {
		fmt.Printf("Entry %s does not exist\n", Cyan(name))
		return false
	}

	delete(warps.Warps, name)
	fmt.Printf("Entry %s deleted\n", Cyan(name))
	return true
}

func (warps *Warps) GetEntry(name string) (*string, bool) {
	path, exists := warps.Warps[name]
	if !exists {
		fmt.Printf("%s does not exist\n", Cyan(name))
		return nil, false
	}
	return &path.Path, true
}