package main

import "errors"

var (
	ErrDocumentExists   = errors.New("document already exists")
	ErrDocumentNotFound = errors.New("document not found")
	ErrAccessDenied     = errors.New("access denied")
	ErrVersionNotFound  = errors.New("version not found")
	ErrNotOwner         = errors.New("only the owner can share the document")
)
