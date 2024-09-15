package main

import (
	"fmt"
	"os"
	"strings"

	parser "my_ls/Parser"
	execute "my_ls/execute"
	myStructs "my_ls/models"
	utils "my_ls/utils"
)

func main() {
	args := os.Args[1:]
	Input := myStructs.Args{}
	path := myStructs.Path{}
	Input.Flags = utils.InitStruct()
	if len(args) != 0 {
		Input = parser.GetArgs(args)
	} else {
		var err error
		path.Path = "."
		path.OpenedPath, err = os.Open(".")
		if err != nil {
			str := fmt.Sprintln(err)
			fmt.Fprint(os.Stderr, strings.Replace(str, "open "+".", "my_ls: cannot access '"+"."+"'", 1))
			os.Exit(2)
		}
		Input.Path = append(Input.Path, path)
		utils.SortPaths(&Input.Path)
		fmt.Print(execute.ExecNoFlags(path.OpenedPath))
		return
	}
	execute.ExecFlags(Input)
}
