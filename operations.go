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
