package main

import (
	"html/template"
	"net/http"

	utils "ascii_web/utils"
)

func AsciiArtResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "404 : Not Found", 404)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == "POST" {
		data := r.PostFormValue("textInput")
		banner := r.PostFormValue("bannerType")
		if len(data) == 0 {
			http.Error(w, "Empty input", http.StatusBadRequest)
			return
		}
		result := utils.AsciiArtGenerator(data, banner)
		t, err := template.ParseFiles("templates/result.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func RootPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootPage)
	mux.HandleFunc("/ascii-art", AsciiArtResult)
	http.ListenAndServe("127.0.0.1:3333", mux)
}
