package utils

import (
	"os"
	"strconv"
	"strings"
)

// Reads the lines of a file. Error is handed over.
func ReadLines(path string) (*[]string, error) {
	if content, err := os.ReadFile(path); err != nil {
		return nil, err

	} else {
		lines := strings.Split(string(content), "\n")
		return &lines, nil
	}
}

func StringSliceToInt(slice []string) (*[]int, error) {
	intSlice := []int{}

	for _, s := range slice {
		if value, err := strconv.Atoi(s); err != nil {
			return nil, err

		} else {
			intSlice = append(intSlice, value)
		}
	}

	return &intSlice, nil
}

func IntSliceSum(slice []int) int {
	sum := 0

	for _, i := range slice {
		sum += i
	}

	return sum
}
