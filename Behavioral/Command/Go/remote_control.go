package main

type RemoteControl struct {
	command Command
}

func NewRemoteControl(command Command) *RemoteControl {
	return &RemoteControl{
		command: command,
	}
}

func (rc *RemoteControl) PressButton() {
	rc.command.Execute()
}
