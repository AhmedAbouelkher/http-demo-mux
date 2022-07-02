package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	SeedBookIntoMemory()
	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		json.NewEncoder(res).Encode(map[string]any{
			"message": "Welcome to my api",
		})
	}).Methods(http.MethodGet)

	// Books Router
	booksRouter := router.PathPrefix("/books").Subrouter()
	
	booksRouter.HandleFunc("/", FetchAllBooksHandler).Methods(http.MethodGet)
	booksRouter.HandleFunc("/", CreateBookHandler).Methods(http.MethodPost)
	booksRouter.HandleFunc("/{id}", FetchBookByIdHandler).Methods(http.MethodGet)
	booksRouter.HandleFunc("/{id}", UpdateBookByIdHandler).Methods(http.MethodPut)
	booksRouter.HandleFunc("/{id}", DeleteBookByIdHandler).Methods(http.MethodDelete)

	// Middleware
	router.Use(contentTypeApplicationJsonMiddleware)

	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8005",
	}
	log.Println("Starting the server...")
	log.Fatal(srv.ListenAndServe())
}

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request)  {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}