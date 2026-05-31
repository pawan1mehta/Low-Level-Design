package main

import "fmt"

func main() {
	careTaker := NewCareTaker()
	editor := NewEditor()

	editor.Type("Pawan ")
	careTaker.Save(editor.Save())

	fmt.Println("Editor Content: ", editor.GetContent())

	editor.Type("Kumar")
	careTaker.Save(editor.Save())

	fmt.Println("Editor Content: ", editor.GetContent())

	editor.Restore(careTaker.Undo())

	fmt.Println("Editor Content: ", editor.GetContent())
}
