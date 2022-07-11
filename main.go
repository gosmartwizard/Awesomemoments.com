package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> Welcome to the awesome and aha site!<h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:go@smartwizard.io\">go@smartwizard.io</a>.</p>")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on port:4949")
	http.ListenAndServe(":4949", nil)
}
