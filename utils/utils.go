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

func ExtractInts(s string) []int {
	numbers := [][]rune{}
	lastDigit := -2

	for i, r := range s {
		if _, err := strconv.Atoi(string(r)); err == nil {
			if i-lastDigit > 1 {
				numbers = append(numbers, []rune{r})
			} else {
				numbers[len(numbers)-1] = append(numbers[len(numbers)-1], r)
			}

			lastDigit = i
		}
	}

	ints := []int{}
	for _, number := range numbers {
		_int, _ := strconv.Atoi(string(number))
		ints = append(ints, _int)
	}

	return ints
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

func SplitParagraphs(lines []string) [][]string {
	paragraphs := [][]string{}
	paragraph := []string{}

	for _, line := range lines {
		if line == "" {
			paragraphs = append(paragraphs, paragraph)
			paragraph = []string{}
		} else {
			paragraph = append(paragraph, line)
		}
	}

	return paragraphs
}

func NormalizeInt(i int) int {
	if i > 0 {
		return 1
	} else if i == 0 {
		return 0
	}
	return -1
}
