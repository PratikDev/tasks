/*
Copyright © 2024 Pratik Dev pratikdevofficial1@gmail.com
This file is part of the tasks project.
*/
package cmd

import (
	"os"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A command line interface to manage your tasks",
	Long: heredoc.Docf(`
	Tasks is a command line tool to manage your tasks.
	You can add, list, edit, and remove your tasks using this tool.
	`),
	Example: heredoc.Doc(`
	# Add a new task
	tasks add "Buy groceries"

	# List all "Working" status tasks
	tasks list

	# List all tasks
	tasks list -a

	# Edit the task title with id 1
	tasks edit 1 -t "Buy groceries and vegetables"

	# Remove the task with id 1
	tasks rm 1

	# Mark the task with id 1 as Done
	tasks edit 1 -d

	# Mark the task with id 1 as Pending
	tasks edit 1 -p

	# Remove all Cancelled tasks
	tasks prune

	# Remove all Done tasks
	tasks prune -d
	`),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
