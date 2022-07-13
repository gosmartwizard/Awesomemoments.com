package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gosmartwizard/Awesomemoments.com/controllers"
	"github.com/gosmartwizard/Awesomemoments.com/templates"
	"github.com/gosmartwizard/Awesomemoments.com/views"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "helpline.gohtml"))
	r.Get("/helpline", controllers.StaticHandler(tpl))

	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port:4949")

	http.ListenAndServe(":4949", r)
}
