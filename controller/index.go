package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nomnom0452/potatolshop/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("due-sue-resu")
	if err != nil {
		fmt.Println(err)
	}

	if cookie != nil {

	}

	temp := template.Must(template.ParseFiles("view/index.html"))
	product, err := model.Products()

	if err != nil {
		fmt.Println("Cannot load Product", err)
	}

	temp.Execute(w, product)
}
