package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLinestoInt(FilePath string) ([]int, error){
	file, err := os.Open(FilePath)
	if err != nil {
		return []int{} , err
	}
	defer file.Close()

	var result []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		int, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, int)
	}
	return result, scanner.Err()
}