package routes

import (
	"github.com/gorilla/mux"
	"halyklife-lib/controllers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors", controllers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", controllers.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", controllers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", controllers.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	r.HandleFunc("/readers", controllers.GetReaders).Methods("GET")
	r.HandleFunc("/readers", controllers.CreateReader).Methods("POST")
	r.HandleFunc("/readers/{id}", controllers.GetReader).Methods("GET")
	r.HandleFunc("/readers/{id}", controllers.UpdateReader).Methods("PUT")
	r.HandleFunc("/readers/{id}", controllers.DeleteReader).Methods("DELETE")

	return r
}
