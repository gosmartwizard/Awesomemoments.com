package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gosmartwizard/Awesomemoments.com/controllers"
	"github.com/gosmartwizard/Awesomemoments.com/models"
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

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userService,
	}

	tpl = views.Must(views.ParseFS(templates.FS, "helpline.gohtml"))
	r.Get("/helpline", controllers.StaticHandler(tpl))

	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)

	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port:4949")

	http.ListenAndServe(":4949", r)
}
