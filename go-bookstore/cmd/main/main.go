package main

import (
	"bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "bookstore/pkg/routes"
	// "bookstore/pkg/routes"
	// // _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("hi")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
