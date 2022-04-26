package cmd

import (
	"fmt"

	"github.com/evgeniy-dammer/taskmanager/db"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete.")
			return
		}

		fmt.Println("You have following tasks: ")

		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCommand.AddCommand(listCommand)
}
