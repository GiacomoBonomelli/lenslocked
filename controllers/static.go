package controllers

import (
	"net/http"
)

// Viene eseguito il template
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

// Viene eseguito il template FAQ
func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string // si potrebbe usare il tipo template.HTMl. Dice che è ok renderizzare la risposta come
		// HTML. NON UTILIZZARLO IN PRODUZIONE dato che potrebbe permettere attacchi di tipo XSS.
		// Accertarsi che la fonte da dove provengono i dati, sia sicura.
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days.",
		},

		{
			Question: "What are your support hours?",
			Answer:   "We have support staff....",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
		{
			Question: "Where is your office located?",
			Answer:   "Our entire team is remote",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}

}
