/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/pratikdev/tasks/cmdUtils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func runPrune(cmd *cobra.Command, args []string) error {
	statusFlag := "cancelled"

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if f.Changed {
			statusFlag = f.Name
		}
	})

	err := (&cmdUtils.Task{}).Prune(statusFlag)
	if err != nil {
		return cmdUtils.FlagErrorf(err.Error())
	}

	return nil
}

// pruneCmd represents the prune command
var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Cleans your tasks list",
	Long:  `prune command cleans your tasks list. By default, it removes all the cancelled tasks. Optionally you can pass different flags clean tasks of other status`,
	Example: heredoc.Docf(`
	# Removes all the tasks with Cancelled status
	tasks prune

	# Removes all the tasks with Pending status
	tasks prune -p

	# Removes all the tasks with Working status
	tasks prune -w

	# Removes all the tasks with Done status
	tasks prune -d
	`),
	Args: cobra.NoArgs,
	RunE: runPrune,
}

func init() {
	rootCmd.AddCommand(pruneCmd)

	pruneCmd.Flags().BoolP("pending", "p", false, "Removes all the tasks with Pending status")
	pruneCmd.Flags().BoolP("working", "w", false, "Removes all the tasks with Working status")
	pruneCmd.Flags().BoolP("done", "d", false, "Removes all the tasks with Done status")
	pruneCmd.Flags().BoolP("cancelled", "c", false, "Removes all the tasks with Cancelled status")
}
