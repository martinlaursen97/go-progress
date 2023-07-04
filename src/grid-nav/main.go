package main

/////////////////////////////////////
// Navigate a grid with w, a, s, d //
/////////////////////////////////////

import (
	"fmt"
)

const (
	Rows = 6
	Cols = 4
)

type Grid [][]uint8

type Position struct {
	x, y uint8
}

const (
	MoveUp    = "w"
	MoveLeft  = "a"
	MoveDown  = "s"
	MoveRight = "d"
)

var validMoves = []string{MoveUp, MoveLeft, MoveDown, MoveRight}

func main() {
	grid := initializeGrid(Rows, Cols)
	mainLoop(&grid)
}

func initializeGrid(rows, cols int) Grid {
	grid := make(Grid, cols)

	for i := 0; i < cols; i++ {
		grid[i] = make([]uint8, rows)
	}
	return grid
}

func printGrid(grid Grid) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func contains(s string, arr []string) bool {
	for _, value := range arr {
		if value == s {
			return true
		}
	}
	return false
}

func convertToCoords(move string) (int8, int8) {
	switch move {
	case MoveUp:
		return -1, 0
	case MoveLeft:
		return 0, -1
	case MoveDown:
		return 1, 0
	case MoveRight:
		return 0, 1
	default:
		return 0, 0
	}
}

func changePos(gridP *Grid, pos *Position, move string) bool {
	nx, ny := convertToCoords(move)
	newPx := int8(pos.x) + nx
	newPy := int8(pos.y) + ny

	rows := len(*gridP)
	cols := len((*gridP)[0])

	if newPx >= 0 && newPx < int8(rows) && newPy >= 0 && newPy < int8(cols) {
		clearPos(gridP, *pos)
		pos.x = uint8(newPx)
		pos.y = uint8(newPy)
		setPos(gridP, *pos)
		return true
	}

	return false
}

func setPos(gridP *Grid, pos Position) {
	(*gridP)[pos.x][pos.y] = 1
}

func clearPos(gridP *Grid, pos Position) {
	(*gridP)[pos.x][pos.y] = 0
}

func mainLoop(gridP *Grid) {
	pos := Position{x: 0, y: 0}
	setPos(gridP, pos)

	for {
		printGrid(*gridP)

		var move string
		fmt.Print("Make move (w, a, s, d): ")
		fmt.Scan(&move)

		if contains(move, validMoves) {
			if changePos(gridP, &pos, move) {
				continue
			}
		}

		fmt.Println("Invalid move. Try again.")
	}
}
