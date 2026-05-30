package main

type Command interface {
	Execute()
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (lc *LightOnCommand) Execute() {
	lc.light.TurnOn()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

func (lc *LightOffCommand) Execute() {
	lc.light.TurnOff()
}
