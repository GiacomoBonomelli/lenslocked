package main

import (
	"fmt"
	"net/http"

	"github.com/GiacomoBonomelli/lenslocked/controllers"
	"github.com/GiacomoBonomelli/lenslocked/models"
	"github.com/GiacomoBonomelli/lenslocked/templates"
	"github.com/GiacomoBonomelli/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	r := chi.NewRouter()

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	// Comunicazione con la base di dati
	// Crea un servizio per gestire le operazioni relative agli utenti
	userService := models.UserService{
		DB: db,
	}

	// Crea un servizio per gestire le sessioni
	sessionService := models.SessionService{
		DB: db,
	}

	// Crea un controller per gestire le operazioni relative agli utenti
	usersC := controllers.Users{
		UserService: &userService,
		SessionService: &sessionService,
	}

	// Setup our templates
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))
	// r.Get("/signup", TimerMiddleware(usersC.New))
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	// Gestione dei form di registrazione e login
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")

	// Crea una chiave per la protezione CSRF
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix this before deploying
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"localhost:3000"}),
	)
	// Avvia il server
	http.ListenAndServe(":3000", csrfMw(r))
}

// Uncomment the TimerMiddleware func and use it above in main() to see
// it in action.
// func TimerMiddleware(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		h(w, r)
// 		fmt.Println("Request time:", time.Since(start))
// 	}
// }
