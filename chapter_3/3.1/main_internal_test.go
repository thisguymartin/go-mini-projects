package main

import (
	"testing"
)

type testCase struct {
	bookworkFile string
	want         []BookWormData
	wantErr      bool
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
				},
			},
			{
				Name: "Peggy",
				Books: []Book{
					{
						Author: "Cat Atwood",
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
