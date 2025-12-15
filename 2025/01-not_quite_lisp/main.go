package main

import (
	"fmt"
	"os"
)

func main() {
	instructions, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error while reading input file:", err)
		return
	}

	partOneResult := partOne(string(instructions))

	fmt.Println("The part one response is: ", partOneResult)

	partTwoResult := partTwo(string(instructions))
	if partTwoResult < 0 {
		fmt.Println("Error while processing input. Never reached the basement")
		return
	}

	fmt.Println("The part two response is: ", partTwoResult)
}

func partOne(instructions string) int {
	floor := 0
	for _, r := range instructions {
		if r == '(' {
			floor++
		}
		if r == ')' {
			floor--
		}
	}
	return floor
}

func partTwo(instructions string) int {
	floor := 0
	for n, r := range instructions {
		if r == '(' {
			floor++
		}
		if r == ')' {
			floor--
		}
		if floor < 0 {
			return n + 1
		}
	}

	return -1
}
