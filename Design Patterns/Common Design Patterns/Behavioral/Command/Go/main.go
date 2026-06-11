package main

func main() {
	light := NewLight()

	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	remoteControl := NewRemoteControl(lightOnCommand)
	remoteControl.PressButton()

	remoteControl = NewRemoteControl(lightOffCommand)
	remoteControl.PressButton()
}
