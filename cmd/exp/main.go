package main

import (
	"os"
	"html/template"
)

type User struct {
	Name string
	Bio string
	//Age int
	//Meta UserMeta
}
/*type UserMeta struct {
	Visits int
}*/

func main() {
	t,err:= template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Giacomo",
		Bio: `<script>alert("You've been hacked")</script>`,
		//Age: 20,
		//Meta: UserMeta{
		//	Visits: 10,
		//},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}