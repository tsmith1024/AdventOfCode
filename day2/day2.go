package day2

import "fmt"

func main() {
	fmt.Printf("hello, world")
}

// WrappingArea calculates the total wrapping needed for a single package
func WrappingArea(dimensions []int) int {
	frontArea := dimensions[0] * dimensions[1]
	sideArea := dimensions[1] * dimensions[2]
	topArea := dimensions[2] * dimensions[0]
	return (frontArea + sideArea + topArea) * 2
}
