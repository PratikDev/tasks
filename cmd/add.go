/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/pratikdev/tasks/cmdUtils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func runAdd(cmd *cobra.Command, args []string) error {
	title := args[0]
	statusFlag := "Pending"

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if f.Changed {
			statusFlag = strings.ToUpper(string(f.Name[0])) + string(f.Name[1:])
		}
	})

	err := (&cmdUtils.Task{
		Title:  title,
		Status: statusFlag,
	}).New()
	if err != nil {
		return cmdUtils.FlagErrorf(err.Error())
	}

	fmt.Printf("Task \"%s\" added successfully\n", title)
	return nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [<title>]",
	Short: "Adds a new task to your task list",
	Long: heredoc.Docf(`
	%[1]sadd%[1]s command is used to add a new task to your task list.
	
	Example:
	%[1]stasks add "Learn Go"%[1]s
	
	This will add a new task "Learn Go" to your task list.
	
	Optionally, you can pass %[1]s-p%[1]s, %[1]s-w%[1]s, %[1]s-d%[1]s, or %[1]s-c%[1]s flag.
	%[1]s-p%[1]s flag sets the status of the task to Pending.
	%[1]s-w%[1]s flag sets the status of the task to Working.
	%[1]s-d%[1]s flag sets the status of the task to Done.
	%[1]s-c%[1]s flag sets the status of the task to Cancelled.
	
	Example with status:
	%[1]stasks add -t "Learn Go" -w%[1]s
	
	This will add a new task "Learn Go" to your task list with status Working.
	
	Default status is Pending.
	`, "`"),
	Example: heredoc.Doc(`
			# add a new task
			tasks add "Learn Go"

			# add a new task with status Pending (default)
			tasks add "Learn Go" -p

			# add a new task with status Working
			tasks add "Learn Go" -w

			# add a new task with status Done
			tasks add "Learn Go" -d

			# add a new task with status Cancelled
			tasks add "Learn Go" -c
		`),
	Args: cobra.ExactArgs(1),
	RunE: runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolP("pending", "p", false, "Sets status to Pending")
	addCmd.Flags().BoolP("working", "w", false, "Sets status to Working")
	addCmd.Flags().BoolP("done", "d", false, "Sets status to Done")
	addCmd.Flags().BoolP("cancelled", "c", false, "Sets status to Cancelled")
}
