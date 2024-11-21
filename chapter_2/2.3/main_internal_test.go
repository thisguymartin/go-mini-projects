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

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello World!",
		},
		"Spanish": {
			lang: "sp",
			want: "Hola Mundo!",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `Nothing found!: "akk"`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("Expected %q, got: %q", tc.want, got)
			}
		})
	}

}
