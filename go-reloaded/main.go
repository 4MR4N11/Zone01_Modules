package main

import (
	"os"

	utils "reloaded/utils"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		os.Exit(1)
	}
	file, err := os.Create(outputFile)
	if err != nil {
		os.Exit(1)
	}
	newData := utils.SearchAndReplaceOp(string(data))
	newData = utils.PunctHandler(newData)
	newData = utils.Quotehandler(newData)
	_, err = file.WriteString(newData)
	if err != nil {
		os.Exit(1)
	}
}
