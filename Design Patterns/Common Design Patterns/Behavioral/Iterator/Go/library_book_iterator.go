package main

import (
	"log"
)

type LibraryBookIterator struct {
	books []*Book
	pos   int
}

func NewLibraryBookIterator(books []*Book) *LibraryBookIterator {
	return &LibraryBookIterator{
		books: books,
		pos:   0,
	}
}

func (l *LibraryBookIterator) HashNext() bool {
	if l.pos < len(l.books) {
		return true
	} else {
		return false
	}
}

func (l *LibraryBookIterator) Next() *Book {
	if !l.HashNext() {
		log.Fatalln("No more books found")
		return &Book{}
	} else {
		res := l.books[l.pos]
		l.pos++
		return res
	}
}
