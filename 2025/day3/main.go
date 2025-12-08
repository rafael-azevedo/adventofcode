package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

func main() {
	// // Uncomment to test with example
	// input, _ := utils.ReadLines("example.txt")

	//Real inputs
	input, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))

}
