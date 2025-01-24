package main

type bookRecommendations map[Book]bookCollection

type bookCollection map[Book]struct{}

func newCollection() bookCollection {
	return make(bookCollection)
}
