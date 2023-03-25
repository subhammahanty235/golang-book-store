package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subhammahanty235/golang-book-store/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
