package main

import (
	"testing"
)

func TestFindByISBN(t *testing.T) {

	// given
	p := []Publication{
		{title: "1984", isbn: "YYY", authors: []string{"A"}, otherNotes: ""},
		{title: "Cento anni di solitudine", isbn: "xxx", authors: []string{"B"}, otherNotes: ""},
		{title: "Il sentiero dei nidi di ragno", isbn: "ZzZ", authors: []string{"C"}, otherNotes: ""},
	}
	expected := "Cento anni di solitudine"

	// than
	actual := findPublicationByISBN(p, "xxx")

	// that
	if actual.title != expected {
		t.Fatalf(`TestFindByISBN fails, expected %s but was %s`, expected, actual.title)
	}
}
