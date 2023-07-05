package main

import (
	"fmt"
)

const (
	Rows = 6
	Cols = 4
)

type Grid interface {
	Initialize(rows, cols int)
	Print()
	Contains(s string) bool
	ConvertToCoords(move string) (int8, int8)
	ChangePos(pos *Position, move string) bool
	SetPos(pos Position)
	ClearPos(pos Position)
}

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

type TwoDGrid struct {
	data [][]uint8
}

func (g *TwoDGrid) Initialize(rows, cols int) {
	g.data = make([][]uint8, cols)
	for i := 0; i < cols; i++ {
		g.data[i] = make([]uint8, rows)
	}
}

func (g *TwoDGrid) Print() {
	for _, row := range g.data {
		fmt.Println(row)
	}
}

func (g *TwoDGrid) Contains(s string) bool {
	for _, value := range validMoves {
		if value == s {
			return true
		}
	}
	return false
}

func (g *TwoDGrid) ConvertToCoords(move string) (int8, int8) {
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

func (g *TwoDGrid) ChangePos(pos *Position, move string) bool {
	nx, ny := g.ConvertToCoords(move)
	newPx := int8(pos.x) + nx
	newPy := int8(pos.y) + ny

	rows := len(g.data)
	cols := len(g.data[0])

	if newPx >= 0 && newPx < int8(rows) && newPy >= 0 && newPy < int8(cols) {
		g.ClearPos(*pos)
		pos.x = uint8(newPx)
		pos.y = uint8(newPy)
		g.SetPos(*pos)
		return true
	}

	return false
}

func (g *TwoDGrid) SetPos(pos Position) {
	g.data[pos.x][pos.y] = 1
}

func (g *TwoDGrid) ClearPos(pos Position) {
	g.data[pos.x][pos.y] = 0
}

func main() {
	var grid Grid = &TwoDGrid{}
	grid.Initialize(Rows, Cols)

	pos := Position{x: 0, y: 0}
	grid.SetPos(pos)

	for {
		grid.Print()

		var move string
		fmt.Print("Make move (w, a, s, d): ")
		fmt.Scan(&move)

		if grid.Contains(move) {
			if grid.ChangePos(&pos, move) {
				continue
			}
		}

		fmt.Println("Invalid move. Try again.")
	}
}
