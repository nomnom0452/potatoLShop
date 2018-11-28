package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nomnom0452/potatolshop/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if cookie, _ := r.Cookie("due-sue-resu"); cookie != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	temp := template.Must(template.ParseFiles("view/login.html"))
	temp.Execute(w, nil)
}

func Auth(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	cust, err := model.CustByEmail(email)

	if err != nil {
		fmt.Println("Cannot Load from database, ", err)
	}

	if cust.Compare(password) {
		session, err := cust.CreateSession()
		if err != nil {
			fmt.Println("Error", err)
		}

		cookie := http.Cookie{
			Name:  "due-sue-resu",
			Value: session.Uuid,
		}

		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusResetContent)
	}
}
