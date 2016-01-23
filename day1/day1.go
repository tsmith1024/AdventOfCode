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
	total, entryPoint := doWork(file, goalFloor)
	fmt.Println("Final Floor:", total)
	fmt.Printf("Entered Basement at: %d\n", entryPoint)
}

// read each line of the file, building the total and looking for the expected
// entry point
func doWork(file *os.File, goalFloor int) (int, int) {
	scanner := bufio.NewScanner(file)
	total := 0
	entryPoint := 0
	entryFound := false
	for scanner.Scan() {
		line := scanner.Text()
		if !entryFound {
			index := 0
			index, entryFound = FindEntryPoint(line, goalFloor)
			entryPoint += index
		}
		total += Count(line)
	}
	return total, entryPoint
}

// Count returns the final floor from a string
func Count(x string) int {
	return strings.Count(x, "(") - strings.Count(x, ")")
}

// FindEntryPoint returns the first step at which a given floor is reached
func FindEntryPoint(input string, floor int) (int, bool) {
	total := 0
	for index, runeValue := range input {
		strVal := string(runeValue)
		if strVal == "(" {
			total++
		} else if strVal == ")" {
			total--
		}
		if total == floor {
			return index + 1, true
		}
	}
	return len(input), false
}
