package main

import (
	"strconv"
	"strings"
)

func part2(lines []string) int {
	total := 0
	ids := strings.Split(lines[0], ",")
	for _, id := range ids {
		parts := strings.Split(id, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			numString := strconv.Itoa(i)
			if !isValid(numString) {
				total += i
			}
		}
	}
	return total
}

func isValid(s string) bool {
	doubled := s + s
	return !strings.Contains(doubled[1:len(doubled)-1], s)
}
