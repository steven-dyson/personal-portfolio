package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/steven-dyson/personal-portfolio/pages"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>Submitted</p>")
}

func main() {
	// Static assets
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle home and public assets
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			templ.Handler(pages.Home()).ServeHTTP(w, r)
			return
		}
		filePath := filepath.Join("static/public", r.URL.Path[1:])
		http.ServeFile(w, r, filePath)
	})

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
