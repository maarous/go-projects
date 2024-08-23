package cmd

import (
	"encoding/csv"
	"log"
	"os"
)

// ReadCSV reads a CSV file and returns its contents.
//
// Parameters:
//
//	fileName: The name of the CSV file to be read.
//
// Returns:
//
//	A 2D slice of strings representing the CSV records.
func ReadCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CSV file opened successfully on read only mode")
	defer file.Close()

	fileReader := csv.NewReader(file)
	records, err := fileReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

// WriteToCSV writes a record to a CSV file.
//
// Parameters:
//
//	fileName: The name of the CSV file.
//	records: A 2D slice of strings representing the record(s) to be written.
//
// Returns:
//
//	None
func WriteToCSV(fileName string, records [][]string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	log.Println("CSV file opened successfully on read/write mode")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.WriteAll(records); err != nil {
		log.Fatal(err)
	}

	log.Println("CSV file updated successfully!")
}

// UpdateCSV updates a CSV file by replacing its contents with the provided contents.
//
// Parameters:
//
//	fileName: The name of the CSV file to be updated.
//	updatedRecords: A 2D slice of strings representing the new record(s) to be written.
//
// Returns:
//
//	None
func UpdateCSV(fileName string, updatedRecords [][]string) {

	tempFileName := FileName + ".tmp"
	WriteToCSV(tempFileName, updatedRecords)

	err := os.Remove(FileName)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename(tempFileName, FileName)
	if err != nil {
		log.Fatal(err)
	}
}
