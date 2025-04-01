package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Bio string
	Age int
	Company string
	Meta UserMeta
}
type UserMeta struct {
	Visits int
	Likes int
}

func helloHandler(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	tpl,err:= template.ParseFiles("hello.gohtml")
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	user := User{
		Name: "Giacomo",
		Bio: "Web Dev",
		Age: 33,
		Company: "Unknown",
		Meta: UserMeta{
			Visits: 10,
			Likes: 50,
		},
	}
	err = tpl.Execute(w,user)
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	
}
func main() {
	http.HandleFunc("/hello",helloHandler)
	http.ListenAndServe(":3000",nil)
}