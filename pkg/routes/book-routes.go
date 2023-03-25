package routes

import (
	"github.com/gorilla/mux"
	controller "github.com/subhammahanty235/golang-book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")

	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/update/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/delete/{bookId}", controller.DeleteBook).Methods("DELETE")
}
