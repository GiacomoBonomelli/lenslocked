package main

import(
	"fmt"
	"html/template"
)

type Template struct{
	htmlTpl *template.Template // questo è un attributo di tipo template
}


func Parse(filepath string) (Template,error){
	tpl, err := template.ParseFiles(filepath) //crea un oggetto di tipo template
	if err != nil {
		return Template{},fmt.Errorf("parsing template %w",err)
	}
	return Template{
		htmlTpl:tpl,
	},nill
}

//Metodo per un oggetto di tipo Template
func (t Template) Execute(w http.ResponseWriter,data interface{}){
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err := t.htmlTpl.Execute(w,data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
	

func main(){

}