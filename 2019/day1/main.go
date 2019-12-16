package main

import (
	"fmt"
	"github.com/rafael-azevedo/adventofcode/utils"
	"os"
)

func main() {
	masses, err := utils.ReadLinestoInt(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 :: %d \n", CalcFuel(masses...))



	fmt.Printf("Part 2 :: %d \n", CalcFuel(masses...))
}


func CalcFuel(masses... int) int {
	totalFuel := 0
	for _, m := range masses{
		totalFuel = totalFuel + ((m / 3) - 2)
	}
	return totalFuel
}

