package parser

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	myStructs "my_ls/models"
)

func InitStruct() myStructs.Flags {
	flags := myStructs.Flags{}
	flags.A = false
	flags.L = false
	flags.LowerR = false
	flags.T = false
	flags.UpperR = false
	return flags
}

func CountDir(dir fs.FileInfo) int {
	path := "." + "/" + dir.Name()
	content, _ := os.ReadDir(path)
	count := 0
	for _, c := range content {
		if c.IsDir() {
			count++
		}
	}
	return count
}

func GetTotalSize(files []fs.FileInfo) int {
	totalSize := 0
	for _, f := range files {
		tmp := 4
		if f.Size() > 4096 {
			tmp = int(f.Size()) / 1024
			if int(f.Size())%1024 != 0 {
				tmp += 4
			}
		}
		totalSize += tmp
	}
	return totalSize
}

func GetSymLink(file fs.FileInfo, path string) (string, fs.FileInfo) {
	fullPath := path + "/" + file.Name()
	info, err := os.Lstat(fullPath)
	if err != nil {
		os.Exit(1)
	}
	symLink := ""
	if info.Mode()&os.ModeSymlink != 0 {
		var err error
		symLink, err = os.Readlink(fullPath)
		if err != nil {
			return "", nil
		}
	}
	return symLink, info
}

func findMaxLen(resultsLen []myStructs.ResultLen, field string) int {
	max := 0
	for _, res := range resultsLen {
		switch field {
		case "Perms":
			if res.Perms > max {
				max = res.Perms
			}
		case "Num":
			if res.Num > max {
				max = res.Num
			}
		case "UserOwn":
			if res.UserOwn > max {
				max = res.UserOwn
			}
		case "GroupOwn":
			if res.GroupOwn > max {
				max = res.GroupOwn
			}
		case "Size":
			if res.Size > max {
				max = res.Size
			}
		case "Date":
			if res.Date > max {
				max = res.Date
			}
		}
	}
	return max
}

func addSpaces(str string, num int) string {
	for i := 0; i < num; i++ {
		str = " " + str
	}
	return str
}

func PrintResults(results []myStructs.Result, resultsLen []myStructs.ResultLen) {
	permsMax := findMaxLen(resultsLen, "Perms")
	numMax := findMaxLen(resultsLen, "Num")
	userOwnMax := findMaxLen(resultsLen, "UserOwn")
	GroupOwnMax := findMaxLen(resultsLen, "GroupOwn")
	sizeMax := findMaxLen(resultsLen, "Size")
	dateMax := findMaxLen(resultsLen, "Date")
	for i, result := range results {
		if resultsLen[i].Perms < permsMax {
			result.Perms = addSpaces(result.Perms, permsMax-resultsLen[i].Perms)
		}
		if resultsLen[i].Num < numMax {
			result.Num = addSpaces(result.Num, numMax-resultsLen[i].Num)
		}
		if resultsLen[i].UserOwn < userOwnMax {
			result.UserOwn = addSpaces(result.UserOwn, userOwnMax-resultsLen[i].UserOwn)
		}
		if resultsLen[i].GroupOwn < GroupOwnMax {
			result.GroupOwn = addSpaces(result.GroupOwn, GroupOwnMax-resultsLen[i].GroupOwn)
		}
		if resultsLen[i].Size < sizeMax {
			result.Size = addSpaces(result.Size, sizeMax-resultsLen[i].Size)
		}
		if resultsLen[i].Date < dateMax {
			result.Date = addSpaces(result.Date, dateMax-resultsLen[i].Date)
		}
		fmt.Println(result.Perms, result.Num, result.UserOwn, result.GroupOwn, result.Size, result.Date, result.Name)
	}
}

func SortFiles(files []fs.FileInfo) []fs.FileInfo {
	for i := 0; i < len(files); i++ {
		if i+1 < len(files) && strings.ToLower(files[i].Name()) > strings.ToLower(files[i+1].Name()) {
			tmp := files[i]
			files[i] = files[i+1]
			files[i+1] = tmp
			i = -1
		}
	}
	return files
}

func RevSortFiles(files []fs.FileInfo) []fs.FileInfo {
	for i := 0; i < len(files); i++ {
		if i+1 < len(files) && strings.ToLower(files[i].Name()) < strings.ToLower(files[i+1].Name()) {
			tmp := files[i]
			files[i] = files[i+1]
			files[i+1] = tmp
			i = -1
		}
	}
	return files
}

// func SortPaths(Paths *[]string) {
// 	for i := 0; i < len(*Paths); i++ {
// 		if i+1 < len((*Paths)) && strings.ToLower((*Paths)[i]) > strings.ToLower((*Paths)[i+1]) {
// 			tmp := (*Paths)[i]
// 			(*Paths)[i] = (*Paths)[i+1]
// 			(*Paths)[i+1] = tmp
// 			i = -1
// 		}
// 	}
// }

// func RevSortPaths(Paths *[]string) {
// 	for i := 0; i < len(*Paths); i++ {
// 		if i+1 < len((*Paths)) && strings.ToLower((*Paths)[i]) < strings.ToLower((*Paths)[i+1]) {
// 			tmp := (*Paths)[i]
// 			(*Paths)[i] = (*Paths)[i+1]
// 			(*Paths)[i+1] = tmp
// 			i = -1
// 		}
// 	}
// }

func SortPaths(Paths *[]myStructs.Path) {
	for i := 0; i < len(*Paths); i++ {
		if i+1 < len((*Paths)) && strings.ToLower((*Paths)[i].Path) < strings.ToLower((*Paths)[i+1].Path) {
			tmp := (*Paths)[i]
			(*Paths)[i] = (*Paths)[i+1]
			(*Paths)[i+1] = tmp
			i = -1
		}
	}
}

func RevSortPaths(Paths *[]myStructs.Path) {
	for i := 0; i < len(*Paths); i++ {
		if i+1 < len((*Paths)) && strings.ToLower((*Paths)[i].Path) > strings.ToLower((*Paths)[i+1].Path) {
			tmp := (*Paths)[i]
			(*Paths)[i] = (*Paths)[i+1]
			(*Paths)[i+1] = tmp
			i = -1
		}
	}
}
