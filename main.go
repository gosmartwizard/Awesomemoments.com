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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@Awesomemoments.com">support@Awesomemoments.com</a>
  </li>
</ul>
`)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, " page not found", http.StatusNotFound)
	}
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathHandler(w, r)
}
func main() {
	var router Router
	fmt.Println("Starting the server on port:4949")
	http.ListenAndServe(":4949", router)
}
