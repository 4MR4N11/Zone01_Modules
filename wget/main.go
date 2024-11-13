package main

import (
	"os"

	cmd "wget/cmd"
	downloader "wget/internal/downloader"
	utils "wget/internal/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		utils.PrintError()
		return
	}
	parsedInput := cmd.ParseArgs(args)
	downloader.HttpDownloader(parsedInput)
}
