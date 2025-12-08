package main

import (
	"strconv"
)

func part2(lines []string) int {
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
			if location-distance <= 0 {
				total += (distance - location) / 100
				if location != 0 {
					total += 1
				}
			}
			location = (location - distance) % 100
		case "R":
			total += (location + distance) / 100
			location = (location + distance) % 100
		default:
		}

		switch {
		case location < 0:
			location += 100
		}

	}

	return total
}
