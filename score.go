package main

var Lives = 1

type CellLife int

const (
	x0 CellLife = 0

	x1 CellLife = 1
	x2 CellLife = 2
	x3 CellLife = 3
	x4 CellLife = 4
	x5 CellLife = 5
	x6 CellLife = 6
	x7 CellLife = 7

	xF CellLife = 8
	xa CellLife = 9
)

var Combos = [][]CellLife{
	// 1 in index 0
	[]CellLife{x1, x0, xa, xa, xa, xa, xa},
	// 1 in index 1
	[]CellLife{x0, x1, x0, xa, xa, xa, xa},
	// 1 in index 2
	[]CellLife{xa, x0, x1, x0, xa, xa, xa},
	// 1 in index 3
	[]CellLife{xa, xa, x0, x1, x0, xa, xa},
	// 1 in index 4
	[]CellLife{xa, xa, xa, x0, x1, x0, xa},
	// 1 in index 5
	[]CellLife{xa, xa, xa, xa, x0, x1, x0},
	// 1 in index 6
	[]CellLife{xa, xa, xa, xa, xa, x0, x1},
	// 2 in index 0
	[]CellLife{x2, xF, x0, xa, xa, xa, xa},
	// 2 in index 1
	[]CellLife{xF, x2, x0, xa, xa, xa, xa},
	[]CellLife{x0, x2, xF, x0, xa, xa, xa},
	// 2 in index 2
	[]CellLife{x0, xF, x2, x0, xa, xa, xa},
	[]CellLife{xa, x0, x2, xF, x0, xa, xa},
	// 2 in index 3
	[]CellLife{xa, x0, xF, x2, x0, xa, xa},
	[]CellLife{xa, xa, x0, x2, xF, x0, xa},
	// 2 in index 4
	[]CellLife{xa, xa, x0, xF, x2, x0, xa},
	[]CellLife{xa, xa, xa, x0, x2, xF, x0},
	// 2 in index 5
	[]CellLife{xa, xa, xa, x0, xF, x2, x0},
	[]CellLife{xa, xa, xa, xa, x0, x2, xF},
	// 2 in index 6
	[]CellLife{xa, xa, xa, xa, x0, xF, x2},
	// 3 in index 0
	[]CellLife{x3, xF, xF, x0, xa, xa, xa},
	// 3 in index 1
	[]CellLife{xF, x3, xF, x0, xa, xa, xa},
	[]CellLife{x0, x3, xF, xF, x0, xa, xa},
	// 3 in index 2
	[]CellLife{xF, xF, x3, x0, xa, xa, xa},
	[]CellLife{x0, xF, x3, xF, x0, xa, xa},
	[]CellLife{xa, x0, x3, xF, xF, x0, xa},
	// 3 in index 3
	[]CellLife{x0, xF, xF, x3, x0, xa, xa},
	[]CellLife{xa, x0, xF, x3, xF, x0, xa},
	[]CellLife{xa, xa, x0, x3, xF, xF, x0},
	// 3 in index 4
	[]CellLife{xa, x0, xF, xF, x3, xa, xa},
	[]CellLife{xa, xa, x0, xF, x3, xF, x0},
	[]CellLife{xa, xa, xa, x0, x3, xF, xF},
	// 3 in index 5
	[]CellLife{xa, xa, x0, xF, xF, x3, x0},
	[]CellLife{xa, xa, xa, x0, xF, x3, xF},
	// 3 in index 6
	[]CellLife{xa, xa, xa, x0, xF, xF, x3},
	// 4 in index 0
	[]CellLife{x4, xF, xF, xF, x0, xa, xa},
	// 4 in index 1
	[]CellLife{xF, x4, xF, xF, x0, xa, xa},
	[]CellLife{x0, x4, xF, xF, xF, x0, xa},
	// 4 in index 2
	[]CellLife{xF, xF, x4, xF, x0, xa, xa},
	[]CellLife{x0, xF, x4, xF, xF, x0, xa},
	[]CellLife{xa, x0, x4, xF, xF, xF, x0},
	// 4 in index 3
	[]CellLife{xF, xF, xF, x4, x0, xa, xa},
	[]CellLife{x0, xF, xF, x4, xF, x0, xa},
	[]CellLife{xa, x0, xF, x4, xF, xF, x0},
	[]CellLife{xa, xa, x0, x4, xF, xF, xF},
	// 4 in index 4
	[]CellLife{x0, xF, xF, xF, x4, x0, xa},
	[]CellLife{xa, x0, xF, xF, x4, xF, x0},
	[]CellLife{xa, xa, x0, xF, x4, xF, xF},
	// 4 in index 5
	[]CellLife{xa, x0, xF, xF, xF, x4, x0},
	[]CellLife{xa, xa, x0, xF, xF, x4, xF},
	// 4 in index 6
	[]CellLife{xa, xa, x0, xF, xF, xF, x4},
	// 5 in index 0
	[]CellLife{x5, xF, xF, xF, xF, x0, xa},
	// 5 in index 1
	[]CellLife{xF, x5, xF, xF, xF, x0, xa},
	[]CellLife{x0, x5, xF, xF, xF, xF, x0},
	// 5 in index 2
	[]CellLife{xF, xF, x5, xF, xF, x0, xa},
	[]CellLife{x0, xF, x5, xF, xF, xF, x0},
	[]CellLife{xa, x0, x5, xF, xF, xF, xF},
	// 5 in index 3
	[]CellLife{xF, xF, xF, x5, xF, x0, xa},
	[]CellLife{x0, xF, xF, x5, xF, xF, x0},
	[]CellLife{xa, x0, xF, x5, xF, xF, xF},
	// 5 in index 4
	[]CellLife{xF, xF, xF, xF, x5, x0, xa},
	[]CellLife{x0, xF, xF, xF, x5, xF, x0},
	[]CellLife{xa, x0, xF, xF, x5, xF, xF},
	// 5 in index 5
	[]CellLife{x0, xF, xF, xF, xF, x5, x0},
	[]CellLife{xa, x0, xF, xF, xF, x5, xF},
	// 5 in index 6
	[]CellLife{xa, x0, xF, xF, xF, xF, x5},
	// 6 in index 0
	[]CellLife{x6, xF, xF, xF, xF, xF, x0},
	// 6 in index 1
	[]CellLife{xF, x6, xF, xF, xF, xF, x0},
	[]CellLife{x0, x6, xF, xF, xF, xF, xF},
	// 6 in index 2
	[]CellLife{xF, xF, x6, xF, xF, xF, x0},
	[]CellLife{x0, xF, x6, xF, xF, xF, xF},
	// 6 in index 3
	[]CellLife{xF, xF, xF, x6, xF, xF, x0},
	[]CellLife{x0, xF, xF, x6, xF, xF, xF},
	// 6 in index 4
	[]CellLife{xF, xF, xF, xF, x6, xF, x0},
	[]CellLife{x0, xF, xF, xF, x6, xF, xF},
	// 6 in index 5
	[]CellLife{xF, xF, xF, xF, xF, x6, x0},
	[]CellLife{x0, xF, xF, xF, xF, x6, xF},
	// 6 in index 6
	[]CellLife{x0, xF, xF, xF, xF, xF, x6},
	// 7 in index 0
	[]CellLife{x7, xF, xF, xF, xF, xF, xF},
	// 7 in index 1
	[]CellLife{xF, x7, xF, xF, xF, xF, xF},
	// 7 in index 2
	[]CellLife{xF, xF, x7, xF, xF, xF, xF},
	// 7 in index 3
	[]CellLife{xF, xF, xF, x7, xF, xF, xF},
	// 7 in index 4
	[]CellLife{xF, xF, xF, xF, x7, xF, xF},
	// 7 in index 5
	[]CellLife{xF, xF, xF, xF, xF, x7, xF},
	// 7 in index 6
	[]CellLife{xF, xF, xF, xF, xF, xF, x7},
}

func getCombos(droppedCell *Cell, index int) (ret [][]CellLife) {
	for _, combo := range Combos {
		if combo[index] == CellLife(droppedCell.Value) {
			ret = append(ret, combo)
		}
	}
	return
}

func checkValid(checkCell *Cell, index int, combos [][]CellLife) (ret [][]CellLife) {
	cellVal := CellLife(checkCell.Value)
	for _, combo := range combos {
		switch combo[index] {
		case xa:
			ret = append(ret, combo)
		case xF:
			if cellVal != x0 {
				ret = append(ret, combo)
			}
		case x0:
			if cellVal == x0 {
				ret = append(ret, combo)
			}

		}
	}
	return
}

// func filterCombos(combos [][]CellLife, index int, c CellLife) (ret [][]CellLife) {
// 	for _, combo := range combos {
// 		if combo[index] == c {
// 			ret = append(ret, combo)
// 		}
// 	}
// 	return
// }
