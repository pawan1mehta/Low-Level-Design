package main

import "time"

type Version struct {
	id        int
	content   string
	userID    string
	createdAt time.Time
}

func NewVersion(id int, content string, userID string, createdAt time.Time) Version {
	return Version{
		id:        id,
		content:   content,
		userID:    userID,
		createdAt: createdAt,
	}
}

func (v Version) GetID() int {
	return v.id
}

func (v Version) getContent() string {
	return v.content
}

type VersionHistory struct {
	versions []Version
}

func NewVersionHistory() *VersionHistory {
	return &VersionHistory{
		versions: []Version{},
	}
}

func (vh *VersionHistory) Save(version Version) {
	vh.versions = append(vh.versions, version)
}

func (vh *VersionHistory) Get(id int) (Version, error) {
	for _, version := range vh.versions {
		if version.GetID() == id {
			return version, nil
		}
	}
	return Version{}, ErrVersionNotFound
}

func (vh *VersionHistory) List() []Version {
	var versionList []Version
	for _, version := range vh.versions {
		versionList = append(versionList, version)
	}
	return versionList
}

func (vh *VersionHistory) NextID() int {
	return len(vh.versions) + 1
}
