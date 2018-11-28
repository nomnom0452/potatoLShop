package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/nomnom0452/potatolshop/model"
)

func BasketCart(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("view/basketcart.html"))
	temp.Execute(w, nil)
}

func AddBasket(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("due-sue-resu")
	productId, _ := strconv.Atoi(r.URL.Query().Get("Id"))

	session, _ := model.SessionGET(cookie.Value)

	result, err := model.AddItemToBasketBy(productId, session.UserId)

	if err != nil {
		fmt.Println("Fail Adding Item")
	}

	if result == true {
		http.Redirect(w, r, "/wishlist", http.StatusFound)
	} else {
		fmt.Fprintln(w, "fail to adding item")
	}
}
