package main

import "fmt"

type FileComponent interface {
	GetName() string
	GetSize() int
	Display()
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Display() {
	fmt.Printf("File: %s (%d KB)\n", f.name, f.size)
}

type Folder struct {
	name     string
	children []FileComponent
}

func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		children: []FileComponent{},
	}
}

func (f *Folder) Add(component FileComponent) {
	f.children = append(f.children, component)
}

func (f *Folder) Remove(component FileComponent) {
	for i, child := range f.children {
		if child == component {
			f.children = append(f.children[:i], f.children[i+1:]...)
		}
	}
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) GetSize() int {
	total := 0
	for _, child := range f.children {
		total += child.GetSize()
	}
	return total
}

func (f *Folder) Display() {
	fmt.Printf("Folder: %s (%d KB)\n", f.name, f.GetSize())
	for _, child := range f.children {
		child.Display()
	}
}
