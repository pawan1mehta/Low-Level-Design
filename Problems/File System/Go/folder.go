package main

import (
	"fmt"
	"strings"
)

type Folder struct {
	name     string
	children map[string]FileSystemEntity
}

func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		children: map[string]FileSystemEntity{},
	}
}

func (folder *Folder) Name() string {
	return folder.name
}

func (folder *Folder) Size() int64 {
	var total int64
	for _, child := range folder.children {
		total += child.Size()
	}
	return total
}

func (folder *Folder) IsDirectory() bool {
	return true
}

func (folder *Folder) Print(indent int) {
	label := folder.name
	if folder.name == "" || folder.name == "/" {
		label = "/"
	}
	fmt.Println(strings.Repeat("  ", indent) + label)
	for _, child := range folder.children {
		child.Print(indent + 1)
	}
}

func (folder *Folder) GetChild(name string) (FileSystemEntity, bool) {
	child, ok := folder.children[name]
	return child, ok
}

func (folder *Folder) AddChild(child FileSystemEntity) error {
	name := child.Name()
	if _, exists := folder.children[name]; exists {
		return ErrAlreadyExists
	}
	folder.children[name] = child
	return nil
}

func (folder *Folder) RemoveChild(name string) error {
	if _, ok := folder.children[name]; !ok {
		return ErrEntityNotFound
	}
	delete(folder.children, name)
	return nil
}

func (folder *Folder) Rename(newName string) error {
	folder.name = newName
	return nil
}
