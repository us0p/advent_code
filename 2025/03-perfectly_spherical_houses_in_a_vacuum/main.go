package main

import (
	"fmt"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

const (
	NORTH = byte('^')
	EAST  = byte('>')
	SOUTH = byte('v')
	WEST  = byte('<')
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	partOne := PartOne(file)
	fmt.Println("Part one response is:", partOne)

	partTwo := PartTwo(file)
	fmt.Println("Part two response is:", partTwo)
}

func PartOne(directions []byte) int {
	visitedPoints := make(map[Point]bool)
	x, y := 0, 0
	for _, direction := range directions {
		point := Point{x, y}
		visitedPoints[point] = true
		switch direction {
		case NORTH:
			y++
		case EAST:
			x++
		case SOUTH:
			y--
		case WEST:
			x--
		default:
			continue
		}
	}

	return len(visitedPoints)
}

func PartTwo(directions []byte) int {
	visitedPoints := make(map[Point]bool)
	santa := Point{0, 0}
	robSanta := Point{0, 0}
	for n, direction := range directions {
		var turn *Point
		if n%2 == 0 {
			turn = &santa
		} else {
			turn = &robSanta
		}
		visitedPoints[*turn] = true
		switch direction {
		case NORTH:
			turn.y++
		case EAST:
			turn.x++
		case SOUTH:
			turn.y--
		case WEST:
			turn.x--
		default:
			continue
		}
	}

	return len(visitedPoints)
}
