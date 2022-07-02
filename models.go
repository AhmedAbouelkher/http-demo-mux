package main

type Author struct {
	ID int8 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Book struct {
	ID int8 `json:"id"`
	Title string `json:"title"`
	Price float32 `json:"price"`
	Author *Author `json:"author"`
}