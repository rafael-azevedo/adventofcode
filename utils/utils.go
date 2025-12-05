package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n"), nil
}

func MustAtoi(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n, nil
}
