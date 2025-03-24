package main

import (
	"flag"
	"fmt"
	"os"
)

type Recommendation struct {
	Book  Book
	Score float64
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "The filepath is required")
	flag.Parse()
	books, err := loadBookData(filePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load data %s", err)
		os.Exit(1)
	}
	fmt.Println("Loading common books")
	commonBooks := findCommonBook(books)
	fmt.Println("Here are your commong books recommended")
	for _, book := range commonBooks {
		fmt.Printf("Title: %s, Author: %s \n", book.Title, book.Author)
		fmt.Println("====")
	}
}
