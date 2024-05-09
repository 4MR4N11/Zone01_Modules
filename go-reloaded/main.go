package main

import (
	"fmt"
	"os"
	"strings"

	utils "reloaded/utils"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: Arguments number is incorrect")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: File not found or could not be opened.")
		os.Exit(1)
	}
	file, err := os.Create(outputFile)
	if err != nil || !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Error: File could not be created or file is not a .txt file.")
		os.Exit(1)
	}
	tmpData := strings.Split(string(data), "\n")
	for i, line := range tmpData {
		newData := utils.SearchAndReplaceOp(string(line))
		newData = utils.VowelHandler(newData)
		newData = utils.PunctHandler(newData)
		newData = utils.Quotehandler(newData)
		if i != len(tmpData)-1 {
			newData += "\n"
		}
		_, err = file.WriteString(newData)
		if err != nil {
			fmt.Println("Error: Could not write to file.")
			os.Exit(1)
		}
	}
}
