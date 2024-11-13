package downloader

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"wget/internal/utils"
)

func HttpDownloader(args utils.Args) {
	client := &http.Client{}
	for i, url := range args.Url {
		if args.UrlFile[i] == "" {
			args.UrlFile[i] = "index.html"
		}
		output := ""
		startTime := time.Now()
		output += fmt.Sprintln("start at " + startTime.Format("2006-01-02 15:04:05"))
		output += fmt.Sprint("sending request, awaiting response... ")
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error downloading", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		output += fmt.Sprintln("status " + resp.Status)
		if resp.StatusCode != 200 {
			os.Exit(1)
		}
		contentSize := resp.ContentLength
		if contentSize == -1 {
			output += fmt.Sprintln("Warning: Unknown content size, progress will not be accurate")
			contentSize = 0
		} else if contentSize != -1 {
			output += fmt.Sprintf("content size: %d [~%s]\n", contentSize, formatBytestoMb(contentSize))
		}
		if len(args.Output) > 0 {
			args.UrlFile[i] = args.Output
		}
		if len(args.Path) > 0 {
			if strings.HasPrefix(args.Path, "~") {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Println("Error", err)
					os.Exit(1)
				}
				args.Path = filepath.Join(homeDir, args.Path[1:])
			}
			args.UrlFile[i] = args.Path + "/" + args.UrlFile[i]
		} else {
			args.UrlFile[i] = "./" + args.UrlFile[i]
		}
		outputFile, _ := os.Create(args.UrlFile[i])
		output += fmt.Sprintf("saving file to: %s\n", args.UrlFile[i])
		defer outputFile.Close()
		var downloadedBytes int64
		buffer := make([]byte, 8*1024)
		if !args.Bdownload {
			fmt.Print(output)
			output = ""
		}
		for {
			n, err := resp.Body.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				output += fmt.Sprintln("Error during download:", err)
				os.Exit(1)
			}
			outputFile.Write(buffer[:n])
			downloadedBytes += int64(n)
			elapsedTime := time.Since(startTime)
			downloadSpeed := float64(downloadedBytes) / 1024 / 1024 / elapsedTime.Seconds()
			progressPercentage := 0.0
			remainingBytes := contentSize - downloadedBytes
			remainingTimeStr := ""
			if downloadSpeed > 0 {
				if contentSize == 0 {
					remainingTimeStr = "unknown"
				} else {
					remainingTime := int(remainingBytes / int64(downloadSpeed*1024*1024))
					remainingTimeStr = fmt.Sprintf("%ds", remainingTime)
				}
			} else {
				remainingTimeStr = "unknown"
			}
			if contentSize == 0 {
				progressPercentage = 0.0
			} else {
				progressPercentage = float64(downloadedBytes) / float64(contentSize) * 100
			}
			if !args.Bdownload {
				fmt.Printf("\r%v / %v [%s] %.2f%% %.2f MiB/s %s   ",
					formatBytes(downloadedBytes),
					formatBytes(contentSize),
					progressbar(progressPercentage),
					progressPercentage,
					downloadSpeed,
					remainingTimeStr,
				)
			}
		}
		output += fmt.Sprintf("\nDownloaded [%s]\n", args.Url[i])
		output += fmt.Sprintf("finished at [%s]\n", time.Now().Format("2006-01-02 15:04:05"))
		if !args.Bdownload {
			fmt.Print(output)
			output = ""
		} else {
			utils.PrintOutput(args, output)
		}
	}
}

func progressbar(progressPercentage float64) string {
	bar := ""
	for i := 0; i < int(progressPercentage); i++ {
		bar += "="
	}
	for i := 0; i < 100-int(progressPercentage); i++ {
		bar += " "
	}
	return bar
}

func formatBytes(bytes int64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	} else if bytes < 1024*1024 {
		return fmt.Sprintf("%.2f KiB", float64(bytes)/1024)
	} else if bytes < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MiB", float64(bytes)/1024/1024)
	} else {
		return fmt.Sprintf("%.2f GiB", float64(bytes)/1024/1024/1024)
	}
}

func formatBytestoMb(bytes int64) string {
	if bytes < 1000 {
		return fmt.Sprintf("%d B", bytes)
	} else if bytes < 1000*1000 {
		return fmt.Sprintf("%.2f KB", math.Round(float64(bytes)/1000))
	} else if bytes < 1000*1000*1000 {
		return fmt.Sprintf("%.2f MB", math.Round(float64(bytes)/1000/1000))
	} else {
		return fmt.Sprintf("%.2f GB", math.Round(float64(bytes)/1000/1000/1000))
	}
}
