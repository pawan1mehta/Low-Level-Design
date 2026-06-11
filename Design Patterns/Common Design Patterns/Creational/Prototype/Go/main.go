package main

import "fmt"

type Clonable interface {
	Clone() Clonable
}

type Document struct {
	Title   string
	Content string
	Author  string
}

func (d *Document) Clone() Clonable {
	return &Document{
		Title:   d.Title,
		Content: d.Content,
		Author:  d.Author,
	}
}

type Report struct {
	Title       string
	Description string
	Department  string
}

func (r *Report) Clone() Clonable {
	return &Report{
		Title:       r.Title,
		Description: r.Description,
		Department:  r.Department,
	}
}

func main() {
	docObj := &Document{
		Title:   "Template Doc",
		Content: "Sample content",
		Author:  "Admin",
	}

	clonedDocObj := docObj.Clone().(*Document)
	clonedDocObj.Title = "My Document"

	fmt.Println("Original Document: ", docObj)
	fmt.Println("Cloned Document: ", clonedDocObj)
}
