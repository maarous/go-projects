package cmd

import (
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var Task string

// init Initializes the root command with the add command.
//
// No parameters.
// No return value.
func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		task := args[0]
		id := generateID(task)
		record := []string{strconv.Itoa(id), task, strconv.FormatBool(false), time.Now().Format(time.RFC3339)}

		WriteToCSV(FileName, [][]string{record})

		log.Println("Task added successfully!")
	},
}

// generateID generates a new ID for a task by reading the last ID from the TODO file and incrementing it by 1.
//
// No parameters.
// Returns an integer representing the new ID.
func generateID(task string) int {
	records := ReadCSV(FileName)

	lastID := 0
	if len(records) > 0 {
		lastID, _ = strconv.Atoi(records[len(records)-1][0])
	}

	log.Println("New ID was generated successfully for the task:", task)

	return lastID + 1
}
