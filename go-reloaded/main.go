package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	utils "reloaded/utils"
)

func punctHandling(data string) string {
	re := regexp.MustCompile(`[.,;:!?]\S`)
	regMatches := re.FindAllString(data, -1)
	fmt.Println(regMatches)
	for _, m := range regMatches {
		index := strings.Index(data, m)
		if index != -1 {
			if index+1 < len(data) {
				data = data[:index+1] + " " + data[index+1:]
			}
		}
	}
	fmt.Println(data)
	re = regexp.MustCompile(`\s+[.,;:!?]`)
	regMatches = re.FindAllString(data, -1)
	for _, m := range regMatches {
		index := strings.Index(data, m)
		if index != -1 {
			data = data[:index] + data[len(m)+index-1:]
		}
	}
	return data
}

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
	newData = punctHandling(newData)
	_, err = file.WriteString(newData)
	if err != nil {
		os.Exit(1)
	}
}
