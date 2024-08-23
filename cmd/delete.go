package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing task using its id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id := args[0]
		records := ReadCSV(FileName)
		found := false
		var updatedRecords [][]string

		for _, record := range records {
			if record[0] != id {
				updatedRecords = append(updatedRecords, record)
			} else {
				found = true
				break
			}
		}

		if found {
			UpdateCSV(FileName, updatedRecords)
			log.Println("Task deleted successfully!")
		} else {
			log.Printf("Error: ID %s not found in the CSV file.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
