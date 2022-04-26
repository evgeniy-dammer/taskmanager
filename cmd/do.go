package cmd

import (
	"fmt"
	"strconv"

	"github.com/evgeniy-dammer/taskmanager/db"
	"github.com/spf13/cobra"
)

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)

			if err != nil {
				fmt.Println("Faild to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return
		}

		for _, val := range ids {
			if val <= 0 || val > len(tasks) {
				fmt.Println("Invalid task number:", val)

				continue
			}

			task := tasks[val-1]

			err := db.DeleteTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", val, err)
			} else {
				fmt.Printf("Task \"%d\" marked as completed\n", val)
			}
		}
	},
}

func init() {
	RootCommand.AddCommand(doCommand)
}
