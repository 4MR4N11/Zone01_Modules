package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	utils "ascii_web/utils"
)

var fileData string

func AsciiArtResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400: Bad request.", http.StatusBadRequest)
		return
	}
	if r.Method == "POST" {
		data := r.PostFormValue("textInput")
		banner := r.PostFormValue("bannerType")
		if len(data) == 0 {
			http.Error(w, "400: Bad request.", http.StatusBadRequest)
			return
		}
		result, check := utils.AsciiArtGenerator(data, banner)
		if check == 1 {
			http.Error(w, "400: Bad request.", http.StatusBadRequest)
			return
		}
		t, err := template.ParseFiles("templates/result.html")
		if err != nil {
			http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, result)
		if err != nil {
			http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
			return
		}
		fileData = result
	}
}

func RootPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404: Page not found", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
}

func exportFile(w http.ResponseWriter, r *http.Request) {
	if len(fileData) == 0 {
		http.Error(w, "400: Bad request.", http.StatusBadRequest)
		return
	}
	fileOutput, err := os.Create("result.txt")
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
	_, err = fileOutput.WriteString(fileData)
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
	fileInfo, err := fileOutput.Stat()
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
	w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
	http.ServeFile(w, r, "result.txt")
	fileData = ""
}

func main() {
	http.HandleFunc("/", RootPage)
	http.HandleFunc("/ascii-art", AsciiArtResult)
	http.HandleFunc("/export", exportFile)
	fmt.Println("\033[32mServer started at http://127.0.0.1:3333\033[0m")
	http.ListenAndServe("127.0.0.1:3333", nil)
}
