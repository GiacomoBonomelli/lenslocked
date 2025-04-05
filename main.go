package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	tpl, err := template.ParseFiles(filepath) //crea un oggetto di tipo template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w,nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := "templates/home.gohtml"
	executeTemplate(w, tplpath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := "templates/contact.gohtml"
	executeTemplate(w, tplpath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := "templates/faq.gohtml"
	executeTemplate(w, tplpath)
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

/* type Router struct{}

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
	r:= chi.NewRouter()
	 

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	// Route con logger
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger) // logger solo per questo gruppo
		r.Get("/user/{userID}", userHandler)
	})

	r.NotFound(notFoundHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
