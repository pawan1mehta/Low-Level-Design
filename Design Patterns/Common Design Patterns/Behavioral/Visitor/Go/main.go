package main

import "fmt"

func main() {
	var elements []Element

	elements = append(elements, NewHeading("first heading"))
	elements = append(elements, NewHeading("second heading"))
	elements = append(elements, NewParagraph("Paragraph"))

	htmlExportVisitor := NewHTMLExportVisitor()

	document := NewDocument(elements)
	document.Accept(htmlExportVisitor)

	fmt.Println("html: ", htmlExportVisitor.GetHtml())
}
