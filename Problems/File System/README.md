# File System

Link: <https://www.hellointerview.com/learn/low-level-design/problem-breakdowns/file-system>

## Clarifying Questions

- Hierarchy? Single root like Unix system or multiple root like Windows? [And: single root]
- Operation? [Ans: create, read, update, delete, move, rename & navigate]

## Requirements

- Hierarchical file system with single root directory
- Files store string content
- Folder contains files and other folders
- Names are unique within a directory
- Perform CRUD operations on directories/files
- Rename & Remove the file/directories
- Navigate to absolute paths

Out Of Scope:

- UI
- Search
- RBAC (Role based access control)
- Symbolic Links
- Persistence Disks storage

## Entities & Relationships

- FileSystem
- Folder
- File

Relationships:

```code
    Abstract: FileSystemEntity
        - Folder
        - File

    FileSystem <|----- composed-of -------- rooot: Folder

    Folder <|----- composed-of ----- children: []FileSystemEntity
    
    File <----- has-a ------ content: string
```

## Class Design

```code
Class FileSystem:

    - root: Folder

    - resolve(path: string) -> FileSystemEntity
    - resolveFolder(path: string) -> Folder

    + createFolder(path: string, name: string)
    + createFile(path: string, name: string, content: string)

    + readFile(path: string) -> string
    + updateFile(path: string, content: string)
    + list(path: string) -> List<String>
    + navigate(path: string) -> FileSystemEntity

    + deleteFolder(path: string)
    + deleteFile(path: string)

    + renameFolder(path: string, newName: string)  
    + renameFile(path: string, newName: string)

    + moveFolder(sourcePath: string, destinationFolderPath: string)
    + moveFile(sourcePath: string, destinationFolderPath: string)
```

```code
Abstract FileSystemEntity:

    + getName() -> string
    + getSize() -> int
    + isDirectory() -> bool
```

```code
class Folder:

    - name: string
    - children: Map<String, FileSystemEntity>

    + getName() -> string
    + getSize() -> int
    + isDirectory

    + addChild(fileSystemEntity: FileSystemEntity)
    + getChild(name: string) -> FileSystemEntity
    + removeChild(name: string)
    + listChild() -> List<FileSystemEntity>
```

```code
class File:

    - name: string
    - content: string

    + getName() -> string
    + getSize() -> int
    + isDirectory() -> bool

    + update(content: string)
    + read() -> string
```
