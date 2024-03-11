package utils

import (
	"os"
	"regexp"
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

// Only use this if the string `s` contains only one integer value.
func ExtractInt(s string) (int, error) {
	digits := []rune{}

	for _, r := range s {
		if _, err := strconv.Atoi(string(r)); err == nil {
			digits = append(digits, r)
		}
	}

	return strconv.Atoi(string(digits))
}

func DescriptorInt(descriptor, s string, delimiter *string) (int, error) {
	re := regexp.MustCompile(descriptor)
	indexes := re.FindStringIndex(s)
	valueStart := indexes[1] + 1

	if delimiter == nil {
		return strconv.Atoi(s[valueStart:])
	}

	var valueEnd int

	for i := valueStart; i < len(s); i++ {
		if strings.HasPrefix(s[i:], *delimiter) {
			valueEnd = i - 1
			break
		}
	}

	return strconv.Atoi(s[valueStart : valueEnd+1])
}
