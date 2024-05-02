package main

import (
	"os"
	"strings"

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
			os.Exit(1)
		}
	}
}
