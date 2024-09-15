package flags

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"

	myStructs "my_ls/models"
	utils "my_ls/utils"
)

func LFlag(d *os.File, path string, lowerR bool) []string {
	results := []myStructs.Result{}
	resultsLen := []myStructs.ResultLen{}
	files, _ := d.Readdir(-1)
	totalSize := fmt.Sprintln("total", strconv.Itoa(utils.GetTotalSize(files)))
	if lowerR {
		files = utils.RevSortFiles(files)
	} else {
		files = utils.SortFiles(files)
	}
	for _, file := range files {
		if file.Name()[0] == '.' {
			continue
		}
		tmp := myStructs.Result{}
		lenTmp := myStructs.ResultLen{}
		symlink, info := utils.GetSymLink(file, path)
		if len(symlink) != 0 {
			tmp.Name = file.Name() + " -> " + symlink
			tmp.Perms = info.Mode().String()
			tmp.Perms = "l" + tmp.Perms[1:]
			lenTmp.Perms = len(tmp.Perms)
		} else {
			tmp.Name = file.Name()
			tmp.Perms = file.Mode().Perm().String()
			lenTmp.Perms = len(tmp.Perms)
		}
		tmp.Num = "1"
		lenTmp.Num = len(tmp.Num)
		if file.IsDir() {
			tmp.Num = strconv.Itoa((utils.CountDir(file) + 2))
			lenTmp.Num = len(tmp.Num)
			tmp.Perms = "d" + tmp.Perms[1:]
		}
		Group, err := user.LookupGroupId(strconv.Itoa(int(file.Sys().(*syscall.Stat_t).Gid)))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tmp.GroupOwn = Group.Name
		lenTmp.UserOwn = len(tmp.UserOwn)
		user, err := user.LookupId(fmt.Sprintf("%d", file.Sys().(*syscall.Stat_t).Uid))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tmp.UserOwn = user.Username
		lenTmp.GroupOwn = len(tmp.GroupOwn)
		tmp.Date = file.ModTime().Format("Jan 02 15:04")
		lenTmp.Date = len(tmp.Date)
		tmp.Size = strconv.FormatInt(file.Size(), 10)
		lenTmp.Size = len(tmp.Size)
		results = append(results, tmp)
		resultsLen = append(resultsLen, lenTmp)
	}
	return utils.PrintResults(results, resultsLen, totalSize)
}
