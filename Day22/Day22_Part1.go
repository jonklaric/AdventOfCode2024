package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mix(secret, value int64) int64 {
	return secret ^ value
}

func prune(secret int64) int64 {
	return secret % int64(16777216)
}

func getNextSecret(secret int64) int64 {
	var result int64 = 64
	result *= secret
	secret = mix(secret, result)
	result = prune(secret)
	result /= int64(32)
	secret = mix(secret, result)
	result = prune(secret)
	result *= int64(2048)
	secret = mix(secret, result)
	result = prune(secret)
	return result
}

func get2000thSecret(secret int64) int64 {
	for i := 0; i < 2000; i++ {
		secret = getNextSecret(secret)
	}
	return secret
}

func read_input(filepath string) []int64 {
	byte_data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	input_string := string(byte_data)
	var data []string
	if strings.Contains(input_string, "\r\n") {
		data = strings.Split(input_string, "\r\n")
	} else {
		data = strings.Split(input_string, "\n")
	}
	intial_list := make([]int64, 0)
	for _, value := range data {
		a, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
		intial_list = append(intial_list, a)
	}
	return intial_list
}

func main() {
	fmt.Println("Hello AoC2024ers! Welcome to day 22 (part 1)!")
	var secret int64 = 42
	var result int64 = 15
	secret = secret ^ result
	fmt.Println("Secret:", secret)
	secret = 100000000
	fmt.Println("Prune(100000000)=", prune(secret))
	secret = 123
	for i := 0; i < 10; i++ {
		fmt.Println("Secret =", secret)
		secret = getNextSecret(secret)
	}
	fmt.Println("10th Secret =", secret)
	test_initial_list := read_input("./my_input.txt")
	sum := int64(0)
	for _, secret := range test_initial_list {
		final_secret := get2000thSecret(secret)
		fmt.Println("secret", secret, "Final secret (secret):", final_secret)
		sum += final_secret
	}
	fmt.Println("Sum:", sum)
}
