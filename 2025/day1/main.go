package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

func main() {
	input, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))

}
