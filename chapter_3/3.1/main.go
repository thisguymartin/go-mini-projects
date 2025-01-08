package main

import (
	"encoding/json"
	"flag"
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

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "The filepath is required")
	flag.Parse()
	_, err := LoadBookData(filePath)
	if err != nil {
		panic("Failed to look at book data")
	}
	// fmt.Printf("Here are the worms %s", worms)
}

func LoadBookData(filePath string) ([]BookWormData, error) {

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
