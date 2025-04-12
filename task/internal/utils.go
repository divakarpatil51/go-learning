package internal

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func WriteToTerminal(tasks *[]Task) {

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	defer w.Flush()
	fmt.Fprintln(w, "Id\tTask Description\tCreated\tStatus")
	for _, task := range *tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.Id, task.Description, task.CreatedAt, task.Status)
	}

	fmt.Fprintf(w, "\nTasks count: %d\n", len(*tasks))
}

func GetStringEnv(key, defaultValue string) string {
	if value, err := os.LookupEnv(key); !err {
		return value
	}
	return defaultValue
}
