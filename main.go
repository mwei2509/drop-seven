package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// init + cleanup
func init() {
	// enable terminal cbreak mode
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin
	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode")
	}
}

func cleanup() {
	// restore cooked mode
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin
	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode")
	}
}

func main() {
	// init game
	defer cleanup()

	// load grid
	makeGrid()

	// make go routine for processing input
	input := make(chan string) // channel
	// string/intpu tin channel?
	go func(ch chan<- string) {
		for {
			input, err := readInput()
			if err != nil {
				log.Printf("Error reading input")
				ch <- "ESC" // string "esc" to the channel
			}
			ch <- input
		}
	}(input)

	// init first round
	printScreen()

	// game loop
	for {
		select {
		case inp := <-input:
			playRound(inp)
		default:
		}

		if Lives == 0 {
			printLose()
			break
		}
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100) // make slice/array of bytes
	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}
	// escape
	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		// arrow keys
		// The escape sequence for the arrow keys are 3 bytes long, starting with ESC+[ and then a letter from A to D.
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}
	return "", nil
}

func clearScreen() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}

func printScreen() {
	clearScreen()

	// print player
	moveCursor(0, player.col)
	printCell(player.cell)
	// print maze
	moveCursor(1, 0)
	for _, row := range Grid {
		for _, cell := range row {
			printCell(cell)
			fmt.Printf("|")
		}
		fmt.Printf("\n")
	}
}

func printLose() {
	clearScreen()
	moveCursor(0, 0)
	fmt.Println("YOU LOSE, GOODBYE")
}
func printCell(cell *Cell) {
	if cell.Value == 0 {
		// empty
		fmt.Printf(" ")
	} else if cell.Lives == 2 {
		// has value and lives
		fmt.Printf("*")
	} else if cell.Lives == 1 {
		fmt.Printf("+")
	} else {
		fmt.Printf("%d", cell.Value)
	}
}
