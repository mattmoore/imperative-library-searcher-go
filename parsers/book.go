package parsers

import (
	"github.com/mattmoore/library-searcher-go-imperative/types"
	"strings"
)

func ParseBooks(data []string) []types.Book {
	books := []types.Book{}
	for _, line := range data {
		book := ParseBook(line)
		if len(book.Title) > 0 &&
			len(book.Author.FirstName) > 0 && len(book.Author.LastName) > 0 {
			books = append(books, book)
		}
	}
	return books
}

func ParseBook(data string) types.Book {
	clean := clean(data)
	return types.Book{
		Title: title(clean),
		Author: types.Author{
			FirstName: authorFirstName(clean),
			LastName:  authorLastName(clean),
		},
	}
}

func title(data string) string {
	return strings.Join(strings.Split(data, " by ")[0:1], " ")
}

func authorFirstName(data string) string {
	author := strings.Join(strings.Split(data, " by ")[1:], " ")
	authorTokens := strings.Split(author, " ")
	return strings.Join(authorTokens[:len(authorTokens)-1], " ")
}

func authorLastName(data string) string {
	// tokens := reverse(strings.Split(data, " "))
	author := strings.Join(strings.Split(data, " by ")[1:], " ")
	tokens := strings.Split(author, " ")
	return tokens[len(tokens)-1]
}

func clean(data string) string {
	tokens := strings.Split(data, " ")
	return strings.Join(tokens[:len(tokens)-1], " ")
}

func reverse(data []string) []string {
	reversed := data
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return reversed
}
