package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/index.html")
	})
	http.HandleFunc("/hyperscript-input", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/hyperscript-input.html")
	})
	http.HandleFunc("/htmx-input", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/htmx-input.html")
	})
	fmt.Println("serving html on :8080")
	http.ListenAndServe(":8080", nil)
}
