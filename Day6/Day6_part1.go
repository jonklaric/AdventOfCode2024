package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row int
	col int
}

var directions [4]Position = [...]Position{Position{row: -1, col: 0},
	Position{row: 0, col: 1},
	Position{row: 1, col: 0},
	Position{row: 0, col: -1}}

type guard struct {
	currectPos Position          // current position of guard
	direction  Position          // direction guard is facing: (1,0), (-1,0), (0, 1) or (0, -1)
	history    map[Position]bool // history of visited locations
}

type game struct {
	board board
	guard guard
}

func (g guard) step(b board) bool {
	if (g.currectPos.row == 0 && g.direction.row == -1) || (g.currectPos.row == len(b.grid) && g.direction.row == 1) ||
		(g.currectPos.col == 0 && g.direction.col == -1) || (g.currectPos.col == len(b.grid) && g.direction.row == 1) {
		return true
	}
	temp_position := g.currectPos
	temp_position.row += g.direction.row

}

type board struct {
	grid           []string
	blockLocations map[Position]string
}

func (b board) findKeyLocations() guard {
	// appends block locations and returns guard starting position.
	var guard_start guard
	b.blockLocations = make(map[Position]string)
	for i := 0; i < len(b.grid); i++ {
		for j := 0; j < len(b.grid[i]); j++ {
			if string(b.grid[i][j]) == "#" {
				b.blockLocations[Position{row: i, col: j}] = "#"
			} else if string(b.grid[i][j]) == "^" {
				guard_start.currectPos.row = i
				guard_start.currectPos.col = j
				guard_start.direction.row = -1
				guard_start.direction.col = 0
			}
		}
	}
	return guard_start
}

func read_input(input_file string) []string {
	test_byte_data, err := os.ReadFile(input_file)
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(test_byte_data), "\r") {
		return strings.Split(string(test_byte_data), "\r\n")
	} else {
		return strings.Split(string(test_byte_data), "\n")
	}
}

func main() {
	fmt.Println("Hello AoCers!")
	fmt.Println("Welcome to day 6 part 1...")

	test_input := read_input("./test_input.txt")
	test_board := board{grid: test_input}
	test_guard := test_board.findKeyLocations()
}
