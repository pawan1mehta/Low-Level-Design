# Document Manager

Link: <https://leetcode.com/discuss/post/4589653/phonepe-machine-coding-round-2-yoe-by-an-67om/>

## Problem Statement

A document manager is an application that can be used to create, share and manage text documents. It provides an interface to easily save and retrieve documents.

- Document to be a page with only text data.
- History of changes to a document to be maintained.
- A document can be reverted to a previous version.
- Only the user who created the document can edit, revert or delete the document. Anyone can view the document.

## Requirements

- User should be able to create a document
  - Creater is the owner
- Owner (a users with EDIT) can update the document
  - Each update should be saved in the history
- User should be able to share a docuement with other user  as VIEW or EDIT
- User with access can list versions and restore a previous version

Out of Scope:

- UI
- Persistence of docuement
- Real-time collaboration
- Comments on document
- Notification

## Entities & Relationship

- Document
- DocumentManager
- User

Relationships:

DocumentManager <---- composed of--- Document
