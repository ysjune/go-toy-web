package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "snack", Email: "snack@naver.com", Age: 23}
	user2 := User{Name: "aaa", Email: "aaa@naver.com", Age: 50}
	users := []User{user, user2}
	templ, err := template.New("Tmpl1").ParseFiles("template/tmpl1.tmpl", "template/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	//templ.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
	templ.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
