package utils

import (
	"os"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n"), nil
}
