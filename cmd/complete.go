/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed using its id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id := args[0]
		records := ReadCSV(FileName)
		found := false

		for i, record := range records {
			if record[0] == id {
				records[i][2] = strconv.FormatBool(true)
				found = true
				break
			}
		}

		if found {
			UpdateCSV(FileName, records)
			log.Println("Task marked as completed successfully!")
		} else {
			log.Printf("Error: ID %s not found in the CSV file.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
