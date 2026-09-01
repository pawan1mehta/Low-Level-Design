package main

func main() {
	player1 := NewPlayer("pawa1", RED)
	player2 := NewPlayer("pawa2", RED)

	game := NewGame(player1, player2)

	game.MakeMove(0, player1)

	game.GetWinner()
}
