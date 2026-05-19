package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing a Rectangle")
}

// Option 1: Static Factory Method
func NewShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{}
	case "rectangle":
		return &Rectangle{}
	default:
		return nil
	}
}

// Option 2: Factory Method
type ShapeFactory interface {
	CreateShape() Shape
}

type CircleFactory struct{}

func (cf *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

type RectangleFactory struct{}

func (rf *RectangleFactory) CreateShape() Shape {
	return &Rectangle{}
}

func main() {
	shape1 := NewShape("circle")
	shape1.Draw()

	var shapeFactory ShapeFactory

	shapeFactory = &RectangleFactory{}
	shapeFactory.CreateShape().Draw()
}
