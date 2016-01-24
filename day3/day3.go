package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "day3_input.txt"

type house struct {
	x int
	y int
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	total := 0
	dualTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += houseCount(line)
		dualTotal += dualWork(line)
	}
	fmt.Println("Total Houses:", total)
	fmt.Println("Dual Houses:", dualTotal)
}

func houseCount(input string) int {
	splitInput := strings.Split(input, "")
	currentHouse := house{}
	houseMap := make(map[house]int)
	houseMap[currentHouse] = 0
	for _, letter := range splitInput {
		nextAddress := nextHouse(currentHouse, letter)
		houseMap[nextAddress] = 0
		currentHouse = nextAddress
	}
	return len(houseMap)
}

func dualWork(input string) int {
	splitInput := strings.Split(input, "")
	santaHouse := house{}
	roboHouse := house{}
	houseMap := make(map[house]int)
	houseMap[santaHouse] = 0
	for index, letter := range splitInput {
		if index%2 == 0 {
			nextAddress := nextHouse(santaHouse, letter)
			houseMap[nextAddress] = 0
			santaHouse = nextAddress
		} else {
			nextAddress := nextHouse(roboHouse, letter)
			houseMap[nextAddress] = 0
			roboHouse = nextAddress
		}
	}

	return len(houseMap)
}

func nextHouse(currentHouse house, direction string) house {
	newHouse := currentHouse
	switch direction {
	case ">":
		newHouse.x++
	case "<":
		newHouse.x--
	case "^":
		newHouse.y++
	case "v":
		newHouse.y--
	}
	return newHouse
}
