package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "day1_input.txt"
const goalFloor = -1

var currentFloor = 0
var entryPoint = 0

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	doWork(file, goalFloor)
	fmt.Println("Final Floor:", currentFloor)
	fmt.Printf("Entered Basement at: %d\n", entryPoint)
}

// read each line of the file, building the total and
//looking for the expected entry point
func doWork(file *os.File, goalFloor int) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	counter := 0
	entryFound := false
	for scanner.Scan() {
		counter += Count(scanner.Text())
		if !entryFound && currentFloor == goalFloor {
			entryFound = true
			entryPoint = counter
		}
	}
}

// Count returns the final floor from a string
func Count(x string) int {
	if x == "(" {
		currentFloor++
		return 1
	} else if x == ")" {
		currentFloor--
		return 1
	}
	return 0
}
