package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The langue is required")
	flag.Parse()

	fmt.Printf("Key %q", lang)
}
