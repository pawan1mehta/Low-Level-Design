package main

type Book struct {
	Title   string
	Author  string
	Content string
}

func NewBook(title, author, content string) *Book {
	return &Book{
		Title:   title,
		Author:  author,
		Content: content,
	}
}
