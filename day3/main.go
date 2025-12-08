package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	res := 0
	for _, seq := range input {
		batteries := strings.Split(seq, "")
		maxJoltage := 0

		for i, battery1 := range batteries {
			for _, battery2 := range batteries[i+1:] {
				joltageStr := battery1 + battery2
				joltage, err := strconv.Atoi(joltageStr)
				if err != nil {
					panic(err.Error())
				}
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}

		}
		fmt.Println(maxJoltage)
		res += maxJoltage
	}

	fmt.Println("sum max joltage: ", res)

}

func readInput() []string {
	fb, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}
	rawInput := string(fb)
	return strings.Split(rawInput, "\n")
}
