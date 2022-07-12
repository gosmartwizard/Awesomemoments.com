package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gosmartwizard/Awesomemoments.com/controllers"
	"github.com/gosmartwizard/Awesomemoments.com/views"
	"net/http"
	"path/filepath"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port:4949")
	
	http.ListenAndServe(":4949", r)
}
