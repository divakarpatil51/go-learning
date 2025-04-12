package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/divakarpatil51/go_learning/task/internal"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "A brief description of your command",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("Please provide id for task to be deleted")
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("Invalid id %s", args[0])
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
				(*tasks)[i].Deleted = true
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
		fmt.Printf("Task with id %d deleted successfully\n\n", id)
		listAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
