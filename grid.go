package main

import (
	"math/rand"
	"time"
)

// CELLS
type Cell struct {
	Value    int // Empty = 0, 1-7
	Lives    int // Starts with 2, 0 = revealed
	Revealed bool
}

func newEmptyCell() *Cell {
	return &Cell{
		Value: 0,
	}
}

func newHiddenCell() *Cell {
	return &Cell{
		Value:    rand.Intn(7) + 1,
		Lives:    2,
		Revealed: false,
	}
}

func (c *Cell) IsEmpty() bool {
	return c.Value == 0
}

// NewHidden populates cell with value and set at 2 lives (hidden)
func (c *Cell) NewHidden() {
	c.Value = rand.Intn(7) + 1
	c.Lives = 2
	c.Revealed = false
}

// DropInto populates cell with known value and sets at 0 lives
func (c *Cell) DropInto(value int) {
	c.Value = value
	c.Lives = 0
	c.Revealed = true
}

// Explode resets the cell to empty
func (c *Cell) Explode() {
	c.Reset()
}

// Reset sets cell to empty
func (c *Cell) Reset() {
	c.Value = 0
	c.Lives = 2
}

// Shatter decreases the cell life
func (c *Cell) Shatter() {
	// if not empty and has lives
	if c.Value > 0 && c.Lives > 0 {
		c.Lives = c.Lives - 1
	}
}

// Grid
var Grid [][]*Cell

func makeGrid() {
	seven := []int{0, 1, 2, 3, 4, 5, 6}
	for row, _ := range seven {
		newRow := []*Cell{}
		for range seven {
			if row == 6 {
				newRow = append(newRow, newHiddenCell())
			} else {
				newRow = append(newRow, newEmptyCell())
			}
		}
		Grid = append(Grid, newRow)
	}
}

var ticker = 0
var maxTicker = 7

func advanceGrid() {
	ticker++
	if ticker == maxTicker {
		// make new grid
		newGrid := [][]*Cell{}
		for i, row := range Grid {
			if i == 0 {
				for _, cell := range row {
					if !cell.IsEmpty() {
						// if there is stuff in here, lose the game
						Lives = 0
						return
					}
				}
				continue
			}
			newGrid = append(newGrid, row)
		}
		// add new hidden grid
		seven := []int{0, 1, 2, 3, 4, 5, 6}
		newRow := []*Cell{}
		for range seven {
			newRow = append(newRow, newHiddenCell())
		}
		newGrid = append(newGrid, newRow)

		// update grid
		Grid = newGrid

		// reset ticker
		ticker = 0
	}
}

func makeMove(oldCol int, dir string) (newCol int) {
	newCol = oldCol
	switch dir {
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(Grid[0]) {
			newCol = oldCol
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = oldCol
		}
	}
	return
}

func updateGrid(newExplosion bool) {
	if !newExplosion {
		return
	}

	cellsToExplode := []map[string]int{}
	// loop through all and check visible numbers, if they should explode, explode them (shattering side lives)
	for rowIndex, row := range Grid {
		for colIndex, cell := range row {
			explodes := checkExplode(cell, colIndex, rowIndex)
			if explodes {
				cellsToExplode = append(cellsToExplode, map[string]int{
					"row": rowIndex,
					"col": colIndex,
				})
			}
		}
	}

	hasNewExplosion := false
	if len(cellsToExplode) > 0 {
		hasNewExplosion = true
		for _, toExplode := range cellsToExplode {
			explodeAdjacentCells(toExplode["col"], toExplode["row"]) // breaks adjacent cells
			Grid[toExplode["row"]][toExplode["col"]].Explode()
		}
	}

	// loop through all colls, rows from the bottom! mark empty spots and drop the cells
	// make these less explicit
	cols := []int{0, 1, 2, 3, 4, 5, 6}
	rows := []int{6, 5, 4, 3, 2, 1, 0}
	for _, colIndex := range cols {
		emptyRowIndex := -1
		for _, rowIndex := range rows {
			if Grid[rowIndex][colIndex].IsEmpty() {
				emptyRowIndex = rowIndex
			} else {
				if emptyRowIndex > -1 {
					// drop into empty row
					Grid[emptyRowIndex][colIndex] = Grid[rowIndex][colIndex]
					// empty current row
					Grid[rowIndex][colIndex] = newEmptyCell()
					// update empty row index
					emptyRowIndex = rowIndex
				}
			}
		}
	}

	printScreen()
	if hasNewExplosion {
		time.Sleep(400 * time.Millisecond)
	}
	// repeat above until no explosions
	updateGrid(hasNewExplosion)
}

