package main

import (
	"DirectoryWarp/warps"
	"flag"
	"fmt"
	"os"
)

const DATABASE_PATH = "test.json"

func main() {
	//setup command line parsing
	database, err := warps.Load(DATABASE_PATH)
	if err != nil {
		fmt.Printf("Error loading database: %v\n", err)
	}

	//list
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	//set
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Error. No command given")
		os.Exit(1)
	}

	switch os.Args[1]{
	case "list":
		err := listCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error parsing commands: %v\n", err)
			os.Exit(1)
		}
		listEntries(database)

	case "set":
		err := setCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error parsing commands: %v\n", err)
			os.Exit(1)
		}
		params := setCmd.Args()
		if len(params) < 2 {
			fmt.Printf("Name and path not specified")
			os.Exit(1)
		}
		setEntry(database, params[0], params[1])
		err = database.Write(DATABASE_PATH)
		if err != nil {
			fmt.Printf("Error writing database back out: %v\n", err)
			os.Exit(1)
		}
	}

}

