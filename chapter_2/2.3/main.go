package main

import "fmt"

type language string

var pharsebook = map[language]string{
	"sp": "Hola Mundo!",
	"en": "Hello World!",
}

func greet(l language) string {
	greeting, ok := pharsebook[l]
	if !ok {
		return fmt.Sprintf("Nothing found!: %q", l)
	}
	return greeting
}

func main() {
	greet := greet("en")
	fmt.Println(greet)
}
