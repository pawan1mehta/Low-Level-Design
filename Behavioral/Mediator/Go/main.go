package main

func main() {

	chatMediator := NewChatRoom()

	user1 := NewChatUser(chatMediator, "user1")
	user2 := NewChatUser(chatMediator, "user2")
	user3 := NewChatUser(chatMediator, "user3")

	chatMediator.Add(user1)
	chatMediator.Add(user2)
	chatMediator.Add(user3)

	user1.SendMessage("hii")
}
