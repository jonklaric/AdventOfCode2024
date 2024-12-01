package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func read_lists(input_string string) ([]int64, []int64) {
	var data []string
	if strings.Contains(input_string, "\r\n") {
		data = strings.Split(input_string, "\r\n")
	} else {
		data = strings.Split(input_string, "\n")
	}
	list1 := make([]int64, 0)
	list2 := make([]int64, 0)
	for _, row := range data {
		row = strings.Replace(row, "\r", "", -1)
		values := strings.Split(row, "   ")
		if len(values) != 2 {
			panic("not two elements?!?")
		}
		a, err := strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			panic(err)
		}
		list1 = append(list1, a)
		b, err := strconv.ParseInt(values[1], 10, 64)
		if err != nil {
			panic(err)
		}
		list2 = append(list2, b)
	}
	if len(list1) != len(list2) {
		panic("Lists are different lengths!")
	}
	return list1, list2
}

func day1_part1(list1, list2 []int64) int64 {
	slices.Sort(list1)
	slices.Sort(list2)
	var running_sum int64 = 0
	for j := 0; j < len(list1); j++ {
		switch {
		case list1[j] > list2[j]:
			running_sum += list1[j] - list2[j]
		case list1[j] < list2[j]:
			running_sum += list2[j] - list1[j]
		case list1[j] == list2[j]:
			running_sum += 0
		}
	}
	return running_sum
}

func main() {
	fmt.Println("Hello AoC 2024!")
	fmt.Println("Welcome to Day 1: Part 1...")
	// read in (test) data
	test_byte_data, err := os.ReadFile("./MyInput.txt")
	if err != nil {
		panic(err)
	}
	list1, list2 := read_lists(string(test_byte_data))
	part1_result := day1_part1(list1, list2)
	fmt.Println("Result:", part1_result)
	fmt.Println("*********************")
	fmt.Println("Now for part 2...")

}
