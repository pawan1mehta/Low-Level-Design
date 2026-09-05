package main

import (
	"strings"
	"sync"
)

type FileSystem struct {
	root *Folder
}

var (
	fileSystem *FileSystem
	once       sync.Once
)

func NewFileSystem() *FileSystem {
	once.Do(func() {
		fileSystem = &FileSystem{
			root: NewFolder("/"),
		}
	})
	return fileSystem
}

func (fs *FileSystem) PrintTree() {
	fs.root.Print(0)
}

func (fs *FileSystem) Navigate(path string) (FileSystemEntity, error) {
	folder, err := fs.resolveFolderPath(path)
	if err == nil {
		return folder, nil
	}
	return fs.resolveFilePath(path)
}

func (fs *FileSystem) resolveFilePath(path string) (*File, error) {
	parentPath, name, err := fs.splitLast(path)
	if err != nil {
		return nil, err
	}

	folder, err := fs.resolveFolderPath(parentPath)
	if err != nil {
		return nil, err
	}

	child, ok := folder.GetChild(name)
	if !ok {
		return nil, ErrInvalidPath
	}

	file, ok := child.(*File)
	if !ok {
		return nil, ErrInvalidPath
	}

	return file, nil
}

func (fs *FileSystem) resolveFolderPath(path string) (*Folder, error) {
	folders := strings.Split(path, "/")

	currFolder := fs.root

	for _, folderName := range folders {
		if folderName == "" {
			continue
		}

		child, ok := currFolder.GetChild(folderName)
		if !ok {
			return nil, ErrInvalidPath
		}
		if !child.IsDirectory() {
			return nil, ErrInvalidPath
		}

		folder, ok := child.(*Folder)
		if !ok {
			return nil, ErrInvalidPath
		}

		currFolder = folder
	}

	return currFolder, nil
}

func (fs *FileSystem) splitLast(path string) (parentPath string, name string, err error) {
	if path == "" || path[0] != '/' || path == "/" {
		return "", "", ErrInvalidPath
	}

	idx := strings.LastIndex(path, "/")
	parent, name := path[:idx], path[idx+1:]
	if name == "" {
		return "", "", ErrInvalidPath
	}
	if parent == "" {
		parent = "/"
	}

	return parent, name, nil
}
