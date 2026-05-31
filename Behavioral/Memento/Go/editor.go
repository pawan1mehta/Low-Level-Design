package main

type Editor struct {
	content string
}

func NewEditor() *Editor {
	return &Editor{
		content: "",
	}
}

func (e *Editor) Type(word string) {
	e.content += word
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) Save() *Memento {
	return NewMemento(e.content)
}

func (e *Editor) Restore(memento *Memento) {
	e.content = memento.GetState()
}
