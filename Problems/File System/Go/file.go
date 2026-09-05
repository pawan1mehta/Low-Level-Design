package main

import (
	"fmt"
	"strings"
)

type File struct {
	name    string
	content string
}

func NewFile(name, content string) *File {
	return &File{
		name:    name,
		content: content,
	}
}

func (file *File) Name() string {
	return file.name
}

func (file *File) Size() int64 {
	return int64(len(file.content))
}

func (file *File) IsDirectory() bool {
	return false
}

func (file *File) Print(indent int) {
	fmt.Println(strings.Repeat(" ", indent) + file.name)
}

func (file *File) Read() string {
	return file.content
}

func (file *File) Update(content string) error {
	file.content = content
	return nil
}

func (file *File) Rename(newName string) error {
	file.name = newName
	return nil
}
