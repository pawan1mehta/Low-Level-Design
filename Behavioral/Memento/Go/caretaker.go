package main

type CareTaker struct {
	history []*Memento
}

func NewCareTaker() *CareTaker {
	return &CareTaker{
		history: []*Memento{},
	}
}

func (c *CareTaker) Save(memento *Memento) {
	c.history = append(c.history, memento)
}

func (c *CareTaker) Undo() *Memento {
	n := len(c.history)

	if n == 0 {
		return nil
	}

	c.history = c.history[:n-1]

	if len(c.history) > 0 {
		return c.history[len(c.history)-1]
	}

	return nil
}
