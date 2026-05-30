package main

type Projector struct{}

func NewProjector() *Projector {
	return &Projector{}
}

func (ss *Projector) On() {
	println("Projector is on")
}

func (ss *Projector) Off() {
	println("Projector is off")
}

func (ss *Projector) SetInput(input string) {
	println("Setting projector input to", input)
}
