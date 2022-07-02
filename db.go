package main

import (
	"errors"
	"math/rand"
	"strconv"
)

var books []Book
var authors []Author

func SeedBookIntoMemory() {
	authors = append(authors, Author{
		ID: 1,
		FirstName: "Ahmed",
		LastName: "Egypt",
	})
	
	for i := 0; i < 5; i++ {
		rand.Seed(int64(i))
		books = append(books, Book{
			ID: int8(rand.Intn(8)),
			Title: strconv.Itoa(rand.Int()),
			Price: rand.Float32(),
			Author: &authors[0],
		})
	}
}

func CreateNewBook(book *Book) error {
	book.ID = int8(rand.Intn(8))
	book.Author = &authors[0]
	books = append(books, *book)
	return nil
}

func FetchBookById(id int64) (book Book, err error) {
	for _, book := range books {
		if book.ID == int8(id) {
			return book, nil
		}
	}
	return Book{}, errors.New("Book was not found")
}

func DeleteBookById(id int64) (book Book, err error) {
	for i, book := range books {
		if book.ID == int8(id) {
			books = removeIndex(books, i)
			return book, nil
		}
	}
	return Book{}, errors.New("Book was not found")
}

func removeIndex(s []Book, index int) []Book {
    return append(s[:index], s[index+1:]...)
}

func UpdateBookById(id int64, updatedBook *Book) (book Book, err error) {
	for i, book := range books {
		if book.ID == int8(id) {
			books[i] = Book{
				ID: int8(id),
				Title: updatedBook.Title,
				Price: updatedBook.Price,
				Author: book.Author,
			}
			return books[i], nil
		}
	}
	return Book{}, errors.New("Book was not found")
}


