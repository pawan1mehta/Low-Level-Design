package main

type FileSystemEntity interface {
	Name() string
	Size() int64
	IsDirectory() bool
	Print(indent int)
}
