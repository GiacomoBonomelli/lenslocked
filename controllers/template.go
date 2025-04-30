package controllers

import (
	"net/http"
)

type Template interface {
	// i placeholder non sono obbligatori
	Execute(w http.ResponseWriter,r *http.Request, data interface{})
}
