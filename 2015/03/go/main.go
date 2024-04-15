package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type grid struct {
	x int
	y int
}

type direction string

const (
	up    direction = "^"
	right direction = ">"
	down  direction = "v"
	left  direction = "<"
)

func visitHouse(houseGrid [][]bool, position grid) int {
	presentsDelivered := 0

	if !houseGrid[position.y][position.x] {
		presentsDelivered = 1
	}

	houseGrid[position.y][position.x] = true

	return presentsDelivered
}

func expandY(houseGrid [][]bool) [][]bool {
	tmp := make([][]bool, len(houseGrid)+1)
	for idx := range tmp {
		tmp[idx] = make([]bool, len(houseGrid[0]))
	}
	return tmp
}

func expandX(houseGrid [][]bool) [][]bool {
	tmp := make([][]bool, len(houseGrid))
	for idx := range tmp {
		tmp[idx] = make([]bool, len(houseGrid[0])+1)
	}
	return tmp
}

func copyFromStart(destination, source [][]bool) {
	for i := range source {
		for j := range source[i] {
			destination[i][j] = source[i][j]
		}
	}
}

func calculateVisitedHouses(
	character []byte,
	visitedHousesCount *int,
	houseGrid *[][]bool,
	position *grid,
    secondaryPosition *grid,
) {
	switch char := direction(character[:1]); char {
	case up:
		if position.y-1 < 0 {
			tmp := expandY(*houseGrid)
			for i := range tmp {
				for j := range tmp[i] {
					if i > 0 {
						tmp[i][j] = (*houseGrid)[i-1][j]
					}
				}
			}
			*houseGrid = tmp
			// reseting y to 1
			position.y += 1
            secondaryPosition.y += 1
		}
		position.y -= 1
		*visitedHousesCount += visitHouse(*houseGrid, *position)
	case right:
		if position.x+1 >= len((*houseGrid)[0]) {
			// add a new collumn to the end of each line
			tmp := expandX(*houseGrid)
			copyFromStart(tmp, *houseGrid)
			*houseGrid = tmp
		}
		position.x += 1
		*visitedHousesCount += visitHouse(*houseGrid, *position)
	case down:
		if position.y+1 >= len(*houseGrid) {
			// append a new line to the matrix;
			tmp := expandY(*houseGrid)
			copyFromStart(tmp, *houseGrid)
			*houseGrid = tmp
		}
		position.y += 1
		*visitedHousesCount += visitHouse(*houseGrid, *position)
	case left:
		if position.x-1 < 0 {
			// add a new collumn to beginning of each line
			tmp := expandX(*houseGrid)
			for i := range *houseGrid {
				for j := range (*houseGrid)[i] {
					tmp[i][j+1] = (*houseGrid)[i][j]
				}
			}
			*houseGrid = tmp
			// reseting x to 1
			position.x += 1
            secondaryPosition.x += 1
		}
		position.x -= 1
		*visitedHousesCount += visitHouse(*houseGrid, *position)
	}
}

func main() {
	fileName := filepath.Join("..", "inputs.txt")
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	houseGrid := make([][]bool, 10)

	for idx := range houseGrid {
		houseGrid[idx] = make([]bool, 10)
		// setting start position to the middle of the matrix
		if idx == 4 {
			houseGrid[4][4] = true
		}
	}

	santaPosition := grid{4, 4}
	robotSantaPosition := grid{4, 4}
	character := make([]byte, 1)
	visitedHouses := 1
	isSantaTurn := true
	for _, err := file.Read(character); err != io.EOF; _, err = file.Read(character) {
		if isSantaTurn {
			calculateVisitedHouses(
                character,
                &visitedHouses,
                &houseGrid,
                &santaPosition,
                &robotSantaPosition,
            )
			isSantaTurn = false
			continue
		}

		calculateVisitedHouses(
            character, 
            &visitedHouses,
            &houseGrid,
            &robotSantaPosition,
            &santaPosition,
        )
		isSantaTurn = true
	}

	fmt.Println(visitedHouses)
}
