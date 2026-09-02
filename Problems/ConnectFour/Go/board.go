package main

type DiscColor int

const (
	Empty DiscColor = iota
	RED
	YELLOW
)

type Board struct {
	rows int
	cols int
	grid [][]DiscColor
}

func NewBoard(rows, cols int) *Board {
	grid := make([][]DiscColor, rows)
	for r := 0; r < rows; r++ {
		grid[r] = make([]DiscColor, cols)
	}

	return &Board{
		rows: rows,
		cols: cols,
		grid: grid,
	}
}

func (b *Board) canPlace(column int) bool {
	if column < 0 || column >= b.cols {
		return false
	}
	if b.grid[0][column] != Empty {
		return false
	}
	return true
}

func (b *Board) PlaceDisc(column int, color DiscColor) []int {
	if !b.canPlace(column) {
		return []int{-1, -1}
	}

	for row := 0; row < b.rows-1; row++ {
		if b.grid[row+1][column] != Empty {
			b.grid[row][column] = color
			return []int{row, column}
		}
	}

	return nil
}

func (b *Board) CheckWin(row, colum int, color DiscColor) bool {
	dirs := [][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
	}

	for _, dir := range dirs {
		dr := dir[0]
		dc := dir[1]

		count := 0

		count += b.coundDirection(row, colum, dr, dc, color)
		count += b.coundDirection(row, colum, -dr, -dc, color)

		if count >= 4 {
			return true
		}
	}

	return false
}

func (b *Board) coundDirection(row, col, dr, dc int, color DiscColor) int {
	count := 0

	for {
		row += dr
		col += dc

		if !b.isValid(row, col) {
			break
		}

		if b.grid[row][col] != color {
			break
		}

		count++
	}

	return count
}

func (b *Board) isValid(row, col int) bool {
	return row >= 0 && col >= 0 && row < b.rows && col < b.cols
}

func (b *Board) IsFull() bool {
	count := b.rows * b.cols

	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			if b.grid[i][j] != Empty {
				count--
			}
		}
	}

	if count == 0 {
		return true
	}
	
	return false
}
