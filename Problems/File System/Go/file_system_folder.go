package main

func (fs *FileSystem) CreateFolder(path, name string) error {
	parentFolder, err := fs.resolveFolderPath(path)
	if err != nil {
		return err
	}
	return parentFolder.AddChild(NewFolder(name))
}

func (fs *FileSystem) GetFolder(path string) (*Folder, error) {
	return fs.resolveFolderPath(path)
}

func (fs *FileSystem) DeleteFolder(path string) error {
	parentPath, folderName, err := fs.splitLast(path)
	if err != nil {
		return err
	}

	parentFolder, err := fs.resolveFolderPath(parentPath)
	if err != nil {
		return err
	}

	return parentFolder.RemoveChild(folderName)
}

func (fs *FileSystem) RenameFolder(path, newName string) error {
	parentPath, folderName, err := fs.splitLast(path)
	if err != nil {
		return err
	}

	parentFolder, err := fs.resolveFolderPath(parentPath)
	if err != nil {
		return err
	}

	child, ok := parentFolder.GetChild(folderName)
	if !ok {
		return ErrInvalidPath
	}

	folder, ok := child.(*Folder)
	if !ok {
		return ErrInvalidPath
	}

	if _, exists := parentFolder.GetChild(newName); exists {
		return ErrFolderAlreadyExistsInDestinationFolder
	}
	if err := parentFolder.RemoveChild(folderName); err != nil {
		return err
	}

	folder.name = newName
	return parentFolder.AddChild(folder)
}

func (fs *FileSystem) MoveFolder(sourcePath, destinationFolderPath string) error {
	source, err := fs.resolveFolderPath(sourcePath)
	if err != nil {
		return err
	}
	dest, err := fs.resolveFolderPath(destinationFolderPath)
	if err != nil {
		return err
	}

	parentPath, folderName, err := fs.splitLast(sourcePath)
	if err != nil {
		return err
	}

	parentFolder, err := fs.resolveFolderPath(parentPath)
	if err != nil {
		return err
	}

	if _, exists := dest.GetChild(folderName); exists {
		return ErrFolderAlreadyExistsInDestinationFolder
	}
	if err := parentFolder.RemoveChild(folderName); err != nil {
		return err
	}

	return dest.AddChild(source)
}
