package main

import "fmt"

func greet() string {
	return "Hello World!!"
}

func main() {
	greet := greet()
	fmt.Println(greet)
}
