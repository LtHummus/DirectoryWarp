package main

import (
	"DirectoryWarp/warps"
	"fmt"
	"path/filepath"

	. "github.com/logrusorgru/aurora"
)


func listEntries(warps *warps.Warps) {
	for name, entry := range warps.Warps {
		fmt.Printf("%s -> %s\n", Cyan(name), Green(entry.Path))
	}
}

func setEntry(warps *warps.Warps, name string, path string) {
	cleanedPath, err := filepath.Abs(path)
	if err != nil {
		msg := fmt.Sprintf("Can not clean path: %v", err)
		fmt.Printf("%s\n", Red(msg))
	}

	overwritten := warps.Add(name, cleanedPath)
	if overwritten {
		fmt.Printf("%s\n", Yellow("Overwriting old entry"))
	}

	fmt.Printf("Set %s to %s\n", Cyan(name), Green(cleanedPath))
}

func deleteEntry(warps *warps.Warps, name string) bool {
	_, exists := warps.Warps[name]
	if !exists {
		fmt.Printf("Entry %s does not exist\n", Cyan(name))
		return false
	}

	delete(warps.Warps, name)
	fmt.Printf("Entry %s deleted\n", Cyan(name))
	return true
}

func getEntry(warps *warps.Warps, name string) (*string, bool) {
	path, exists := warps.Warps[name]
	if !exists {
		fmt.Printf("%s does not exist\n", Cyan(name))
		return nil, false
	}
	return &path.Path, true
}