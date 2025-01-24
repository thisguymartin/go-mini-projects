package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type BookWormData struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func loadBookData(filePath string) ([]BookWormData, error) {

	f, err := os.Open(filePath)
	if err != nil {
		panic("done here")
	}
	defer f.Close()

	var bookworms []BookWormData
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		panic("Error here in decorder")
	}

	for b := range len(bookworms) {
		d := bookworms[b]
		fmt.Println(d.Name)
	}
	return bookworms, nil
}

func countBooks(bookworms []BookWormData) map[Book]uint {
	count := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}

func findCommonBook(books []BookWormData) []Book {
	common := countBooks(books)
	var commonBooks []Book
	for book, count := range common {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return commonBooks

}
