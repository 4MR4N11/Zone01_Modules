package main

import (
	"fmt"
	"os"
	"strings"

	utils "ascii_art/utils"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("==>ERROR: IVALID NUMBER OF ARGUMENTS\n==>Usage: go run . [STRING] AND/OR [BANNER]\n\n==>EX: go run . something standard\nEX: go run . something")
		os.Exit(1)
	}
	rawData := []byte{}
	var err error
	if len(os.Args) == 3 {
		if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
			fmt.Println("==>ERROR: IVALID BANNER\n==>Usage: go run . [STRING] [BANNER]\n\n==>EX: go run . something standard")
			os.Exit(1)
		}
		rawData, err = os.ReadFile(os.Args[2] + ".txt")
	} else {
		rawData, err = os.ReadFile("standard.txt")
	}
	if err != nil {
		fmt.Println("Error: Can't open assets file.")
		os.Exit(1)
	}
	asciiTable := utils.AsciiTableMaker(strings.ReplaceAll(string(rawData), "\r", ""))
	arg := strings.ReplaceAll(os.Args[1], "\\n", "\n")
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
	for _, str := range arr {
		if len(str) == 0 {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, c := range str {
					if c >= ' ' && c <= '~' {
						fmt.Print(asciiTable[int(c)-int(' ')][i])
					}
				}
				fmt.Println()
			}
		}
	}
}
