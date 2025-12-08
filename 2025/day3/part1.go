package main

import (
	"strconv"
)

func part1(lines []string) int {
	total := 0

	for _, line := range lines {
		first := 0
		second := 0
		for pos, joltS := range line {
			joltInt, _ := strconv.Atoi(string(joltS))

			if first == 0 {
				first = joltInt
				continue
			}
			if pos < len(line)-1 {
				if first < joltInt {
					first = joltInt
					second = 0
					continue
				}
				if second < joltInt {
					second = joltInt
					continue
				}
			}
			if second < joltInt {
				second = joltInt
			}
		}
		tString := strconv.Itoa(first) + strconv.Itoa(second)
		tInt, _ := strconv.Atoi(tString)
		total += tInt
	}
	return total
}
