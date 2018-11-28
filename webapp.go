package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	c "github.com/nomnom0452/potatolshop/controller"
	"github.com/nomnom0452/potatolshop/data"
)

func main() {
	fmt.Println("Alpha PotatoLshop 1.0 running")

	mux := mux.NewRouter()
	files := http.FileServer(http.Dir(config.static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", c.Index).Methods(http.MethodGet)

	mux.HandleFunc("/login", c.Login).Methods(http.MethodGet)
	mux.HandleFunc("/login", c.Auth).Methods(http.MethodPost)

	mux.HandleFunc("/register", c.Register).Methods(http.MethodGet)
	mux.HandleFunc("/register", c.SignAuth).Methods(http.MethodPost)

	mux.HandleFunc("/product", c.ProductDetail).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    config.Addr,
		Handler: mux,
	}

	fmt.Println(server.ListenAndServe())
}

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("config went wrong. ", err)
	}
	decoder := json.NewDecoder(file)
	config = configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("cannot decode config file. ", err)
	}

	data.Connect()
}

var config configuration

type configuration struct {
	Addr   string
	static string
}
