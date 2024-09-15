package flags

import (
	"fmt"
	"os"
	"strings"

	paths "my_ls/models"
)

func dirTraversal(file string, dirs *[]string) error {
	d, err := os.Open(file)
	if err != nil {
		return err
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return err
	}

	for _, f := range files {
		fullPath := file + "/" + f.Name()
		if f.IsDir() {
			*dirs = append(*dirs, fullPath)
			if err := dirTraversal(fullPath, dirs); err != nil {
				return err
			}
		}
	}
	return nil
}

func UpperR(d *os.File, path string) []paths.Path {
	files, _ := d.Readdir(-1)
	dirs := []string{}
	result := []paths.Path{}
	if path == "." {
		path = "./"
		dirs = append(dirs, path)
	} else {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		dirs = append(dirs, path)
	}
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, path+file.Name())
		}
		dirTraversal(path+file.Name(), &dirs)
	}
	for _, dir := range dirs {
		tmp, err := os.Open(dir)
		if err != nil {
			str := fmt.Sprintln(err)
			fmt.Fprint(os.Stderr, strings.Replace(str, "open "+dir, "my_ls: cannot access '"+dir+"'", 1))
		} else {
			p := paths.Path{}
			p.OpenedPath = tmp
			p.Path = dir
			result = append(result, p)
		}
	}
	return result
}
