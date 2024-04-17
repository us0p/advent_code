package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type coordinate struct {
	x int
	y int
}

const (
	UP    = "^"
	RIGHT = ">"
	DOWN  = "v"
	LEFT  = "<"
)

func main() {
	fileName := filepath.Join("..", "inputs.txt")
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	houseGrid := createHouseGrid()

	santaPosition := coordinate{4, 4}
	robotSantaPosition := coordinate{4, 4}
	direction := make([]byte, 1)
	visitedHouses := 1
	isSantaTurn := true
	for _, err := file.Read(direction); err != io.EOF; _, err = file.Read(direction) {
		if isSantaTurn {
			houseGrid, robotSantaPosition = expandGrid(
				houseGrid,
				string(direction),
				santaPosition,
				robotSantaPosition,
			)
			visitedHouses += walkDirections(string(direction), houseGrid, &santaPosition)

			isSantaTurn = false
			continue
		}

		houseGrid, santaPosition = expandGrid(
			houseGrid,
			string(direction),
			robotSantaPosition,
			santaPosition,
		)
		visitedHouses += walkDirections(string(direction), houseGrid, &robotSantaPosition)
		isSantaTurn = true
	}

	fmt.Println(visitedHouses)
}

func createHouseGrid() [][]bool {
	houseGrid := make([][]bool, 10)

	for i := range houseGrid {
		houseGrid[i] = make([]bool, 10)
		// setting start position to the middle of the matrix
		if i == 4 {
			houseGrid[4][4] = true
		}
	}

	return houseGrid
}

func expandGrid(
	houseGrid [][]bool,
	direction string,
	primaryPosition,
	secondaryPosition coordinate,
) ([][]bool, coordinate) {
	switch direction {
	case UP:
		if primaryPosition.y-1 < 0 {
			temp := expandY(houseGrid)
			for i := range temp {
				for j := range temp[i] {
					if i > 0 {
						temp[i][j] = houseGrid[i-1][j]
					}
				}
			}
			secondaryPosition.y += 1
			return temp, secondaryPosition
		}
		return houseGrid, secondaryPosition
	case RIGHT:
		if primaryPosition.x+1 >= len(houseGrid[0]) {
			// add a new collumn to the end of each line
			temp := expandX(houseGrid)
			copyGridFromStart(temp, houseGrid)
			return temp, secondaryPosition
		}
		return houseGrid, secondaryPosition
	case DOWN:
		if primaryPosition.y+1 >= len(houseGrid) {
			// append a new line to the matrix;
			temp := expandY(houseGrid)
			copyGridFromStart(temp, houseGrid)
			return temp, secondaryPosition
		}
		return houseGrid, secondaryPosition
	case LEFT:
		if primaryPosition.x-1 < 0 {
			// add a new collumn to beginning of each line
			temp := expandX(houseGrid)
			for i := range houseGrid {
				for j := range houseGrid[i] {
					temp[i][j+1] = houseGrid[i][j]
				}
			}
			secondaryPosition.x += 1
			return temp, secondaryPosition
		}
		return houseGrid, secondaryPosition
	default:
		return houseGrid, secondaryPosition
	}
}

func expandY(houseGrid [][]bool) [][]bool {
	temp := make([][]bool, len(houseGrid)+1)
	for i := range temp {
		temp[i] = make([]bool, len(houseGrid[0]))
	}
	return temp
}

func expandX(houseGrid [][]bool) [][]bool {
	temp := make([][]bool, len(houseGrid))
	for i := range temp {
		temp[i] = make([]bool, len(houseGrid[0])+1)
	}
	return temp
}

func copyGridFromStart(newGrid, oldGrid [][]bool) {
	for y := range oldGrid {
		for x := range oldGrid[y] {
			newGrid[y][x] = oldGrid[y][x]
		}
	}
}

func walkDirections(
	direction string,
	houseGrid [][]bool,
	primaryPosition *coordinate,
) int {
	switch direction {
	case UP:
		if primaryPosition.y-1 >= 0 {
			primaryPosition.y -= 1
		}
		return checkVisitedHouse(houseGrid, *primaryPosition)
	case RIGHT:
		if primaryPosition.x+1 < len(houseGrid[0]) {
			primaryPosition.x += 1
		}
		return checkVisitedHouse(houseGrid, *primaryPosition)
	case DOWN:
		if primaryPosition.y+1 < len(houseGrid) {
			primaryPosition.y += 1
		}
		return checkVisitedHouse(houseGrid, *primaryPosition)
	case LEFT:
		if primaryPosition.x-1 >= 0 {
			primaryPosition.x -= 1
		}
		return checkVisitedHouse(houseGrid, *primaryPosition)
	default:
		return 0
	}
}

func checkVisitedHouse(houseGrid [][]bool, position coordinate) int {
	if !houseGrid[position.y][position.x] {
		houseGrid[position.y][position.x] = true
		return 1
	}

	return 0
}

