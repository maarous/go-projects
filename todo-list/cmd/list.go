package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in your TODO list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		records := ReadCSV(FileName)
		// Create a new tabwriter.Writer with specified settings
		writer := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)

		// Write formatted data to the writer
		fmt.Fprintln(writer, "ID\tTask\tCompleted\tCreatedAt")
		for _, record := range records {
			fmt.Fprintln(writer, record[0], "\t", record[1], "\t", record[2], "\t", record[3])
		}
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
