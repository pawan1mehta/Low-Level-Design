# Document Manager

Link: <https://leetcode.com/discuss/post/4589653/phonepe-machine-coding-round-2-yoe-by-an-67om/>

## Problem Statement

A document manager is an application that can be used to create, share and manage text documents. It provides an interface to easily save and retrieve documents.

- Document to be a page with only text data.
- History of changes to a document to be maintained.
- A document can be reverted to a previous version.
- Only the user who created the document can edit, revert or delete the document. Anyone can view the document.

## Requirements

- User can create a text document.
  - creater is the owner and has full access

- Owner can share a document with another user as VIEW and EDIT
  - VIEW: read content & list versions
  - EDIT: VIEW + update content and restore a previous version

- Anyone with VIEW or EDIT can read the document

- Owner or EDIT can restore version N

- Only the owner can delete the document

- Reject: missing doc/user/version, share by non-owner, no access.

Out of Scope:

- UI
- Persistence of docuement
- Real-time collaboration
- Comments on document
- Notification

## Entities & Relationship

- DocumentManager
- Document
- Version
- User
- AccessLevel
- VersionManager

Relationships:

```code
  DocumentManager <---- composed of--- Document

  Document <---- contains --- user & AccessLevel
  Document <---- contains ---- VersionManager

  VersionManager <---- containts ---- Version
```

## Design Class

```code
Class DocumentManager:

  - documentsMap : Map<ID, Document>

  + createDocument(user: User)-> Document
  + getDocument(id) -> Document
  + readDocument(id) -> string
  + updateDocument(id: string, user: User, content: string)
  + restore(id: string, user: User, version: id)
  + deleteDocument(id: string, user: User)
  + shareDocument(id: string, owner: User, user: User, acessLevel: AcessLevel)
```

```code
Class Document:

  - id: string
  - content: string
  - owner: User
  - shares: Map<Id, AcessLevel>
  - versionManbager: VersionHistory

  + update(content: string, user: User)
  + read(user: User)
  + restore(user: User, version: int)
  + share(user: User, userID: string, accessLevel: AcessLevel)
  + listVersion() -> List<Version>
  - canView(user: User)
  - canEdit(user: User)
```

```code
Class Version:  
  - id: int
  - authorId: string
  - content: string
  - createdAt: time

  + getState() -> string
```

```code
Class VersionHistory:

  - versions: []Version
  - current: int

  save(Version)
  get(id: int) -> Version
  list() -> List<Version>
```

```code
Class User:

  - id: string
  - name: string

  + getID() -> string
  + getName() -> string
```
