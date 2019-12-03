package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const input = "yzbqklnj"
const target = "000000"

func main() {
	extension := findHashExtension(input, target)
	fmt.Println("Extension is:", extension)
}

func findHashExtension(input string, targetString string) int {
	counter := 1
	messages := make(chan string)

	for {
		go func(counter int, input string) {
			currentTest := input + strconv.Itoa(counter)
			hasher := md5.New()
			hasher.Write([]byte(currentTest))
			sum := hasher.Sum(nil)
			sumString := hex.EncodeToString(sum)
			firstFive := sumString[0:len(targetString)]
			messages <- firstFive
			// if firstFive == "00000" {
			// 	messages <- counter
			// }
		}(counter, input)

		msg := <-messages
		if msg == targetString {
			return counter
		}
		counter++
	}
}
