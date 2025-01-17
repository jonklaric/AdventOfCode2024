package main

import (
	"fmt"
	"os"
	"strings"
)

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

func isXMAS(test_string string) bool {
	if test_string == "XMAS" {
		return true
	}
	return false
}

func isSAMX(test_string string) bool {
	if test_string == "SAMX" {
		return true
	}
	return false
}

func sum_chars(chars []string) string {
	var temp_string string = ""
	for i := 0; i < len(chars); i++ {
		temp_string += chars[i]
	}
	return temp_string
}

func find_xmas(textgrid []string) int {
	xmas_count := 0
	directions := [8][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for row_num, row := range textgrid {
		row_split := strings.Split(row, "")
		for col_num := range row_split {
			for _, direction := range directions {
				if string(textgrid[row_num][col_num]) != "X" {
					continue
				}
				row_start := row_num + 3*direction[0]
				col_start := col_num + 3*direction[1]
				if row_start < 0 || row_start >= len(textgrid) || col_start < 0 || col_start >= len(textgrid[row_start]) {
					continue
				}
				var temp_string string = ""
				for i := 0; i < 4; i++ {
					temp_string += string(textgrid[row_num+i*direction[0]][col_num+i*direction[1]])
				}
				if isXMAS(temp_string) || isSAMX(temp_string) {
					xmas_count++
				}
			}
		}
	}
	return xmas_count
}

func isCrossMas(testString string) bool {
	if testString == "MS" || testString == "SM" {
		return true
	}
	return false
}

func part_two(textgrid []string) int {
	xmas_count := 0
	//directions := [4][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	charGrid := make([][]string, len(textgrid))
	for i := 0; i < len(textgrid); i++ {
		charGrid[i] = make([]string, len(textgrid[0]))
		row_split := strings.Split(textgrid[i], "")
		for j := 0; j < len(textgrid[0]); j++ {
			charGrid[i][j] = row_split[j]
		}
	}
	for row_num, row := range charGrid {
		for col_num, char := range row {
			if col_num < 1 || row_num < 1 {
				continue
			}
			if col_num >= len(textgrid[0])-1 || row_num >= len(textgrid)-1 {
				continue
			}
			if char != "A" {
				continue
			}
			var diag_string1 string = charGrid[row_num-1][col_num-1] + charGrid[row_num+1][col_num+1]
			var diag_string2 string = charGrid[row_num+1][col_num-1] + charGrid[row_num-1][col_num+1]
			if isCrossMas(diag_string1) && isCrossMas(diag_string2) {
				xmas_count++
			}
		}
	}
	return xmas_count
}

func main() {
	fmt.Println("Welcome to day 4!")
	grid := read_input("./my_input.txt")
	xmas_count := find_xmas(grid)
	fmt.Println("xmas_count:", xmas_count)
	fmt.Println("Now onto day 4 part 2!")
	grid2 := read_input("./my_input.txt")
	xmas_count_part2 := part_two(grid2)
	fmt.Println("xmas_count:", xmas_count_part2)
}
