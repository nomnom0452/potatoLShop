package model

import (
	"github.com/nomnom0452/potatolshop/data"
)

type Product struct {
	Id    int    `json:"ID"`
	Name  string `json:"Name"`
	Qty   int    `json:"Quantity"`
	Price int    `json:"Price"`
	Desc  string `json:"Desc"`
}

func ProductById(id int) (product Product, err error) {
	row := data.Db.QueryRow("SELECT * FROM items WHERE ID=?", id)

	err = row.Scan(&product.Id, &product.Name, &product.Qty, &product.Price, &product.Desc)

	return
}

func Products() (products []Product, err error) {
	rows, err := data.Db.Query("SELECT *FROM items")

	if err != nil {
		return
	}

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Qty, &product.Price, &product.Desc)
		if err != nil {
			return
		}
		products = append(products, product)
	}

	defer rows.Close()

	return
}

func AddItemToBasketBy(productId, userId int) (status bool, err error) {
	result, err := data.Db.Exec("INSERT INTO Basket(productId, userId) VALUES(?,?)", productId, userId)
	if res, _ := result.RowsAffected(); res == 0 {
		status = false
		return
	}

	status = true
	return
}
