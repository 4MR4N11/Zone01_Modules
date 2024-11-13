package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"wget/internal/utils"
)

func HttpDownloader(args utils.Args) {
	client := &http.Client{}
	for i, url := range args.Url {
		if args.UrlFile[i] == "" {
			args.UrlFile[i] = "index.html"
			// url = strings.TrimRight(url, "/") + "/" + args.UrlFile[i]
		}
		fmt.Println(url)
		startTime := time.Now()
		fmt.Println("start at " + startTime.Format("2006-01-02 15:04:05"))
		fmt.Print("sending request, awaiting response... ")
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error downloading", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Println("status " + resp.Status)

		contentSize := resp.ContentLength
		if contentSize == -1 {
			fmt.Println("Warning: Unknown content size, progress will not be accurate")
		}
		if contentSize != -1 {
			fmt.Printf("content size: %d [~%dMB]\n", contentSize, contentSize/1024)
		}

		outputFile, _ := os.Create(args.UrlFile[i])
		fmt.Printf("saving file to: ./%s\n", args.UrlFile[i])
		defer outputFile.Close()

		// bar := progressbar.NewOptions(int(contentSize),
		// 	progressbar.OptionShowBytes(true),
		// 	progressbar.OptionSetWidth(100),
		// 	progressbar.OptionSetTheme(progressbar.Theme{
		// 		Saucer:        "=",
		// 		SaucerHead:    ">",
		// 		SaucerPadding: " ",
		// 		BarStart:      "[",
		// 		BarEnd:        "]",
		// 	}))

		var downloadedBytes int64
		buffer := make([]byte, 8*1024)
		for {
			n, err := resp.Body.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error during download:", err)
				os.Exit(1)
			}
			outputFile.Write(buffer[:n])
			downloadedBytes += int64(n)
			elapsedTime := time.Since(startTime)
			downloadSpeed := float64(downloadedBytes) / 1024 / 1024 / elapsedTime.Seconds()
			progressPercentage := float64(downloadedBytes) / float64(contentSize) * 100
			elapsedTimeStr := fmt.Sprintf("%.0fs", elapsedTime.Seconds())
			fmt.Printf("\r%v / %v [%s] %.2f%% %.2f MiB/s %v",
				formatBytes(downloadedBytes),
				formatBytes(contentSize),
				progressbar(progressPercentage),
				progressPercentage,
				downloadSpeed,
				elapsedTimeStr,
			)

		}
		fmt.Printf("\nDownloaded [%s]\n", args.Url[i])
		fmt.Printf("finished at [%s]\n", time.Now().Format("2006-01-02 15:04:05"))
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
