package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/khizarIqbal93/aoc25/day2/utils"
)

func main() {
	part1()
	part2()
}

func part2() {
	input := readInput()
	total := 0

	for _, idRange := range input {
		rngStr := strings.Split(idRange, "-")

		start, err := strconv.Atoi(rngStr[0])
		if err != nil {
			panic(err.Error())
		}

		end, err := strconv.Atoi(rngStr[1])
		if err != nil {
			panic(err.Error())
		}

		for i := start; i <= end; i++ {
			if !utils.IsValidV2(i) {
				total += i
			}
		}

	}

	fmt.Println("part 2: ", total)
}

func part1() {
	input := readInput()
	total := 0

	for _, idRange := range input {
		rngStr := strings.Split(idRange, "-")

		start, err := strconv.Atoi(rngStr[0])
		if err != nil {
			panic(err.Error())
		}

		end, err := strconv.Atoi(rngStr[1])
		if err != nil {
			panic(err.Error())
		}

		for i := start; i <= end; i++ {
			if !utils.IsValid(i) {
				total += i
			}
		}

	}

	fmt.Println("part 1: ", total)
}

func readInput() []string {
	fb, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}
	rawInput := string(fb)
	return strings.Split(rawInput, ",")
}
