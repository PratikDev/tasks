/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/mergestat/timediff"
	"github.com/pratikdev/tasks/cmdUtils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func runLs(cmd *cobra.Command, args []string) error {
	flag := "working"

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		switch f.Name {
		case "all":
			if f.Changed {
				flag = "all"
			}
		case "pending":
			if f.Changed {
				flag = "pending"
			}
		case "done":
			if f.Changed {
				flag = "done"
			}
		case "cancelled":
			if f.Changed {
				flag = "cancelled"
			}
		}
	})

	tasks, err := (&cmdUtils.Task{}).List(flag)
	if err != nil {
		return cmdUtils.FlagErrorf(err.Error())
	}

	if len(tasks) == 0 {
		fmt.Printf("No tasks found with status %s\n", flag)
		return nil
	}

	writter := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', tabwriter.Debug)
	defer writter.Flush()

	writter.Write([]byte("ID \t Title \t Status \t Created At \n"))

	for _, task := range tasks {
		formattedTime, err := time.Parse(time.RFC3339Nano, task.CreatedAt)
		if err != nil {
			return cmdUtils.FlagErrorf(err.Error())
		}

		tabularData := fmt.Sprintf("%d \t %s \t %s \t %s \n", task.ID, task.Title, task.Status, timediff.TimeDiff(formattedTime))
		writter.Write([]byte(tabularData))
	}

	return nil
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists your tasks",
	Long: heredoc.Docf(`
	ls command lists your tasks.
	
	Example:
	%[1]stasks ls%[1]s
	
	This will list all the tasks that are in Working status.

	Optionally, you can pass %[1]s-a%[1]s, %[1]s-p%[1]s, %[1]s-w%[1]s %[1]s-d%[1]s, or %[1]s-c%[1]s flag.
	%[1]s-a%[1]s flag lists all the tasks.
	%[1]s-p%[1]s flag lists all the tasks that are in Pending status.
	%[1]s-w%[1]s flag lists all the tasks that are in Working status.
	%[1]s-d%[1]s flag lists all the tasks that are in Done status.
	%[1]s-c%[1]s flag lists all the tasks that are in Cancelled status.

	Example with status:
	%[1]stasks ls -a%[1]s

	This will list all the tasks.

	Default flag is Working.
	`, "`"),
	Example: heredoc.Doc(`
	# list tasks that are in Working status
	tasks ls

	# list tasks that are in Pending status
	tasks ls -p

	# list tasks that are in Done status
	tasks ls -d

	# list tasks that are in Cancelled status
	tasks ls -c

	# list all tasks
	tasks ls -a
	`),
	RunE: runLs,
	Args: cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.Flags().BoolP("all", "a", false, "List all tasks")
	lsCmd.Flags().BoolP("pending", "p", false, "List tasks that are in Pending status")
	lsCmd.Flags().BoolP("working", "w", false, "List tasks that are in Working status")
	lsCmd.Flags().BoolP("done", "d", false, "List tasks that are in Done status")
	lsCmd.Flags().BoolP("cancelled", "c", false, "List tasks that are in Cancelled status")
}
