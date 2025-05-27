package controllers

import (
	"fmt"
	"net/http"
	
	"github.com/GiacomoBonomelli/lenslocked/models"
)

// Definisce la struttura Users che contiene i template,
// il servizio per gestire le operazioni relative agli utenti e il servizio per gestire le sessioni.
// E' un layer di astrazione tra la view e il model.
type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

// In questa funzione vengono gestiti i dati in entrata per poi inviarli alle funzioni che
// si occupano di creare l'utente e di creare la sessione.
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Crea l'utente nel database
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	// Viene creata la sessione per l'utente appena creato
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		// TODO: Long term, we should  show a warning about not being able to sign
		//the user in
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	// Creiamo e settiamo il cookie con valore pari al token della sessione
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

// Invoca la view per il login
func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)
}

// In questa funzione vengono gestiti i dati in entrata per poi inviarli alle funzioni che
// si occupano di autenticare l'utente e di creare la sessione.
func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	// Viene creata la sessione per l'utente appena autenticato
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// Creiamo e settiamo il cookie con valore pari al token della sessione
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

// Questa funzione recupera l'utente corrente dalla sessione
// e lo visualizza nella pagina.
// Viene utilizzata per verificare se l'utente è autenticato.
// Se l'utente non è autenticato, viene rediretto alla pagina di login.
// Se l'utente è autenticato, viene visualizzato il nome dell'utente.

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	user, err := u.SessionService.User(token)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Current user: %s\n", user.Email)
}

func (u Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	err = u.SessionService.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	deleteCookie(w, CookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)
}
