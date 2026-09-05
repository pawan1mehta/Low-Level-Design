package main

import "errors"

var (
	ErrInvalidPath                            = errors.New("invalid path")
	ErrFolderAlreadyExists                    = errors.New("folder already exists")
	ErrEntityNotFound                         = errors.New("child entity not found")
	ErrFileNotFound                           = errors.New("file not found")
	ErrFileAlreadyExistsInDestinationFolder   = errors.New("file already exists in destination folder")
	ErrFolderAlreadyExistsInDestinationFolder = errors.New("folder already exists in destination folder")
	ErrAlreadyExists                          = errors.New("name already exists")
)