func checkExplode(cell *Cell, col int, row int) bool {
	if cell.Value > 0 && cell.Lives == 0 {
		// horizontal
		// filter current combos
		combos := getCombos(cell, col)
		for idx := range []int{0, 1, 2, 3, 4, 5, 6} {
			if idx != col {
				c := Grid[row][idx]
				combos = checkValid(c, idx, combos)
				// early break
				if len(combos) == 0 {
					break
				}
			}
		}
		// if valid combo
		if len(combos) > 0 {
			return true
		}

		// vertical
		combos = getCombos(cell, row)
		for idx := range []int{0, 1, 2, 3, 4, 5, 6} {
			if idx != row {
				c := Grid[idx][col]
				combos = checkValid(c, idx, combos)
				// if no combos apply, skip
				if len(combos) == 0 {
					return false
				}
			}
		}
		return true
		// switch cell.Value {
		// case 1:
		// 	fallthrough
		// case 6:
		// 	fallthrough
		// case 7:
		// 	seven := []int{0, 1, 2, 3, 4, 5, 6}
		// 	horizontalFilled := 0
		// 	verticalFilled := 0

		// 	for idx := range seven {
		// 		if !Grid[row][idx].IsEmpty() {
		// 			horizontalFilled++
		// 		}
		// 		if !Grid[idx][col].IsEmpty() {
		// 			verticalFilled++
		// 		}
		// 	}

		// 	return horizontalFilled == cell.Value || verticalFilled == cell.Value
		// case 2:
		// 	seven := []int{0, 1, 2, 3, 4, 5, 6}
		// 	horizontalFilled := 0
		// 	// horizontal
		// 	if col == 0 {
		// 		// first
		// 		horizontalFilled := 1 // index 0
		// 		for i = 1; i < 5; i++ {
		// 			if Grid[row][col+1].IsEmpty() {
		// 				return horizontalFilled == cell.Value;
		// 			} else {
		// 				horizontalFilled++
		// 			}
		// 		}
		// 	} else if col == 6 {
		// 		// last
		// 	} else {
		// 		// middle
		// 	}

		// 	if col == 0 && !Grid[row][col+1].IsEmpty() && Grid[row][col+2].IsEmpty() {
		// 		return true
		// 	} else if (col == 6 && !Grid[row][col-1].IsEmpty() && Grid[row][col-2].IsEmpty()) {
		// 		return true
		// 	} else {
		// 		horizontalFilled := 1
		// 		seven := []int{0, 1, 2, 3, 4, 5, 6}
		// 		for idx := range seven {
		// 			if !Grid[row][col-idx].IsEmpty()
		// 		}
		// 	}
		// case 3:
		// case 4:
		// case 5:

		// }

		// // check vertical
		// for _, rows := range Grid {

		// }
		// // vertical explosions
		// if row+cell.Value == 7 {
		// 	return true
		// }

		// // left side check
		// for i := 0; i < cell.Value; i++ {

		// }
		// if col+1 == val {
		// 	// check that they are empty
		// }
	}
	return false
}

func explodeAdjacentCells(col, row int) {
	// left
	if col-1 >= 0 {
		Grid[row][col-1].Shatter()
	}
	// top
	if row-1 >= 0 {
		Grid[row-1][col].Shatter()
	}
	// right
	if col+1 <= 6 {
		Grid[row][col+1].Shatter()
	}
	// bottom
	if row+1 <= 6 {
		Grid[row+1][col].Shatter()
	}
}
