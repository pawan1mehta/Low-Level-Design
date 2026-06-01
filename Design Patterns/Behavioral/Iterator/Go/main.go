package main

import "fmt"

func main() {
	library := NewLibrary()

	book1 := NewBook("THE FOUNDATION", "AYN RAND", "content")
	book2 := NewBook("Thinking Fast & Slow", "Don't Know", "content")

	library.Add(book1)
	library.Add(book2)

	libraryBookIterator := library.CreateIterator()

	for libraryBookIterator.HashNext() {
		book := libraryBookIterator.Next()
		fmt.Println("Book: ", book.Title)
	}
}
