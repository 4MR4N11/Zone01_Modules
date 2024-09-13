package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

type Result struct {
	perms    string
	num      string
	userOwn  string
	GroupOwn string
	size     string
	date     string
	name     string
}

type ResultLen struct {
	perms    int
	num      int
	userOwn  int
	GroupOwn int
	size     int
	date     int
}

func countDir(dir fs.FileInfo) int {
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

func getTotalSize(files []fs.FileInfo) int {
	totalSize := 4096
	result := 0
	for _, file := range files {
		if int(file.Size()) == 0 {
			totalSize += 1024
		} else {
			totalSize += int(file.Size())
		}
	}
	result = totalSize / 1024
	// if result%1024 != 0 {
	// 	result += 1
	// }
	return result
}

func getSymLink(file fs.FileInfo) (string, fs.FileInfo) {
	info, _ := os.Lstat(file.Name())
	symLink := ""
	if info.Mode()&os.ModeSymlink != 0 {
		var err error
		symLink, err = os.Readlink(file.Name())
		if err != nil {
			return "", nil
		}
	}
	return symLink, info
}

func findMaxLen(resultsLen []ResultLen, field string) int {
	max := 0
	for _, res := range resultsLen {
		switch field {
		case "perms":
			if res.perms > max {
				max = res.perms
			}
		case "num":
			if res.num > max {
				max = res.num
			}
		case "userOwn":
			if res.userOwn > max {
				max = res.userOwn
			}
		case "GroupOwn":
			if res.GroupOwn > max {
				max = res.GroupOwn
			}
		case "size":
			if res.size > max {
				max = res.size
			}
		case "date":
			if res.date > max {
				max = res.date
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

func printResults(results []Result, resultsLen []ResultLen) {
	permsMax := findMaxLen(resultsLen, "perms")
	numMax := findMaxLen(resultsLen, "num")
	userOwnMax := findMaxLen(resultsLen, "userOwn")
	GroupOwnMax := findMaxLen(resultsLen, "GroupOwn")
	sizeMax := findMaxLen(resultsLen, "size")
	dateMax := findMaxLen(resultsLen, "date")
	for i, result := range results {
		if resultsLen[i].perms < permsMax {
			result.perms = addSpaces(result.perms, permsMax-resultsLen[i].perms)
		}
		if resultsLen[i].num < numMax {
			result.num = addSpaces(result.num, numMax-resultsLen[i].num)
		}
		if resultsLen[i].userOwn < userOwnMax {
			result.userOwn = addSpaces(result.userOwn, userOwnMax-resultsLen[i].userOwn)
		}
		if resultsLen[i].GroupOwn < GroupOwnMax {
			result.GroupOwn = addSpaces(result.GroupOwn, GroupOwnMax-resultsLen[i].GroupOwn)
		}
		if resultsLen[i].size < sizeMax {
			result.size = addSpaces(result.size, sizeMax-resultsLen[i].size)
		}
		if resultsLen[i].date < dateMax {
			result.date = addSpaces(result.date, dateMax-resultsLen[i].date)
		}
		fmt.Println(result.perms, result.num, result.userOwn, result.GroupOwn, result.size, result.date, result.name)
	}
}

func main() {
	d, _ := os.Open(".")
	defer d.Close()
	results := []Result{}
	resultsLen := []ResultLen{}
	files, _ := d.Readdir(-1)
	totalSize := getTotalSize(files)
	fmt.Println("total", totalSize)
	for _, file := range files {
		tmp := Result{}
		lenTmp := ResultLen{}
		symlink, info := getSymLink(file)
		if len(symlink) != 0 {
			tmp.name = file.Name() + " -> " + symlink
			tmp.perms = info.Mode().String()
			tmp.perms = "l" + tmp.perms[1:]
			lenTmp.perms = len(tmp.perms)
		} else {
			tmp.name = file.Name()
			tmp.perms = file.Mode().Perm().String()
			lenTmp.perms = len(tmp.perms)
		}
		tmp.num = "1"
		lenTmp.num = len(tmp.num)
		if file.IsDir() {
			tmp.num = strconv.Itoa((countDir(file) + 2))
			lenTmp.num = len(tmp.num)
			tmp.perms = "d" + tmp.perms[1:]
		}
		User, _ := user.LookupGroupId(strconv.Itoa(int(file.Sys().(*syscall.Stat_t).Uid)))
		tmp.userOwn = User.Name
		lenTmp.userOwn = len(tmp.userOwn)
		Group, _ := user.LookupId((strconv.Itoa(int(file.Sys().(*syscall.Stat_t).Gid))))
		tmp.GroupOwn = Group.Name
		lenTmp.GroupOwn = len(tmp.GroupOwn)
		tmp.date = file.ModTime().Format("Jan 02 15:04")
		lenTmp.date = len(tmp.date)
		tmp.size = strconv.FormatInt(file.Size(), 10)
		lenTmp.size = len(tmp.size)
		results = append(results, tmp)
		resultsLen = append(resultsLen, lenTmp)
	}
	printResults(results, resultsLen)
}
