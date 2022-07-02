package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FetchAllBooksHandler(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(books)
}

func CreateBookHandler(res http.ResponseWriter, req *http.Request) {
	var book Book

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	CreateNewBook(&book)

	json.NewEncoder(res).Encode(book)
}

func FetchBookByIdHandler(res http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]any{
			"message": fmt.Sprintln("Error", err),
		})
		return
	}
	
	book, err := FetchBookById(id)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]any{
			"message": "book is not found",
		})
		return
	}
	
	json.NewEncoder(res).Encode(book)
}

func UpdateBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]any{
			"message": fmt.Sprintln("Error", err),
		})
		return
	}

	var placeHolderBook Book

	decodingErr := json.NewDecoder(req.Body).Decode(&placeHolderBook)
	if decodingErr != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook, updateErr := UpdateBookById(id, &placeHolderBook)

	if updateErr != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]any{
			"message": fmt.Sprintln("Error", updateErr.Error()),
		})
		return
	}

	json.NewEncoder(res).Encode(updatedBook)
}

func DeleteBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 10, 32)
	
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]any{
			"message": fmt.Sprintln("Error", err),
		})
		return
	}

	deleted, err := DeleteBookById(id)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]any{
			"message": "book is not found",
		})
		return
	}

	json.NewEncoder(res).Encode(deleted)
}