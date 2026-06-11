package main

import "fmt"

type User interface {
	GetName() string
	SendMessage(message string)
	ReceiveMessage(message string)
}

type ChatUser struct {
	chatMediator ChatMediator
	name         string
}

func NewChatUser(chatMediator ChatMediator, name string) *ChatUser {
	return &ChatUser{
		chatMediator: chatMediator,
		name:         name,
	}
}

func (cu *ChatUser) GetName() string {
	return cu.name
}

func (cu *ChatUser) SendMessage(message string) {
	fmt.Printf("%s sends: %s \n", cu.name, message)
	cu.chatMediator.SendMessage("message", cu)
}

func (cu *ChatUser) ReceiveMessage(message string) {
	fmt.Printf("%s received: %s \n", cu.name, message)
}
