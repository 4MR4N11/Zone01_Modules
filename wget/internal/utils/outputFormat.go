package utils

import (
	"fmt"
	"os"
)

func PrintOutput(args Args, output string) {
	if args.Bdownload {
		logFile, _ := os.Create("wget-log")
		logFile.WriteString(output)
		fmt.Println("Output will be written to \"wget-log\".")
	} else {
		fmt.Print(output)
	}
}
