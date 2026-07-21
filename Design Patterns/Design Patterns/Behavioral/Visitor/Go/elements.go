package main

type Element interface {
	Accept(visitor Visitor)
}

type Document struct {
	elements []Element
}

func NewDocument(elements []Element) *Document {
	return &Document{
		elements: elements,
	}
}

func (d *Document) Accept(visitor Visitor) {
	visitor.VisitDocument(d)
}

func (d *Document) Elements() []Element {
	return d.elements
}

type Heading struct {
	text string
}

func NewHeading(text string) *Heading {
	return &Heading{
		text: text,
	}
}

func (d *Heading) Accept(visitor Visitor) {
	visitor.VisitHeading(d)
}

func (d *Heading) GetText() string {
	return d.text
}

type Paragraph struct {
	text string
}

func NewParagraph(text string) *Paragraph {
	return &Paragraph{
		text: text,
	}
}

func (p *Paragraph) Accept(visitor Visitor) {
	visitor.VisitParagraph(p)
}

func (p *Paragraph) GetText() string {
	return p.text
}
