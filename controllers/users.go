package controllers

import (
	"net/http"
	"github.com/GiacomoBonomelli/lenslocked/views"
)

type Users struct {
	Templates struct {
		New views.Template // it allows to create a Template object
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
	u.Templates.New.Execute(w, nil)
}
