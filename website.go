package main

import (
    "log"
    "net/http"
	"html/template"
)

type Page struct {
    Title string
}

func main() {
	// Route handlers
    http.HandleFunc("/", indexHandler)

	// Static handler
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Listen to port and serve
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderPage(w http.ResponseWriter, tmpl string, p *Page) {
    // Parse file
	t, err := template.ParseFiles("templates/" + tmpl + ".html")

    // Handle parse errors
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Execute template
    err = t.Execute(w, p)

    // Handle execute errors
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}