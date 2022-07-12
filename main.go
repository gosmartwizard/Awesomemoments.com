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

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port:4949")
	http.ListenAndServe(":4949", r)
}
