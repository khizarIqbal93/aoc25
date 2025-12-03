package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	current := 50
	input := readInput()
	res := 0
	for _, instruction := range input {
		num, err := strconv.Atoi(instruction[1:])

		if err != nil {
			panic(err.Error())
		}

		if num > 100 {
			num = num % 100
		}

		switch instruction[:1] {
		case "L":
			current -= num

			if current < 0 {
				current += 100
			}

			if current == 100 {
				current = 0
			}

			if current == 0 {
				res++
			}

		case "R":
			current += num

			if current > 99 {
				current -= 100
			}

			if current == 100 {
				current = 0
			}

			if current == 0 {
				res++
			}

		}
	}

	fmt.Println(res)
}

func readInput() []string {
	fb, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}
	rawInput := string(fb)
	return strings.Split(rawInput, "\n")
}
