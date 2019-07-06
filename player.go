package main

import (
	"math/rand"
	"time"
)

var player = initPlayer()
var round = 0

func playRound(inp string) {
	switch inp {
	case "ESC":
		// quit
		Lives = 0
	case "RIGHT":
		fallthrough
	case "LEFT":
		player.movePlayer(inp)
		printScreen()
	case "DOWN":
		round++
		dropPlayerCell()
		resetPlayer()

		printScreen()
		time.Sleep(400 * time.Millisecond)

		updateGrid(true)
		printScreen()

		// wait a bit
		time.Sleep(400 * time.Millisecond)
		advanceGrid()

		if Lives == 0 {
			return
		}
		printScreen()
		updateGrid(true)
		printScreen()
	}
}

func dropPlayerCell() {
	dropCol := player.col / 2
	dropRow := 0
	// find out how to do reverse range
	for i, row := range Grid {
		if !row[dropCol].IsEmpty() {
			break
		}
		dropRow = i
	}
	// update cell value
	Grid[dropRow][dropCol].DropInto(player.cell.Value)
}

type Player struct {
	cell *Cell
	col  int
}

func resetPlayer() {
	player.cell = newPlayerCell()
	player.col = 0
}

func initPlayer() *Player {
	return &Player{
		cell: newPlayerCell(),
		col:  0,
	}
}

// move left or right
func (p *Player) movePlayer(dir string) {
	p.col = makeMove(p.col, dir)
}

func newPlayerCell() *Cell {
	return &Cell{
		Value:    rand.Intn(7) + 1,
		Lives:    0,
		Revealed: true,
	}
}
