package main

import (
	"log"
	"net/http"

	"github.com/mrtzee/go-web-native/config"
	"github.com/mrtzee/go-web-native/controllers/categorycontroller"
	"github.com/mrtzee/go-web-native/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 1. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
