package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"DirectoryWarp/warps"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

const DefaultDatabaseName = ".warps.json"
const DatabaseEnvironmentVariableName = "WARPS_DATABASE_FILE"

func loadDatabase(databasePath string) (*warps.Warps, error) {
	return warps.Load(databasePath)
}

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

func setWarp(c *cli.Context) error {
	if c.Args().Len() < 2 {
		return cli.Exit("setting a warp requires two arguments", 1)
	}

	databasePath := getWarpsPath()
	database, err := loadDatabase(databasePath)
	if err != nil {
		return err
	}

	database.SetEntry(c.Args().Get(0), c.Args().Get(1))
	err = database.Write(databasePath)
	return err
}

func listWarps(c *cli.Context) error {
	databasePath := getWarpsPath()
	database, err := loadDatabase(databasePath)
	if err != nil {
		return err
	}

	database.ListEntries()
	return nil
}

func jumpWarp(c *cli.Context) error {
	if c.Args().Len() < 1 {
		return cli.Exit("need a warp to go to", 1)
	}

	databasePath := getWarpsPath()
	database, err := loadDatabase(databasePath)
	if err != nil {
		return err
	}

	path, exists := database.GetEntry(c.Args().First())
	if !exists {
		return cli.Exit("", 1)
	} else {
		//have to do the print here since the API will send to stderr
		fmt.Printf("%s\n", *path)
		return cli.Exit("", 2)
	}
}

func deleteWarp(c *cli.Context) error {
	if c.Args().Len() < 1 {
		return cli.Exit("need a warp to delete", 1)
	}

	databasePath := getWarpsPath()
	database, err := loadDatabase(databasePath)
	if err != nil {
		return err
	}

	database.DeleteEntry(c.Args().First())
	return database.Write(databasePath)
}

func main() {
	app := &cli.App{
		Name:        "wd",
		Version:     "0.0.1",
		Description: "it's like `cd` but with bookmarks",
		Usage:       "a way to warp directories",
		Commands: []*cli.Command{
			{
				Name:      "set",
				Aliases:   []string{"add", "create", "s", "a"},
				ArgsUsage: "name path",
				Usage:     "add a warp point",
				Action:    setWarp,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all warp points",
				Action:  listWarps,
			},
			{
				Name:    "go",
				Aliases: []string{"g", "j", "jump", "warp", "w"},
				Usage:   "jump to warp point",
				Action:  jumpWarp,
			},
			{
				Name:   "delete",
				Usage:  "delete warp point",
				Action: deleteWarp,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
