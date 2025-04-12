package cmd

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/divakarpatil51/go_learning/task/internal"
	"github.com/spf13/cobra"
)

// updateStatusCmd represents the updateStatus command
var updateStatusCmd = &cobra.Command{
	Use:   "update-status [id] [status]",
	Short: "Command to update status of tasks",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return fmt.Errorf("Please provide id and status for task to be updated")
		}

		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("Invalid id %s", args[0])
		}
		allowedStatuses := slices.Collect(maps.Values(internal.StatusName))
		if !slices.Contains(allowedStatuses, args[1]) {
			return fmt.Errorf("Invalid status. Allowed values: %s", strings.Join(allowedStatuses, ", "))
		}

		return nil
	},
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

		id, _ := strconv.Atoi(args[0])
		var found bool
		for i, task := range *tasks {
			if id == task.Id && !task.Deleted {
				(*tasks)[i].Status = args[1]
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Task with id %d not found\n\n", id)
			listAllTasks()
			return
		}
		writeErr := internal.WriteCsvFile(file, *tasks)
		if writeErr != nil {
			fmt.Printf("Error writing to file: %v", writeErr)
			os.Exit(1)
		}
		fmt.Printf("Task with id %d updated successfully\n\n", id)
		listAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(updateStatusCmd)
}
