package main

type GameState int

const (
	InProgress GameState = iota
	WON
	DRAW
)

type Game struct {
	board         *Board
	player1       *Player
	player2       *Player
	currentPlayer *Player
	state         GameState
	winer         *Player
}

func NewGame(player1, player2 *Player) *Game {
	return &Game{
		board:   NewBoard(6, 7),
		player1: player1,
		player2: player2,
		state:   InProgress,
	}
}

func (g *Game) MakeMove(column int, player *Player) bool {
	if g.state != InProgress {
		return false
	}

	if player != g.currentPlayer {
		return false
	}

	position := g.board.PlaceDisc(column, player.GetColor())
	if position[0] == -1 {
		return false
	}

	if g.board.CheckWin(position[0], position[1], player.GetColor()) {
		g.state = WON
		g.winer = player
	} else if g.board.IsFull() {
		g.state = DRAW
	} else {
		if player == g.player1 {
			g.currentPlayer = g.player2
		} else {
			g.currentPlayer = g.player1
		}
	}

	return true
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.currentPlayer
}

func (g *Game) GameState() GameState {
	return g.state
}

func (g *Game) GetWinner() *Player {
	return g.winer
}
