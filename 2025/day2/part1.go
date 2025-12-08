package main

import (
	"strconv"
	"strings"
)

func part1(lines []string) int {
	total := 0
	ids := strings.Split(lines[0], ",")
	for _, id := range ids {
		parts := strings.Split(id, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			numString := strconv.Itoa(i)
			if len(numString)%2 == 0 {
				numString := strconv.Itoa(i)
				left, _ := strconv.Atoi(numString[0:(len(numString) / 2)])
				right, _ := strconv.Atoi(numString[(len(numString) / 2):])
				if left == right {
					total += i
				}
			}
		}
	}
	return total
}
