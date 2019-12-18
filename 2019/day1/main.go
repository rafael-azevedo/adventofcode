package main

import (
	"fmt"
	"os"

	"github.com/rafael-azevedo/adventofcode/utils"
)

func main() {
	masses, err := utils.ReadLinestoInt(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 :: %d \n", CalcFuel(masses...))

	fmt.Printf("Part 2 :: %d \n", calcFuelAlso(masses...))

	// Recursive part 2
	totalfuel := 0
	for _, m := range masses {
		_, f := calcFuelRecursive(m, 0)
		totalfuel += f
	}

	fmt.Printf("Part 2 recursive :: %d \n", totalfuel)
}

func CalcFuel(masses ...int) int {
	totalFuel := 0
	for _, m := range masses {
		totalFuel = totalFuel + ((m / 3) - 2)
	}
	return totalFuel
}

func calcFuelAlso(masses ...int) int {
	totalFuel := 0
	for _, m := range masses {
		fuel := (m / 3) - 2
		for fuel > 0 {
			totalFuel = totalFuel + fuel
			fuel = (fuel / 3) - 2
		}
	}
	return int(totalFuel)
}

func calcFuelRecursive(mass, fuel int) (int, int) {
	diff := (mass / 3) - 2

	if diff <= 0 {
		return mass, fuel
	}
	fuel += diff

	return calcFuelRecursive(diff, fuel)
}
