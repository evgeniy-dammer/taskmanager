package cmd

import "github.com/spf13/cobra"

var RootCommand = &cobra.Command{
	Use:   "taskmanager",
	Short: "CLI Task Manager",
}
