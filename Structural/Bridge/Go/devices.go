package main

import "fmt"

type Device interface {
	TurnOn()

	TurnOff()

	SetVolume(volume int)
}

type TV struct {
	volume int
}

func NewTV() *TV {
	return &TV{
		volume: 10,
	}
}

func (tv *TV) TurnOn() {
	fmt.Println("TV turned ON!")
}

func (tv *TV) TurnOff() {
	fmt.Println("TV turned Off!")
}

func (tv *TV) SetVolume(volume int) {
	tv.volume = volume
	fmt.Println("TV volume set to: ", volume)
}
