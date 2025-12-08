package main

import (
	"fmt"
	"strconv"
)

func part1(lines []string) int {
	location := 50
	total := 0

	for _, line := range lines {
		if line == "" {
			break
		}

		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])
		switch direction {
		case "L":
			location = (location - distance) % 100
		case "R":
			location = (location + distance) % 100
		default:
		}

		switch {
		case location < 0:
			location += 100
		case location == 0:
			total++
		}

		fmt.Printf("The dial is rotated %s to point at %d\n", line, location)
	}

	return total
}
