package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	expect := "Hello World!!"
	greet := greet()

	if greet != expect {
		t.Errorf("expect %q, got: %q", expect, greet)
	}

}
