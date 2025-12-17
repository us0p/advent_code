package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrInvalidDimension = errors.New("Invalid Dimension received")
	ErrInvalidNumber    = errors.New("Invalid Number received")
)

type RightRectangularPrism struct {
	l int
	w int
	h int
}

func (d RightRectangularPrism) SurfaceArea() int {
	return 2*d.l*d.w + 2*d.w*d.h + 2*d.l*d.h
}

func (d RightRectangularPrism) SmallestArea() int {
	return min(
		d.l*d.w,
		d.w*d.h,
		d.l*d.h,
	)
}

func (d RightRectangularPrism) SmallestPerimeter() int {
	return min(
		2*d.l+2*d.w,
		2*d.w+2*d.h,
		2*d.h+2*d.l,
	)
}

func (d RightRectangularPrism) CubicVolume() int {
	return d.l * d.w * d.h
}

func NewRightRectangularPrism(l, w, h int) RightRectangularPrism {
	return RightRectangularPrism{
		l,
		w,
		h,
	}
}

func extractRightRectangularPrism(input string) (RightRectangularPrism, error) {
	d := strings.Split(input, "x")
	if len(d) != 3 {
		return RightRectangularPrism{}, fmt.Errorf(
			"%w: expected 3 values, got %d, (%q)",
			ErrInvalidDimension,
			len(d),
			input,
		)
	}
	l, err := strconv.Atoi(d[0])
	if err != nil {
		return RightRectangularPrism{}, fmt.Errorf(
			"%w: length=%q: %v",
			ErrInvalidNumber,
			d[0],
			err,
		)
	}
	w, err := strconv.Atoi(d[1])
	if err != nil {
		return RightRectangularPrism{}, fmt.Errorf(
			"%w: width=%q: %v",
			ErrInvalidNumber,
			d[1],
			err,
		)
	}
	h, err := strconv.Atoi(d[2])
	if err != nil {
		return RightRectangularPrism{}, fmt.Errorf(
			"%w: height=%q: %v",
			ErrInvalidNumber,
			d[2],
			err,
		)
	}

	return NewRightRectangularPrism(l, w, h), nil
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error while reading input file", err)
		return
	}
	dimensionsString := string(input)
	dimensions := strings.Split(dimensionsString[:len(dimensionsString)-1], "\n")

	totalWrappingPapper, err := PartOne(dimensions)
	if err != nil {
		fmt.Println("Error while running PartOne", err)
		return
	}
	fmt.Println("Part One response is:", totalWrappingPapper)

	totalRibbon, err := PartTwo(dimensions)
	if err != nil {
		fmt.Println("Error while running PartTwo", err)
		return
	}
	fmt.Println("Part Two response is:", totalRibbon)
}

func PartOne(dimensions []string) (int, error) {
	total := 0
	for _, d := range dimensions {
		dimension, err := extractRightRectangularPrism(d)
		if err != nil {
			return 0, fmt.Errorf(
				"invalid dimension %q: %w",
				d,
				err,
			)
		}

		total += dimension.SurfaceArea() + dimension.SmallestArea()
	}
	return total, nil
}

func PartTwo(dimensions []string) (int, error) {
	total := 0
	for _, d := range dimensions {
		dimension, err := extractRightRectangularPrism(d)
		if err != nil {
			return 0, fmt.Errorf(
				"invalid dimension %q: %w",
				d,
				err,
			)
		}
		total += dimension.SmallestPerimeter() + dimension.CubicVolume()
	}
	return total, nil
}
