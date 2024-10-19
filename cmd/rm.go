/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/pratikdev/tasks/cmdUtils"
	"github.com/spf13/cobra"
)

func runRm(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmdUtils.FlagErrorf("task id is required")
	}

	id := args[0]
	err := (&cmdUtils.Task{}).Remove(id)
	if err != nil {
		return cmdUtils.FlagErrorf(err.Error())
	}

	fmt.Println("Task removed successfully")
	return nil
}

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [<task-id>]",
	Short: "Removes a task from your task list",
	Long:  `rm command is used to remove a task from your task list.`,
	Example: heredoc.Doc(`
	# Remove the task with ID 1
	tasks rm 1
	`),
	Args: cobra.MaximumNArgs(1),
	RunE: runRm,
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
