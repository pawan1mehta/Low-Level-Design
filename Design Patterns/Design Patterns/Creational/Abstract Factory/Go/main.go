package main

import "fmt"

func main() {
	osName := "MAC"

	var factory GUIFactory

	switch osName {
	case "MAC":
		factory = MacOSFactory{}
	case "WINDOWS":
		factory = WindowsFactory{}
	default:
		fmt.Println("Error! Unknown operating system.")
		return
	}

	app := NewApplication(factory)
	app.Paint()
}
