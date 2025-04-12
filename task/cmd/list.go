package cmd

import (
	"fmt"
	"os"

	"github.com/divakarpatil51/go_learning/task/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks. By default, lists only todo tasks. Filters can be applied for filtering by status",
	Run: func(cmd *cobra.Command, args []string) {

		value := internal.GetStringEnv("TASKS_FILE_PATH", "tasks.csv")
		file, err := internal.LoadFile(value)
		if err != nil {
			fmt.Printf("Error loading file: %v", err)
			os.Exit(1)
		}
		defer file.Close()

		tasks, readErr := internal.ReadCsvFile(file)
		if readErr != nil {
			fmt.Printf("Error reading file: %v", readErr)
			os.Exit(1)
		}

		flags := cmd.Flags()

		all := flags.Changed("all")
		inProgressFilter := flags.Changed("in-progress") || all
		blockedFilter := flags.Changed("blocked") || all
		doneFilter := flags.Changed("done") || all
		todoFilter := flags.Changed("todo") || !(inProgressFilter || blockedFilter || doneFilter) || all

		var filteredTasks []internal.Task

		statusMap := internal.StatusName
		for _, task := range *tasks {
			if !task.Deleted &&
				((task.Status == statusMap[internal.TODO] && todoFilter) ||
					(task.Status == statusMap[internal.IN_PROGRESS] && inProgressFilter) ||
					(task.Status == statusMap[internal.BLOCKED] && blockedFilter) ||
					(task.Status == statusMap[internal.DONE] && doneFilter)) {
				filteredTasks = append(filteredTasks, task)
			}
		}
		internal.WriteToTerminal(&filteredTasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("todo", "t", false, "List only Todo tasks")
	listCmd.Flags().BoolP("in-progress", "i", false, "List only In Progress tasks")
	listCmd.Flags().BoolP("blocked", "b", false, "List only blocked tasks")
	listCmd.Flags().BoolP("done", "d", false, "List only done tasks")
	listCmd.Flags().BoolP("all", "a", false, "List all tasks")
}
