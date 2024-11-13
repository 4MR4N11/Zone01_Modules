package cmd

import (
	"os"
	"strings"

	"wget/internal/utils"
)

func Init() utils.Args {
	parsedArgs := utils.Args{}
	mirror := utils.Mirror{}
	parsedArgs.Bdownload = false
	parsedArgs.LimitRate = ""
	parsedArgs.MultiDownload = ""
	parsedArgs.Output = ""
	parsedArgs.Path = ""
	parsedArgs.Url = nil
	parsedArgs.UrlFile = nil
	mirror.Active = false
	mirror.ConvertLinks = false
	mirror.Exclude = ""
	mirror.Reject = ""
	parsedArgs.Mirror = mirror
	return parsedArgs
}

func GetOptions(arg string, parsedArgs *utils.Args) {
	if arg == "B" {
		parsedArgs.Bdownload = true
	} else if strings.HasPrefix(arg, "O=") {
		parsedArgs.Output = strings.TrimPrefix(arg, "O=")
	} else if strings.HasPrefix(arg, "P=") {
		parsedArgs.Path = strings.TrimPrefix(arg, "P=")
	} else if strings.HasPrefix(arg, "-rate-limit=") {
		parsedArgs.LimitRate = strings.TrimPrefix(arg, "-rate-limit=")
	} else if strings.HasPrefix(arg, "i=") {
		parsedArgs.LimitRate = strings.TrimPrefix(arg, "i=")
	} else if strings.HasPrefix(arg, "-mirror") {
		parsedArgs.Mirror.Active = true
	} else if strings.HasPrefix(arg, "-reject=") {
		parsedArgs.Mirror.Reject = strings.TrimPrefix(arg, "-reject=")
	} else if strings.HasPrefix(arg, "R=") {
		parsedArgs.Mirror.Reject = strings.TrimPrefix(arg, "R=")
	} else if strings.HasPrefix(arg, "-exclude") {
		parsedArgs.Mirror.Exclude = strings.TrimPrefix(arg, "-exclude")
	} else if strings.HasPrefix(arg, "X=") {
		parsedArgs.Mirror.Exclude = strings.TrimPrefix(arg, "X=")
	} else if strings.HasPrefix(arg, "-convert-links") {
		parsedArgs.Mirror.ConvertLinks = true
	} else if strings.HasPrefix(arg, "-help") {
		utils.GetHelp()
		os.Exit(0)
	} else {
		utils.OptionsError(arg)
		os.Exit(0)
	}
}

func ParseArgs(args []string) utils.Args {
	parsedArgs := Init()
	for _, arg := range args {
		if len(arg) > 0 && arg[0] == '-' {
			GetOptions(arg[1:], &parsedArgs)
		} else if len(arg) > 0 {
			if strings.HasPrefix(arg, "http://") {
				tmp := strings.Split(strings.TrimLeft(arg, "http://"), "/")
				if len(tmp) > 1 {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, tmp[len(tmp)-1])
				} else {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, "")
				}
				parsedArgs.Url = append(parsedArgs.Url, arg)
			} else if strings.HasPrefix(arg, "https://") {
				tmp := strings.Split(strings.TrimLeft(arg, "https://"), "/")
				if len(tmp) > 1 {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, tmp[len(tmp)-1])
				} else {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, "")
				}
				parsedArgs.Url = append(parsedArgs.Url, arg)
			} else if strings.HasPrefix(arg, "ftp://") {
				tmp := strings.Split(strings.TrimLeft(arg, "ftp://"), "/")
				if len(tmp) > 1 {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, tmp[len(tmp)-1])
				} else {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, "")
				}
				parsedArgs.Url = append(parsedArgs.Url, arg)
			} else {
				tmp := strings.Split(arg, "/")
				if len(tmp) > 1 {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, tmp[len(tmp)-1])
				} else {
					parsedArgs.UrlFile = append(parsedArgs.UrlFile, "")
				}
				parsedArgs.Url = append(parsedArgs.Url, "http://"+arg)
			}
		}
	}
	return parsedArgs
}
