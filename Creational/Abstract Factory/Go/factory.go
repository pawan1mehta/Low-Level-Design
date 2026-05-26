package main

type GUIFactory interface {
	createButton() Button
	createCheckbox() Checkbox
}

type MacOSFactory struct{}

func (f MacOSFactory) createButton() Button {
	return MacOSButton{}
}

func (f MacOSFactory) createCheckbox() Checkbox {
	return MacOSCheckbox{}
}

type WindowsFactory struct{}

func (f WindowsFactory) createButton() Button {
	return WindowsOSButton{}
}

func (f WindowsFactory) createCheckbox() Checkbox {
	return WindowsOSCheckbox{}
}
