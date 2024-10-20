/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/pratikdev/tasks/cmdUtils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func runEdit(cmd *cobra.Command, args []string) error {
	id := args[0]

	var flag string = ""
	var title string = ""

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		switch f.Name {
		case "title":
			title = f.Value.String()
		default:
			if f.Changed {
				flag = strings.ToUpper(string(f.Name[0])) + string(f.Name[1:])
			}
		}
	})

	if flag == "" && title == "" {
		return cmdUtils.FlagErrorf("Title or flag is required")
	}

	err := (&cmdUtils.Task{}).Edit(id, title, flag)
	if err != nil {
		return cmdUtils.FlagErrorf(err.Error())
	}

	return nil
}

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [<task-id>]",
	Short: "Edit a task from the task list",
	Long: heredoc.Docf(`
	rm command edits a task from the task list. You can edit both the title and the status of a task by it's id.

	rm command do not have any Default flags
	`),
	Example: heredoc.Docf(`
	# edit the title of task with id 5
	tasks edit 5 -t "New title"

	# mark the task with id 5 as Pending
	tasks edit 5 -p

	# mark the task with id 5 as Working
	tasks edit 5 -w

	# mark the task with id 5 as Done
	tasks edit 5 -d

	# mark the task with id 5 as Cancelled
	tasks edit 5 -c

	# edit the title and mark the task with id 5 as Working
	tasks edit 5 -t "New title" -w
	`),
	Args: cobra.ExactArgs(1),
	RunE: runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("title", "t", "", "New title for the task")
	editCmd.Flags().BoolP("pending", "p", false, "Mark the task as Pending")
	editCmd.Flags().BoolP("working", "w", false, "Mark the task as Working")
	editCmd.Flags().BoolP("done", "d", false, "Mark the task as Done")
	editCmd.Flags().BoolP("cancelled", "c", false, "Mark the task as Cancelled")
}
