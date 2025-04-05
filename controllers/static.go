package controllers

import (
	"net/http"

	"github.com/GiacomoBonomelli/lenslocked/views"
)

// restituisce una handler func
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
