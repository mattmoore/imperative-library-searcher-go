package main

import (
	"fmt"
	"github.com/mattmoore/library-searcher-go-imperative/parsers"
	"github.com/mattmoore/library-searcher-go-imperative/types"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name.")
		return
	}

	file := os.Args[1]
	content, _ := ioutil.ReadFile(file)
	data := strings.Split(string(content), "\n")
	books := parsers.ParseBooks(data[:len(data)-1])

	for _, author := range top15Authors(books) {
		fmt.Print(author.LastName)
		if len(author.FirstName) > 0 {
			fmt.Print(fmt.Sprintf(", %s", author.FirstName))
		}
		fmt.Println()
	}
}

func top15Authors(books []types.Book) []types.Author {
	authors := []types.Author{}

	for _, book := range books {
		authors = append(authors, book.Author)
	}

	sort.Sort(ByAuthorLastNameFirstName(authors))

	uniqueAuthors := []types.Author{}

	for _, author := range authors {
		if len(uniqueAuthors) >= 15 {
			break
		}
		if !contains(uniqueAuthors, author) {
			uniqueAuthors = append(uniqueAuthors, author)
		}
	}

	return uniqueAuthors
}

func contains(haystack []types.Author, needle types.Author) bool {
	for _, item := range haystack {
		if item.FirstName == needle.FirstName && item.LastName == needle.LastName {
			return true
		}
	}
	return false
}

// Author LastNameFirstName sort

type ByAuthorLastNameFirstName []types.Author

func (list ByAuthorLastNameFirstName) Len() int {
	return len(list)
}

func (list ByAuthorLastNameFirstName) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list ByAuthorLastNameFirstName) Less(i, j int) bool {
	si := list[i].LastName + list[i].FirstName
	sj := list[j].LastName + list[j].LastName
	siLower := strings.ToLower(si)
	sjLower := strings.ToLower(sj)
	if siLower == sjLower {
		return si < sj
	}
	return siLower < sjLower
}
