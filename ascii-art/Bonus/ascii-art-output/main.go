package main

import (
	"fmt"
	"os"
	"strings"

	utils "ascii_art/utils"
)

func main() {
	args := os.Args[1:]
	outputFileName := ""
	inputString := ""
	banner := ""

	if len(args) == 3 {
		if !strings.HasPrefix(args[0], "--output=") || !strings.HasSuffix(args[0], ".txt") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(1)
		}
		outputFileName = strings.TrimPrefix(args[0], "--output=")
		inputString = args[1]
		banner = args[2]
	} else if len(args) == 2 {
		inputString = args[0]
		banner = args[1]
	} else if len(args) == 1 {
		inputString = args[0]
		banner = "standard"
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}
	asciiFile, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error: Can't open assets file.")
		os.Exit(1)
	}
	asciiCharacters := strings.ReplaceAll(string(asciiFile), "\r", "")
	asciiTable := utils.AsciiTableMaker(asciiCharacters)
	arg := strings.ReplaceAll(inputString, "\\n", "\n")
	arr := strings.Split(arg, "\n")
	for _, str := range arr {
		for _, c := range str {
			if c < ' ' || c > '~' {
				fmt.Println("Error: Character out of range.")
				os.Exit(1)
			}
		}
	}

	if len(arg) == strings.Count(arg, "\n") {
		arr = arr[:len(arr)-1]
	}
	output := ""
	for _, str := range arr {
		if len(str) == 0 {
			output += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for _, c := range str {
					if c >= ' ' && c <= '~' {
						output += asciiTable[int(c)-int(' ')][i]
					}
				}
				output += "\n"
			}
		}
	}
	if outputFileName != "" {
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println("Error: Can't create output file.")
			os.Exit(1)
		}
		_, err = outputFile.WriteString(output)
		if err != nil {
			fmt.Println("Error: Can't write in file.")
			os.Exit(1)
		}
	} else {
		fmt.Print(output)
	}
}
