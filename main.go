package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type RequestBody struct {
	LongURL string `json:"long_url"`
}

var urlStore = make(map[string]string)

func main() {
	// Serve static files (CSS, JS, etc.) from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	urlHandler := func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		if urlPath == "/favicon.ico" {
			http.NotFound(w, r) // Handle the favicon request by returning 404
			return
		}

		if urlPath == "/" {
			templ := template.Must(template.ParseFiles("template/index.html"))
			templ.Execute(w, nil)
			return
		}

		shortURL := urlPath[len("/"):]
		longURL, exists := urlStore[shortURL]

		if !exists {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, longURL, http.StatusFound)
	}

	shortHandler := func(w http.ResponseWriter, r *http.Request) {

		var reqBody RequestBody

		json.NewDecoder(r.Body).Decode(&reqBody)

		shortKey := generateShortURL(6)
		urlStore[shortKey] = reqBody.LongURL

		resp := map[string]string{"short_url": "http://localhost:8080/" + shortKey}
		json.NewEncoder(w).Encode(resp)

	}
	http.HandleFunc("/", urlHandler)

	http.HandleFunc("/short", shortHandler)
	http.ListenAndServe(":8080", nil)
}

func generateShortURL(n int) string {
	b := make([]byte, n)
	rand.NewSource(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
