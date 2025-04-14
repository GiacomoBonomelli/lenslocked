package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/GiacomoBonomelli/lenslocked/templates"
	"github.com/GiacomoBonomelli/lenslocked/controllers"
	"github.com/GiacomoBonomelli/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath) // fa il parsing del file html e mi restituisce un oggetto di tipo Template
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	userID := chi.URLParam(r, "userID")
	w.Write([]byte(fmt.Sprintf("hi %v", userID)))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you were looking for</h1><p>Please email us if you keep being sent to an invalid page.</p>")
}

/* func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
		//notFoundHandler(w, r)
	}
}  */

/*
Un modo per scrivere un router

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
} */

func main() {
	//var router Router
	// Diversi modi per gestire le routes
	//http.HandleFunc("/", pathHandler)
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))

	r := chi.NewRouter()
	//check the template for error and parse it
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	// Route con logger
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger) // logger solo per questo gruppo
		r.Get("/user/{userID}", userHandler)
	})

	r.NotFound(notFoundHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
