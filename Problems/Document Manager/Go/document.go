package main

import "time"

type AccessLevel int

const (
	AccessNone AccessLevel = iota
	EDIT
	VIEW
)

type Document struct {
	id      string
	content string
	owner   User
	shares  map[string]AccessLevel
	history *VersionHistory
}

func NewDocument(id string, owner User) *Document {
	d := &Document{
		id:      id,
		content: "",
		owner:   owner,
		shares:  map[string]AccessLevel{},
		history: NewVersionHistory(),
	}
	d.history.Save(d.createSnapshot(owner))
	return d
}

func (d *Document) Update(user User, content string) error {
	if !d.canUpdate(user) {
		return ErrAccessDenied
	}

	d.content = content

	snapshot := d.createSnapshot(user)
	d.history.Save(snapshot)

	return nil
}

func (d *Document) ReadContent(user User) (string, error) {
	if !d.canRead(user) {
		return "", ErrAccessDenied
	}
	return d.content, nil
}

func (d *Document) Restore(user User, version int) error {
	if !d.canUpdate(user) {
		return ErrAccessDenied
	}
	ver, err := d.history.Get(version)
	if err != nil {
		return err
	}

	d.content = ver.getContent()
	d.history.Save(d.createSnapshot(user))

	return nil
}

func (d *Document) CanDelete(user User) bool {
	if !d.IsOwner(user) {
		return false
	}
	return true
}

func (d *Document) Share(owner User, targetUser User, accessLevel AccessLevel) error {
	if !d.IsOwner(owner) {
		return ErrNotOwner
	}
	d.shares[targetUser.ID()] = accessLevel
	return nil
}

func (d *Document) Versions(user User) ([]Version, error) {
	if !d.canRead(user) {
		return nil, ErrAccessDenied
	}
	return d.history.List(), nil
}

func (d *Document) createSnapshot(user User) Version {
	return NewVersion(d.history.NextID()+1, d.content, user.ID(), time.Now())
}

func (d *Document) IsOwner(user User) bool {
	if d.owner.ID() == user.ID() {
		return true
	}
	return false
}

func (d *Document) canUpdate(user User) bool {
	if d.IsOwner(user) {
		return true
	}
	if accessLevel, ok := d.shares[user.ID()]; ok && accessLevel == EDIT {
		return true
	}
	return false
}

func (d *Document) canRead(user User) bool {
	if d.IsOwner(user) {
		return true
	}
	if accessLevel, ok := d.shares[user.ID()]; ok && (accessLevel == VIEW || accessLevel == EDIT) {
		return true
	}
	return false
}
