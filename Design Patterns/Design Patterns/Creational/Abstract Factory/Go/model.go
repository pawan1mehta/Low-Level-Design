package main

type Button interface {
	Paint()
}

type MacOSButton struct{}

func (b MacOSButton) Paint() {
	println("Rendering a button in MacOS style.")
}

type WindowsOSButton struct{}

func (b WindowsOSButton) Paint() {
	println("Rendering a button in Windows style.")
}

type Checkbox interface {
	Paint()
}

type MacOSCheckbox struct{}

func (c MacOSCheckbox) Paint() {
	println("Rendering a checkbox in MacOS style.")
}

type WindowsOSCheckbox struct{}

func (c WindowsOSCheckbox) Paint() {
	println("Rendering a checkbox in Windows style.")
}
