package main

type Movie struct {
	id          string
	title       string
	description string
}

func (m Movie) ID() string {
	return m.id
}

func (m Movie) Title() string {
	return m.title
}
