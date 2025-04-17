package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template // it allows to create a Template object
	}
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
		Password string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w,data )
}

//function that is going to handle the post request made on the form submission
func (u Users) Create(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"La mail è:%s",r.FormValue("email"))
	fmt.Fprintf(w,"La password è:%s",r.FormValue("password"))
}