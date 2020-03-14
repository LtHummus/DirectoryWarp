package main

import (
	"DirectoryWarp/warps"
	"flag"
	"fmt"
	"os"
)

func main() {
	//setup command line parsing
	database, err := warps.Load("test.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%+v\n", database)
	os.Exit(0)

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
		fmt.Printf("list\n")
		fmt.Printf("%v\n", listCmd.Args())

	case "set":
		err := setCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error parsing commands: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("set\n")
	}

}