package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {

	// 1. Your software should read all data from the given CSV files in a meaningful structure.
	publications := readAllData()

	// 2. Print out all books and magazines (could be a GUI, console, …) with all their details (with a meaningful output format).
	fmt.Println("2. Print out all books and magazines")
	printAllBooksAndMagazines(publications) // missing Author details...

	// 3. Find a book or magazine by its `isbn`.
	isbnToBeFound := "5554-5545-4518"
	p1 := findPublicationByISBN(publications, isbnToBeFound)
	fmt.Printf("\n3. Book or magazine with ISBN '%s': %v\n", isbnToBeFound, p1.title)

	// 4. Find all books and magazines by their `authors`’ email.
	authorToBeFound := "null-walter@echocat.org"
	p2 := findPublicationsByAuthorEmail(publications, authorToBeFound)
	fmt.Printf("\n4. Books or magazines with Author '%s':\n", authorToBeFound)
	for _, pp2 := range p2 {
		fmt.Printf("\t%s, authors '%s'\n", pp2.title, pp2.authors)
	}

	// 5. Print out all books and magazines with all their details sorted by `title`. This sort should be done for books and magazines together.
	fmt.Println("\n5. Print out all books and magazines sorted by title")
	sort.Slice(publications, func(a, b int) bool {
		return publications[a].title < publications[b].title
	})
	printAllBooksAndMagazines(publications) // missing Author details...
}

func findPublicationByISBN(publications []Publication, search string) Publication {
	var p Publication
	for _, pp := range publications {
		if pp.isbn == search {
			p = pp
		}
	}

	return p
}

func findPublicationsByAuthorEmail(publications []Publication, search string) []Publication {
	var p []Publication

	for _, pp := range publications {
		if slices.Contains(pp.authors, search) {
			p = append(p, pp)
		}
	}

	return p
}

func readAllData() []Publication {

	// debug
	// fmt.Println(readAuthors())
	// fmt.Println(readBooks())
	// fmt.Println(readMagazines())

	return preparePublications(readBooks(), readMagazines())
}

func printAllBooksAndMagazines(publications []Publication) {

	for i, p := range publications {
		notes := p.otherNotes
		if len(notes) > 25 {
			notes = notes[0:25] + "..."
		}
		fmt.Printf("\t%d) Title '%s', ISBN '%s', Authors: '%s', other notes: '%s'\n", i+1, p.title, p.isbn, p.authors, notes)
	}
}

func readFromCsv(filename string) [][]string {
	// Open the CSV
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file")
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Read from file
	lines, err := reader.ReadAll()
	if err != nil {
		panic("cannot read file")
	}

	return lines
}

// Unused...
func readAuthors() []Author {

	lines := readFromCsv("resources/authors.csv")

	authors := []Author{}
	for i, line := range lines {
		// fmt.Println(line)
		if i > 0 {
			authors = append(authors, Author{line[0], line[1], line[2]})
		}
	}

	return authors
}

func readBooks() []Book {

	lines := readFromCsv("resources/books.csv")

	books := []Book{}
	for i, line := range lines {
		// fmt.Println(line)
		if i > 0 {
			books = append(books, Book{line[0], line[1], line[2], line[3]})
		}
	}

	return books
}

func preparePublications(books []Book, magazines []Magazine) []Publication {

	p := []Publication{}

	for _, b := range books {
		p = append(p, Publication{b.title, b.isbn, strings.Split(b.authors, ","), b.description})
	}

	for _, m := range magazines {
		p = append(p, Publication{m.title, m.isbn, strings.Split(m.authors, ","), m.publishedAt})
	}

	return p
}

func readMagazines() []Magazine {

	lines := readFromCsv("resources/magazines.csv")

	magazines := []Magazine{}
	for i, line := range lines {
		// fmt.Println(line)
		if i > 0 {
			magazines = append(magazines, Magazine{line[0], line[1], line[2], line[3]})
		}
	}

	return magazines
}

type Author struct {
	email     string
	firstname string
	lastname  string
}

type Book struct {
	title       string
	isbn        string
	authors     string
	description string
}

type Magazine struct {
	title       string
	isbn        string
	authors     string
	publishedAt string
}

type Publication struct {
	title      string
	isbn       string
	authors    []string
	otherNotes string
}
