package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/evgeniy-dammer/taskmanager/cmd"
	"github.com/evgeniy-dammer/taskmanager/db"
	"github.com/mitchellh/go-homedir"
)

func ifErrorExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	homeDir, err := homedir.Dir()

	if err != nil {
		panic(err)
	}

	dbPath := filepath.Join(homeDir, "tasks.db")

	ifErrorExit(db.InitDB(dbPath))

	ifErrorExit(cmd.RootCommand.Execute())
}
