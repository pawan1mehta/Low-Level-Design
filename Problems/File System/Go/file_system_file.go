package main

func (fs *FileSystem) CreateFile(path, name, content string) error {
	parentFolder, err := fs.resolveFolderPath(path)
	if err != nil {
		return err
	}
	return parentFolder.AddChild(NewFile(name, content))
}

func (fs *FileSystem) GetFile(path string) (*File, error) {
	file, err := fs.resolveFilePath(path)
	if err != nil {
		return nil, ErrFileNotFound
	}
	return file, nil
}

func (fs *FileSystem) DeleteFile(path string) error {
	parentPath, fileName, err := fs.splitLast(path)
	if err != nil {
		return err
	}

	parentFolder, err := fs.resolveFolderPath(parentPath)
	if err != nil {
		return err
	}

	return parentFolder.RemoveChild(fileName)
}

func (fs *FileSystem) ReadFile(path string) (string, error) {
	file, err := fs.GetFile(path)
	if err != nil {
		return "", err
	}
	return file.Read(), nil
}

func (fs *FileSystem) UpdateFile(path, content string) error {
	file, err := fs.resolveFilePath(path)
	if err != nil {
		return err
	}
	return file.Update(content)
}

func (fs *FileSystem) RenameFile(path, newName string) error {
	parentPath, fileName, err := fs.splitLast(path)
	if err != nil {
		return err
	}

	parentFolder, err := fs.resolveFolderPath(parentPath)
	if err != nil {
		return err
	}

	child, ok := parentFolder.GetChild(fileName)
	if !ok {
		return ErrInvalidPath
	}

	file, ok := child.(*File)
	if !ok {
		return ErrInvalidPath
	}

	if _, exists := parentFolder.GetChild(newName); exists {
		return ErrFileAlreadyExistsInDestinationFolder
	}
	if err := parentFolder.RemoveChild(fileName); err != nil {
		return err
	}

	file.name = newName
	return parentFolder.AddChild(file)
}

func (fs *FileSystem) MoveFile(sourcePath, destinationFolder string) error {
	source, err := fs.resolveFilePath(sourcePath)
	if err != nil {
		return err
	}
	dest, err := fs.resolveFolderPath(destinationFolder)
	if err != nil {
		return err
	}

	parentPath, fileName, err := fs.splitLast(sourcePath)
	if err != nil {
		return err
	}

	parentFolder, err := fs.resolveFolderPath(parentPath)
	if err != nil {
		return err
	}

	if _, exists := dest.GetChild(fileName); exists {
		return ErrFileAlreadyExistsInDestinationFolder
	}
	if err := parentFolder.RemoveChild(fileName); err != nil {
		return err
	}

	return dest.AddChild(source)
}
