package main

import (
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{Title: "Dryden High School Coding Club"}
    renderPage(w, "index", p)
}