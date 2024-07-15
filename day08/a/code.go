package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
	"strings"
)

type Number struct {
	Segments string
	Literal  string
	Length   int
}

func NewNumber(segments, literal string) Number {
	return Number{segments, literal, len(segments)}
}

func Numbers() []Number {
	numberLiterals := utils.NumberLiterals()

	return []Number{
		NewNumber("abcefg", numberLiterals[0]),
		NewNumber("cf", numberLiterals[1]),
		NewNumber("acdeg", numberLiterals[2]),
		NewNumber("acdfg", numberLiterals[3]),
		NewNumber("bcdf", numberLiterals[4]),
		NewNumber("abdfg", numberLiterals[5]),
		NewNumber("abdefg", numberLiterals[6]),
		NewNumber("acf", numberLiterals[7]),
		NewNumber("abcdefg", numberLiterals[8]),
		NewNumber("abcdfg", numberLiterals[9]),
	}
}

func NumberMap() map[int][]Number {
	numberMap := map[int][]Number{}

	for _, number := range Numbers() {
		numberMap[number.Length] = append(numberMap[number.Length], number)
	}

	return numberMap
}

func Uniques() []Number {
	numberMap := NumberMap()
	uniques := []Number{}

	for i := 0; i <= 9; i++ {
		if len(numberMap[i]) == 1 {
			uniques = append(uniques, numberMap[i]...)
		}
	}

	return uniques
}

func (n Number) Detect(segments string) bool {
	return n.Length == len(segments)
}

func DetectUniques(segments string) *Number {
	for _, unique := range Uniques() {
		if unique.Detect(segments) {
			return &unique
		}
	}

	return nil
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	entries := ParseEntries(*lines)

	count := 0

	for i, entry := range entries {
		fmt.Print("Entry ", i, ": ")
		for _, symbol := range entry[1] {
			if number := DetectUniques(symbol); number != nil {
				count++
				fmt.Print(symbol, " ")
			}
		}
		fmt.Println()
	}

	fmt.Println(count)
}

func ParseEntries(lines []string) [][2][]string {
	entries := [][2][]string{}

	for _, line := range lines {
		entries = append(entries, ParseEntry(line))
	}

	return entries
}

func ParseEntry(line string) [2][]string {
	symbols := strings.Split(line, " ")

	entry := [2][]string{}
	entry[0] = symbols[0:10]
	entry[1] = symbols[11:]
	return entry
}
