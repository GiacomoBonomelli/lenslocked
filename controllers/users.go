package controllers

import (
	"fmt"
	"net/http"

	"github.com/GiacomoBonomelli/lenslocked/models"
)

type Users struct {
	Templates struct {
		New Template // it allows to create a Template object
		SignIn Template
	}
	UserService *models.UserService // collegamento tra il controller e il model. Permette di creare gli utenti
}

//new and edit to render forms
//create and update to process forms

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	//We need a view to render
	//We need to parse the view before the controller
	//handles this request.
	//So its best practice to parse it somewhere else.

	//execute the template
	var data struct{
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w,data )
}

//function that is going to handle the post request made on the form submission
func (u Users) Create(w http.ResponseWriter, r *http.Request){
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v", user)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	//We need a view to render
	//We need to parse the view before the controller
	//handles this request.
	//So its best practice to parse it somewhere else.

	//execute the template
	var data struct{
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w,data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request){
	var data struct{
		Email string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	cookie:= http.Cookie{
		Name: "email", // di solito sono brevi per non appesantire le risposte inviate dal server
		Value: user.Email,
		Path: "/",
	}	
	http.SetCookie(w,&cookie)
	fmt.Fprintf(w, "User authenticated: %+v", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request){
	email,err:=r.Cookie("email")
	if err!=nil{
		fmt.Fprint(w,"The email cookie could not be read")
		return
	}
	fmt.Fprintf(w,"Email cookie: %s\n",email.Value)
}