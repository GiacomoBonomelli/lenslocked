package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

// creiamo un oggetto Template per poter lavorare con diverse tipologie di file (es:HTML,JSON,ecc...)
type Template struct {
	htmlTpl *template.Template // questo attributo ha come tipo un oggetto template.
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// Per fare il parsing dei templates indipendentemente dalla loro posizione nel file system
func ParseFS(fs fs.FS, patterns ...string) (Template, error) { //variadic parameter
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML{
				return `<!-- TODO: Implement the csrfField -->`
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
	//tpl è una variabile già esistente mentre err,no
	//all'interno della funzione,patterns è usata come slice di stringhe
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %w", err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}

/* func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath) //crea un oggetto di tipo template
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
} */

// Metodo per un oggetto di tipo Template
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
