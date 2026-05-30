package main

import "fmt"

type Light struct{}

func NewLight() *Light {
	return &Light{}
}

func (l *Light) TurnOn() {
	fmt.Println("Light is ON!")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is OFF!")
}
