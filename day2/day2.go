package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const fileName = "day2_input.txt"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	totalWrap := 0
	totalRibbon := 0
	for scanner.Scan() {
		dimensions := strings.Split(scanner.Text(), "x")
		length, width, height := floatDimensions(dimensions)
		totalWrap += wrappingArea(length, width, height)
		totalRibbon += packageRibbon(length, width, height)
	}
	fmt.Println("Total wrap needed is:", totalWrap)
	fmt.Println("Total Ribbon is:", totalRibbon)
}

// wrappingArea calculates the total wrapping needed for a single package
func wrappingArea(length float64, width float64, height float64) int {
	frontArea := length * width
	sideArea := width * height
	topArea := height * length
	bonusArea := math.Min(frontArea, math.Min(sideArea, topArea))
	return int(2*(frontArea+sideArea+topArea) + bonusArea)
}

func packageRibbon(length float64, width float64, height float64) int {
	largestSide := math.Max(length, math.Max(width, height))
	perimeter := 2 * (length + width + height - largestSide)
	volume := (length * width * height)
	return int(perimeter + volume)
}

func floatDimensions(dimensions []string) (float64, float64, float64) {
	length, err := strconv.ParseFloat(dimensions[0], 64)
	width, err := strconv.ParseFloat(dimensions[1], 64)
	height, err := strconv.ParseFloat(dimensions[2], 64)
	if err != nil {
		log.Fatal(err)
	}
	return length, width, height
}
