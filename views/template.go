package views

import (
	"fmt"
	"html/template"
	"net/http"
	"io/fs"
)

type Template struct {
	htmlTpl *template.Template // questo è un attributo di tipo template
}

func Must(t Template, err error) (Template){
	if err!= nil{
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template,error){ //variadic parameter
	tpl,err := template.ParseFS(fs,patterns...) //all'interno della funzione,patterns è usata come slice di stringhe
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}


func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath) //crea un oggetto di tipo template
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

// Metodo per un oggetto di tipo Template
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
