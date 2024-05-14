package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Invalid arguments number.")
		os.Exit(1)
	}
	rawData, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: Can't open assets file.")
		os.Exit(1)
	}
	arg := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	dataTmp := strings.Split(string(rawData), "\n")
	asciiTable := [][]string{}
	tmp := []string{}
	for _, l := range dataTmp {
		if len(l) > 0 {
			tmp = append(tmp, l)
		} else {
			if len(tmp) > 0 {
				asciiTable = append(asciiTable, tmp)
				tmp = []string{}
			}
		}
	}
	str := strings.Split(arg, "\n")
	if len(arg) == strings.Count(arg, "\n") {
		str = str[:len(str)-1]
	}
	for _, s := range str {
		if len(s) == 0 {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, c := range s {
					if c >= ' ' && c <= '~' {
						fmt.Print(asciiTable[int(c)-int(' ')][i])
					}
				}
				fmt.Println()
			}
		}
	}
}
