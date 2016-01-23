package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "day1_input.txt"
const goalFloor = -1

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	finalFloor, entryPoint := doWork(file, goalFloor)
	fmt.Println("Final Floor:", finalFloor)
	fmt.Printf("Entered Basement at: %d\n", entryPoint)
}

// read each line of the file, building the total and
//looking for the expected entry point
func doWork(file *os.File, goalFloor int) (int, int) {
	scanner := bufio.NewScanner(file)
	entryFound := false
	currentFloor := 0
	entryPoint := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !entryFound {
			index := 0
			entryFound, index = findFloor(line, currentFloor, goalFloor)
			entryPoint += index
		}
		currentFloor += count(line)
	}
	return currentFloor, entryPoint
}

// Count takes a string and returns the change in floors on that line
func count(line string) int {
	return strings.Count(line, "(") - strings.Count(line, ")")
}

func findFloor(line string, startingFloor int, goalFloor int) (bool, int) {
	letters := strings.Split(line, "")
	floor := startingFloor
	for char := range letters {
		floor += count(letters[char])
		if floor == goalFloor {
			return true, char + 1
		}
	}
	return false, len(line)
}
