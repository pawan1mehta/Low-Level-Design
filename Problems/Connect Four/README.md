# Connector Four

Link: <https://www.hellointerview.com/learn/low-level-design/problem-breakdowns/connect-four>

## Clarifying Questions

- How do players interact with the game? do they just specify a column number and the disc dropts? [Ans: they just specify the column from 0 to 6]
- What should happen if someone tries to drop a disc in a column that's already full? Should I return an error, throw an exception, or just ignore it? [Ans: Return an error]
- Are we designing this to support one game at time or do we need to handle multiple concurrent games? [Ans: just one game]
- Do we need to track move history or support undo?

## Requirements

- Played vertically on a grid measuring 6 row x 7 cols
- Each disc is colored either red or yellow
- Each player has 21 disc
- On each turn, a player drops one disc from the top into any column
- A disc falls to the lowest available row in the chosen column
- The first player to align four of their own discs horizonally, vertically or diagonal wins the game

Out Of Scope

- UI Support
- Concurrent games
- Move history
- Undo

## Entities & Relationships

- Board
- Disc
- Player
- Game

Relationships:

Game <|--- composed of ---- Board

Game <--- contains ---- player

## Class Design

```code
Class Game

- board: Board
- player1: Player
- player2: Player
- currentPlayer: Player
- state: GameState  (INPROGRESS, WON, DRAW)
- winner: Player

+ Game(player1, player2)
+ makeMove(column, player) -> bool
+ getCurrentPlayer() -> Player
+ getGameState() -> GameState
+ getWinner() -> Player
+ getBoard
```

```code
Class Board

- rows: int = 6
- colums: int = 7
- grid: DiscColor?[rows][cols]

+ Board()
+ getRows() -> int
+ getCols() -> int
+ canPlace(column) -> bool
+ placeDisc(column, color) -> int
+ isFull() -> bool
+ checkWin(row, columnm, color) -> bool
+ getCell() -> DiscColor
```

```code
Class Player

- name: string
- color: DiscColor

+ Player(name, color)
+ getName() -> string
+ getColor() -> DiscColor
```
