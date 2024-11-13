package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func PrintOutput(args Args, resp *http.Response, file *os.File) {
	fmt.Println(resp.ContentLength)
	bar := progressbar.NewOptions(int(resp.ContentLength),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(100),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionUseIECUnits(true),
		progressbar.OptionShowTotalBytes(true))
	if int(resp.ContentLength) == -1 {
		for i := 0; i < 1000; i++ {
			bar.Add(1)
		}
	} else {
		for i := 0; i < int(resp.ContentLength); i++ {
			bar.Add(1)
		}
	}
	io.Copy(io.MultiWriter(file, bar), resp.Body)
}
