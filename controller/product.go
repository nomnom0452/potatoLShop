package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/nomnom0452/potatolshop/model"
)

func ProductDetail(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	queryId := vals.Get("id")
	id, err := strconv.Atoi(queryId)

	if err != nil {
		fmt.Println("Cannot convert Id", err)
	}

	product, err := model.ProductById(id)

	if err != nil {
		fmt.Println("cannot load product, ", err)
	}

	temp := template.Must(template.ParseFiles("view/product.html"))
	temp.Execute(w, product)
}
