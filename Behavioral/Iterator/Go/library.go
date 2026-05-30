package main

type Library struct {
	books []*Book
}

func NewLibrary() *Library {
	return &Library{
		books: []*Book{},
	}
}

func (l *Library) Add(book *Book) {
	l.books = append(l.books, book)
}

func (l *Library) CreateIterator() BookIterator {
	return NewLibraryBookIterator(l.books)
}
