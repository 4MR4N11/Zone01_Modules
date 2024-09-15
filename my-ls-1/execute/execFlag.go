package execute

import (
	"fmt"
	"os"

	"my_ls/flags"
	Flag "my_ls/flags"
	myStructs "my_ls/models"
	utils "my_ls/utils"
)

func getOutput(paths []myStructs.Path, lowerR bool, l bool) []string {
	output := []string{}
	tmp := []string{}
	line := ""
	if lowerR {
		utils.RevSortPaths(&paths)
	} else {
		utils.SortPaths(&paths)
	}
	for _, p := range paths {
		output = append(output, "")
		files, _ := p.OpenedPath.Readdir(-1)
		if lowerR {
			files = utils.RevSortFiles(files)
		} else {
			files = utils.SortFiles(files)
		}
		for _, file := range files {
			if file.Name()[0] == '.' {
				continue
			}
			tmp = append(tmp, file.Name())
		}
		if len(tmp) == 0 && l {
			output = append(output, "total 0\n")
		}
		for i, file := range tmp {
			if l {
				filePath, _ := os.Open(p.Path)
				output = append(output, flags.LFlag(filePath, p.Path, lowerR)...)
				break
			} else {
				if i+1 < len(tmp) {
					line += fmt.Sprint(file + "  ")
				} else {
					line += fmt.Sprint(file + "\n")
				}
			}
		}
		if len(line) != 0 && !l {
			output = append(output, line)
		}
		tmp = nil
		line = ""
	}
	for _, p := range paths {
		for i, o := range output {
			if len(o) == 0 {
				if i == 0 {
					output[i] = p.Path + ":\n"
				} else {
					output[i] = "\n" + p.Path + ":\n"
				}
				break
			}
		}
	}
	return output
}

func ExecFlags(Input myStructs.Args) {
	output := []string{}
	if Input.Flags.LowerR {
		utils.RevSortPaths(&Input.Path)
	} else {
		utils.SortPaths(&Input.Path)
	}
	for _, path := range Input.Path {
		defer path.OpenedPath.Close()
		if len(Input.Path) > 1 && !Input.Flags.UpperR {
			output = append(output, path.Path+":\n")
		}
		if Input.Flags.UpperR {
			output = append(output, getOutput(Flag.UpperR(path.OpenedPath, path.Path), Input.Flags.LowerR, Input.Flags.L)...)
		}
		if Input.Flags.L {
		}
		// if Input.Flags.L {
		// 	execFlag.LFlag(path.OpenedPath, path.Path, Input.Flags.LowerR)
		// 	if i+1 < len(Input.Path) {
		// 		fmt.Println()
		// 	}
		// } else if Input.Flags.UpperR {
		// 	execFlag.UpperR(path.OpenedPath, path.Path)
		// } else {
		// 	ExecNoFlags(path.OpenedPath, Input.Flags.LowerR)
		// 	if i+1 < len(Input.Path) {
		// 		fmt.Println()
		// 	}
		// }
	}
	for _, l := range output {
		fmt.Print(l)
	}
}
