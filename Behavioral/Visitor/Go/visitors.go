package main

import "strings"

type Visitor interface {
	VisitDocument(document *Document)
	VisitHeading(heading *Heading)
	VisitParagraph(paragraph *Paragraph)
}

type HTMLExportVisitor struct {
	html strings.Builder
}

func NewHTMLExportVisitor() *HTMLExportVisitor {
	return &HTMLExportVisitor{}
}

func (h *HTMLExportVisitor) VisitDocument(document *Document) {
	h.html.WriteString("<html><body>")
	for _, el := range document.Elements() {
		el.Accept(h)
	}
	h.html.WriteString("</body></html>")
}

func (h *HTMLExportVisitor) VisitHeading(heading *Heading) {
	h.html.WriteString("<h1>")
	h.html.Write([]byte(heading.GetText()))
	h.html.WriteString("</h1>")
}

func (h *HTMLExportVisitor) VisitParagraph(paragraph *Paragraph) {
	h.html.WriteString("<p>")
	h.html.Write([]byte(paragraph.GetText()))
	h.html.WriteString("</p>")
}

func (h *HTMLExportVisitor) GetHtml() string {
	return h.html.String()
}
