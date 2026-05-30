package main

type BookIterator interface {
	HashNext() bool
	Next() *Book
}
