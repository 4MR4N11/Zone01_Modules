package main

import (
	"fmt"
	"os"
	"strings"

	parser "my_ls/Parser"
	execFlag "my_ls/flags"
	myStructs "my_ls/models"
	utils "my_ls/utils"
)

func ExecFlags(Input myStructs.Args) {
	for i, path := range Input.Path {
		defer path.OpenedPath.Close()
		if len(Input.Path) > 1 {
			fmt.Println(path.Path + ":")
		}
		if Input.Flags.L {
			execFlag.LFlag(path.OpenedPath, path.Path, Input.Flags.LowerR)
			if i+1 < len(Input.Path) {
				fmt.Println()
			}
		} else if Input.Flags.UpperR {
		} else {
			ExecNoFlags(path.OpenedPath, Input.Flags.LowerR)
			if i+1 < len(Input.Path) {
				fmt.Println()
			}
		}
	}
}

func ExecNoFlags(openedFile *os.File, lowerR bool) {
	results := []string{}
	files, err := openedFile.Readdir(-1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if lowerR {
		files = utils.RevSortFiles(files)
	} else {
		files = utils.SortFiles(files)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		if file.Name()[0] == '.' {
			continue
		}
		results = append(results, file.Name())
	}
	for i, file := range results {
		if i+1 < len(results) {
			fmt.Print(file + "  ")
		} else {
			fmt.Print(file + "\n")
		}
	}
}

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
		ExecNoFlags(path.OpenedPath, Input.Flags.LowerR)
		return
	}
	if Input.Flags.LowerR {
		utils.RevSortPaths(&Input.Path)
	} else {
		utils.SortPaths(&Input.Path)
	}
	ExecFlags(Input)
}
