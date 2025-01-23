package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	bookworkFile string
	want         []BookWormData
	wantErr      bool
}

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestBookCount(t *testing.T) {
	var tt = map[string]struct {
		input []BookWormData
		want  map[Book]uint
	}{
		"use case": {
			input: []BookWormData{
				{Name: "Fadi", Books: []Book{
					{
						Author: "Margaret Atwood",
						Title:  "The Handmaid's Tale",
					},
					{
						Author: "Sylvia Plath",
						Title:  "The Bell Jar",
					},
					{
						Author: "George Lucas",
						Title:  "Star Wars",
					}}},
				{Name: "Peggy", Books: []Book{
					{
						Author: "Margaret Atwood",
						Title:  "The Handmaid's Tale",
					},
					{
						Author: "Sylvia Plath",
						Title:  "The Bell Jar",
					},
				}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
	}

	for name, tc := range tt {

		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			fmt.Print(got)
			if !equalBooksCount(got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}

}

func equalBooksCount(got, want map[Book]uint) bool {
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}

	}

	return true
}

func booksCount(bookworms []BookWormData) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}
func TestGreetSpanish(t *testing.T) {
	var test = testCase{
		bookworkFile: "./data/bookworm.json",
		want: []BookWormData{
			{
				Name: "Fadi",
				Books: []Book{
					{
						Author: "Margaret Atwood",
						Title:  "The Handmaid's Tale",
					},
					{
						Author: "Sylvia Plath",
						Title:  "The Bell Jar",
					},
					{
						Author: "George Lucas",
						Title:  "Star Wars",
					},
				},
			},
			{
				Name: "Peggy",
				Books: []Book{
					{
						Author: "Margaret Atwood",
						Title:  "The Handmaid's Tale",
					},
					{
						Author: "Sylvia Plath",
						Title:  "The Bell Jar",
					},
				},
			},
		},
		wantErr: false,
	}

	got, err := LoadBookData(test.bookworkFile)

	if err != nil && test.wantErr {
		t.Fatalf("Expected an error, got none: %s", err.Error())
	}

	if err == nil && test.wantErr {
		t.Fatalf("Expected no error, got an error: %s, expected %s", got, test.want)
	}

	if !equalBookworms(got, test.want) {
		t.Fatal("Failed to equal results")
	}

}

func equalBookworms(bookworms, target []BookWormData) bool {
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(books, target []Book) bool {
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}
