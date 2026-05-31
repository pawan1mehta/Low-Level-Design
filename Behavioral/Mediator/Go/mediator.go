package main

type ChatMediator interface {
	SendMessage(message string, sender User)
}

type ChatRoom struct {
	users []User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: []User{},
	}
}

func (chatRoom *ChatRoom) Add(user User) {
	chatRoom.users = append(chatRoom.users, user)
}

func (chatRoom *ChatRoom) SendMessage(message string, sender User) {
	for _, user := range chatRoom.users {
		if user != sender {
			user.ReceiveMessage(message)
		}
	}
}
