package main

import (
	"testing"
)

func TestGreetSpanish(t *testing.T) {
	expect := "Hola Mundo!"
	greet := greet("sp")

	if greet != expect {
		t.Errorf("expect %q, got: %q", expect, greet)
	}
}

func TestGreetEnglish(t *testing.T) {
	expect := "Hello World!"
	greet := greet("en")

	if greet != expect {
		t.Errorf("expect %q, got: %q", expect, greet)
	}
}

func TestGreetNotFound(t *testing.T) {
	expect := "Nothing found!: \"1en\""
	greet := greet("1en")
	if greet != expect {
		t.Errorf("expect %q, got: %q", expect, greet)
	}
}
