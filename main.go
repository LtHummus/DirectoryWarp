package main

import (
	"DirectoryWarp/warps"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const DefaultDatabaseName = ".warps.json"
const DatabaseEnvironmentVariableName = "WARPS_DATABASE_FILE"

func getWarpsPath() string {
	customPath, useCustomPath := os.LookupEnv(DatabaseEnvironmentVariableName)
	if useCustomPath {
		return customPath
	}

	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Error loading directory: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, DefaultDatabaseName)
}

func main() {
	//setup command line parsing
	databasePath := getWarpsPath()
	database, err := warps.Load(databasePath)
	if err != nil {
		fmt.Printf("Error loading database: %v\n", err)
	}

	//go
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)

	//list
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	//set
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)

	//delete
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Error. No command given")
		os.Exit(1)
	}

	switch os.Args[1]{
	case "go":
		err := goCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error parsing commands: %v\n", err)
			os.Exit(1)
		}
		params := goCmd.Args()
		if len(params) < 1 {
			fmt.Printf("No name specified\n")
			os.Exit(1)
		}
		path, exists := database.GetEntry(params[0])
		if !exists {
			os.Exit(1)
		} else {
			fmt.Printf("%s\n", *path)
			os.Exit(2)
		}

	case "list":
		err := listCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error parsing commands: %v\n", err)
			os.Exit(1)
		}
		database.ListEntries()

	case "delete":
		err := deleteCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error parsing commands: %v\n", err)
			os.Exit(1)
		}
		params := deleteCmd.Args()
		if len(params) < 1 {
			fmt.Printf("Name not specified\n")
			os.Exit(1)
		}
		database.DeleteEntry(params[0])
		err = database.Write(databasePath)
		if err != nil {
			fmt.Printf("Error writing database back out: %v\n", err)
			os.Exit(1)
		}

	case "set":
		err := setCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error parsing commands: %v\n", err)
			os.Exit(1)
		}
		params := setCmd.Args()
		if len(params) < 2 {
			fmt.Printf("Name and path not specified\n")
			os.Exit(1)
		}
		database.SetEntry(params[0], params[1])
		err = database.Write(databasePath)
		if err != nil {
			fmt.Printf("Error writing database back out: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("Command not recognized\n")
		os.Exit(1)
	}

}

