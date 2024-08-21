/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"main.go/cmd"
)

// type item struct {
// 	ID        int
// 	Task      string
// 	Done      bool
// 	CreatedAt time.Time
// }

// type List []item

const FileName = "TODO.csv"

func main() {
	cmd.FileName = FileName
	cmd.Execute()
}
