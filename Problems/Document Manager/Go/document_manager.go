package main

import "fmt"

type DocumentManager struct {
	documentMap map[string]*Document
}

func NewDocumentManager() *DocumentManager {
	return &DocumentManager{
		documentMap: make(map[string]*Document),
	}
}

func (dm *DocumentManager) CreateDocument(id string, user User) (string, error) {
	if _, exists := dm.documentMap[id]; exists {
		return "", ErrDocumentExists
	}

	document := NewDocument(id, user)
	dm.documentMap[id] = document

	return id, nil
}

func (dm *DocumentManager) UpdateDocument(id string, user User, content string) error {
	doc, exists := dm.documentMap[id]
	if !exists {
		return ErrDocumentNotFound
	}

	return doc.Update(user, content)
}

func (dm *DocumentManager) Read(id string, user User) (string, error) {
	if _, exists := dm.documentMap[id]; !exists {
		return "", ErrDocumentNotFound
	}

	content, err := dm.documentMap[id].ReadContent(user)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (dm *DocumentManager) Restore(id string, user User, version int) error {
	if _, exists := dm.documentMap[id]; !exists {
		return ErrDocumentNotFound
	}

	if err := dm.documentMap[id].Restore(user, version); err != nil {
		return err
	}
	fmt.Printf("Version %d restored successfully!", version)

	return nil
}

func (dm *DocumentManager) DeleteDocument(id string, user User) error {
	if _, exists := dm.documentMap[id]; !exists {
		return ErrDocumentNotFound
	}

	doc, _ := dm.documentMap[id]
	if !doc.CanDelete(user) {
		return ErrAccessDenied
	}

	delete(dm.documentMap, id)

	return nil
}

func (dm *DocumentManager) Share(id string, owner User, targetUser User, accessLevel AccessLevel) error {
	if _, exists := dm.documentMap[id]; !exists {
		return ErrDocumentNotFound
	}

	doc, _ := dm.documentMap[id]
	if !doc.IsOwner(owner) {
		return ErrNotOwner
	}

	if err := doc.Share(owner, targetUser, accessLevel); err != nil {
		return err
	}

	return nil
}

func (dm *DocumentManager) ListVersion(id string, user User) ([]Version, error) {
	if _, exists := dm.documentMap[id]; !exists {
		return nil, ErrDocumentNotFound
	}

	versions, err := dm.documentMap[id].Versions(user)
	if err != nil {
		return nil, err
	}

	return versions, nil
}
