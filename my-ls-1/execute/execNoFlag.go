package execute

import (
	"fmt"
	"os"

	utils "my_ls/utils"
)

func ExecNoFlags(openedFile *os.File) string {
	results := []string{}
	output := ""
	files, err := openedFile.Readdir(-1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	files = utils.SortFiles(files)
	for _, file := range files {
		if file.Name()[0] == '.' {
			continue
		}
		results = append(results, file.Name())
	}
	for i, file := range results {
		if i+1 < len(results) {
			output += fmt.Sprint(file + "  ")
		} else {
			output += fmt.Sprint(file + "\n")
		}
	}
	return output
}
