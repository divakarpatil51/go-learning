package cmd

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/divakarpatil51/go_learning/task/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task to your TODO list",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("Please provide description for todo ")
		}

		status := cmd.Flag("status").Value.String()

		allowedStatuses := slices.Collect(maps.Values(internal.StatusName))
		if !slices.Contains(allowedStatuses, status) {
			return fmt.Errorf("Invalid status. Allowed values: %s", strings.Join(allowedStatuses, ", "))
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		value := internal.GetStringEnv("TASKS_FILE_PATH", "tasks.csv")
		file, err := internal.LoadFile(value)
		if err != nil {
			return fmt.Errorf("Error loading file: %v", err)
		}
		defer file.Close()

		var tasks *[]internal.Task
		tasks, readErr := internal.ReadCsvFile(file)
		if readErr != nil {
			return fmt.Errorf("Error reading file: %v", readErr)
		}
		nextTaskId := len(*tasks) + 1

		task := internal.Task{
			Id:          nextTaskId,
			Description: args[0],
			CreatedAt:   time.Now().Format(time.DateTime),
			Status:      cmd.Flag("status").Value.String(),
		}

		writeErr := internal.AppendToCsvFile(file, task.ToCSVFormat())
		if writeErr != nil {
			return fmt.Errorf("Error writing to file: %v", writeErr)
		}

		listAllTasks()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	statuses := strings.Join(slices.Collect(maps.Values(internal.StatusName)), ", ")
	addCmd.Flags().StringP("status", "s", internal.StatusName[internal.TODO], "Task status. Allowed values: "+statuses)
}
