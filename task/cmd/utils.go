package cmd

import (
	"fmt"
)

func listAllTasks() error {
	rootCmd.SetArgs([]string{"list", "-a"})
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("Error listing tasks: %v", err)
	}
	return nil
}
