package main

type BookCollection interface {
	CreateIterator() BookIterator
}
