package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getMetricInts(metric string) ([]int, error) {
	metricsAsStr := strings.Split(metric, "x")
	metricsAsInt := make([]int, len(metricsAsStr))
	for idx, metricStr := range metricsAsStr {
		metricInt, err := strconv.Atoi(metricStr)
		if err != nil {
			return metricsAsInt, err
		}
		metricsAsInt[idx] = metricInt
	}

	return metricsAsInt, nil
}

func getSmallest(l, w, h int) int {
	return min(l, min(w, h))
}

func calculateRectangularWrappingPapper(metric []int) int {
	baseArea := metric[0] * metric[1]
	sideArea := metric[1] * metric[2]
	frontArea := metric[2] * metric[0]

	squareFeetPapperForMetric := 2*baseArea + 2*sideArea + 2*frontArea

	return squareFeetPapperForMetric + getSmallest(baseArea, sideArea, frontArea)
}

func calculateRibbonWrappingPapper(metric []int) int {
	length := metric[0]
	width := metric[1]
	height := metric[2]

    basePerimeter := 2 * (length + width)
    sidePerimeter := 2 * (height + width)
    frontPerimeter := 2 * (height + length)

	wrappingPapper := getSmallest(basePerimeter, sidePerimeter, frontPerimeter)
	bow := length * width * height

	return wrappingPapper + bow
}

func getTotalSquareFeetOfPapper(metricList []string, calculateFn func([]int) int) (int, error) {
	totalSquareFeet := 0

	for _, metric := range metricList {
		if metric == "" {
			continue
		}
		integerMetric, err := getMetricInts(metric)

		if err != nil {
			return 0, err
		}

		totalSquareFeet += calculateFn(integerMetric)
	}

	return totalSquareFeet, nil
}

func main() {
	filename := filepath.Join("..", "inputs.txt")
	inputs, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	elvesList := strings.Split(string(inputs), "\n")

	totalSquareFeetRectangle, err := getTotalSquareFeetOfPapper(elvesList, calculateRectangularWrappingPapper)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total square feet for rectangles:", totalSquareFeetRectangle)

	totalSquareFeetRibbon, err := getTotalSquareFeetOfPapper(elvesList, calculateRibbonWrappingPapper)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total square feet for ribbons:", totalSquareFeetRibbon)
}
