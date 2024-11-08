package main

import "fmt"

type language string

func greet(l language) string {
	switch l {
	case "en":
		return "Hello World!"
	case "sp":
		return "Hola Mundo!"
	default:
		return ""
	}

}

func main() {
	greet := greet("sp")
	fmt.Println(greet)
}
