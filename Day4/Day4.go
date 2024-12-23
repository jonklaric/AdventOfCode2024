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

func isXMAS(chars []string) bool {
	test_string := sum_chars(chars)
	if test_string == "XMAS" {
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

func find_horizontal(textgrid [][]string) int {
	var xmas_found int64 = 0
	//var temp_string string = ""
	for row_num := 0; row_num < len(textgrid); row_num++ {
		for col_num := 0; col_num < len(textgrid[row_num]); col_num++ {
			//temp_string += textgrid[row_num][col_num]
			if col_num < len(textgrid[row_num])-4 {
				if sum_chars(textgrid[row_num][col_num:col_num+4]) == "SAMX" || sum_chars(textgrid[row_num][col_num:col_num+4]) == "XMAS" {
					
				}
			}
		}
	}
	return int(xmas_found)
}

func main() {
	fmt.Println("Welcome to day 4!")
	grid := read_input("./test_input.txt")
	for _, row := range grid {
		fmt.Println(row)
	}
}
