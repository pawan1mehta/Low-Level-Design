package main

type Memento struct {
	state string
}

func NewMemento(content string) *Memento {
	return &Memento{
		state: content,
	}
}

func (m *Memento) GetState() string {
	return m.state
}
