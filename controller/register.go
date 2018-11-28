package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/nomnom0452/potatolshop/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("view/register.html"))
	temp.Execute(w, nil)
}

func SignAuth(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := model.User{0, name, age, email, password}
	err := model.CreateCust(user)

	if err != nil {
		fmt.Println("something is wrong ", err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
