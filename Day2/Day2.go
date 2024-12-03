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

func main() {
	fmt.Println("hello AoCers!")
	test_input := read_input("./test_input.txt")
	for _, record := range test_input {
		fmt.Println(record)
	}
}
