package parser

import (
	"fmt"
	"os"
	"strings"

	models "my_ls/models"
	utils "my_ls/utils"
)

func getFlags(flag string, flags *models.Flags) {
	for _, f := range flag {
		if f == 'l' {
			flags.L = true
		} else if f == 'a' {
			flags.A = true
		} else if f == 'R' {
			flags.UpperR = true
		} else if f == 'r' {
			flags.LowerR = true
		} else if f == 't' {
			flags.T = true
		} else {
			fmt.Fprintf(os.Stderr, "my_ls: option requires an argument -- '%c'\n", f)
			os.Exit(2)
		}
	}
}

func GetArgs(args []string) models.Args {
	result := models.Args{}
	flags := models.Flags{}
	paths := []models.Path{}
	check := 0
	flags = utils.InitStruct()
	for _, arg := range args {
		if len(arg) > 0 && arg[0] == '-' {
			getFlags(arg[1:], &flags)
		} else {
			tmp, err := os.Open(arg)
			if err != nil {
				str := fmt.Sprintln(err)
				fmt.Fprint(os.Stderr, strings.Replace(str, "open "+arg, "my_ls: cannot access '"+arg+"'", 1))
				check = 1
			} else {
				p := models.Path{}
				p.OpenedPath = tmp
				p.Path = arg
				paths = append(paths, p)
			}
		}
	}
	if len(paths) == 0 && check == 0 {
		p := models.Path{}
		tmp, err := os.Open(".")
		if err != nil {
			str := fmt.Sprintln(err)
			fmt.Fprint(os.Stderr, strings.Replace(str, "open "+".", "my_ls: cannot access '"+"."+"'", 1))
			os.Exit(2)
		}
		p.OpenedPath = tmp
		p.Path = "."
		paths = append(paths, p)
	}
	result.Path = paths
	utils.SortPaths(&result.Path)
	result.Flags = flags
	return result
}
