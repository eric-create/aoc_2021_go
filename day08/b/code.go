package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Number struct {
	Code   []rune
	Value  int
	Length int
}

// Sets the code, value, and length.
func NewNumber(code string, value int) Number {
	return Number{[]rune(code), value, len(code)}
}

// Sets the code and length.
func UnknownNumber(code string) Number {
	return Number{[]rune(code), -1, len(code)}
}

func (n Number) CodeContains(segments []rune) bool {
	for _, segment := range segments {
		if !slices.Contains(n.Code, segment) {
			return false
		}
	}

	return true
}

func (n Number) Difference(other Number) []rune {
	return mapset.NewSet(n.Code...).Difference(mapset.NewSet(other.Code...)).ToSlice()
}

func Numbers() []Number {
	return []Number{
		NewNumber("abcefg", 0),
		NewNumber("cf", 1),
		NewNumber("acdeg", 2),
		NewNumber("acdfg", 3),
		NewNumber("bcdf", 4),
		NewNumber("abdfg", 5),
		NewNumber("abdefg", 6),
		NewNumber("acf", 7),
		NewNumber("abcdefg", 8),
		NewNumber("abcdfg", 9),
	}
}

func NumberMap() map[int][]Number {
	numberMap := map[int][]Number{}

	for _, number := range Numbers() {
		numberMap[number.Length] = append(numberMap[number.Length], number)
	}

	return numberMap
}

func Numbers235() []Number {
	return NumberMap()[5]
}

func Numbers069() []Number {
	return NumberMap()[6]
}

func Get1748() []Number {
	numberMap := NumberMap()
	uniques := []Number{}

	for i := 0; i <= 9; i++ {
		if len(numberMap[i]) == 1 {
			uniques = append(uniques, numberMap[i]...)
		}
	}

	return uniques
}

func Identify1748(codes []string, knownNumbers *[10]Number) {
	for _, code := range codes {
		for _, number := range Get1748() {
			if number.Length == len(code) {
				number.Code = []rune(code)
				(*knownNumbers)[number.Value] = number
			}
		}
	}
}

func FilterCodes(codes []string, length int) []string {
	filtered := []string{}

	for _, code := range codes {
		if len(code) == length {
			filtered = append(filtered, code)
		}
	}

	return filtered
}

func Identify3(codes []string, numbers235 []Number, knownNumbers *[10]Number) {
	for _, code := range codes {
		one := knownNumbers[1]
		candidate := UnknownNumber(code)

		if candidate.CodeContains(one.Code) {
			candidate.Value = 3
			(*knownNumbers)[3] = candidate
		}
	}
}

func Identify6(codes []string, numbers069 []Number, knownNumbers *[10]Number) {
	for _, code := range codes {
		one := knownNumbers[1]
		candidate := UnknownNumber(code)

		if !candidate.CodeContains(one.Code) {
			candidate.Value = 6
			(*knownNumbers)[6] = candidate
		}
	}
}

func Identify09(codes []string, numbers []Number, knownNumbers *[10]Number, knownSegments *map[rune]rune) {
	maskCode := append(knownNumbers[4].Code, (*knownSegments)['a'])
	mask := UnknownNumber(string(maskCode))

	for _, code := range codes {

		// Skip known SIX
		if Detect(code, knownNumbers) != nil {
			continue
		}

		candidate := UnknownNumber(code)
		difference := mask.Difference(candidate)

		if len(difference) > 0 {
			// ZERO
			candidate.Value = 0
			(*knownNumbers)[0] = candidate

		} else {
			// NINE
			candidate.Value = 9
			(*knownNumbers)[9] = candidate
		}
	}
}

func Identify25(codes []string, numbers []Number, knownNumbers *[10]Number, knownSegments *map[rune]rune) {
	maskCode := append(knownNumbers[4].Code, (*knownSegments)['a'])
	mask := UnknownNumber(string(maskCode))

	for _, code := range codes {

		// Skip known THREE
		if Detect(code, knownNumbers) != nil {
			continue
		}

		candidate := UnknownNumber(code)
		difference := mask.Difference(candidate)

		if len(difference) == 2 {
			// TWO
			candidate.Value = 2
			(*knownNumbers)[2] = candidate

		} else {
			// FIVE
			candidate.Value = 5
			(*knownNumbers)[5] = candidate
		}
	}
}

// Checks to which known number the code belongs.
func Detect(code string, knownNumbers *[10]Number) *Number {
	for _, knownNumber := range knownNumbers {
		if knownNumber.Length != 0 {
			if SortCode(code) == SortCode(string(knownNumber.Code)) {
				return &knownNumber
			}
		}
	}

	return nil
}

func SortCode(code string) string {
	return string(utils.SortRuneSlice([]rune(code)))
}

func SegmentsMap() map[rune]rune {
	return map[rune]rune{
		'a': ' ',
		'b': ' ',
		'c': ' ',
		'd': ' ',
		'e': ' ',
		'f': ' ',
		'g': ' ',
	}
}

func IdentifyA(one, seven Number, segments *map[rune]rune) {
	a := seven.Difference(one)[0]
	(*segments)['a'] = a
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	// lines, _ := utils.ReadLines("extra-input.txt")
	entries := ParseEntries(*lines)
	sum := 0

	for _, entry := range entries {

		knownSegments := SegmentsMap()
		knownNumbers := [10]Number{}

		Identify1748(entry[0], &knownNumbers)
		// The segment "a" can be deduce from the known numbers ONE and SEVEN.
		IdentifyA(knownNumbers[1], knownNumbers[7], &knownSegments)

		codes235 := FilterCodes(entry[0], 5)
		Identify3(codes235, []Number{Numbers()[2], Numbers()[3], Numbers()[5]}, &knownNumbers)
		Identify25(codes235, []Number{Numbers()[2], Numbers()[5]}, &knownNumbers, &knownSegments)

		codes069 := FilterCodes(entry[0], 6)
		Identify6(codes069, []Number{Numbers()[0], Numbers()[6], Numbers()[9]}, &knownNumbers)
		Identify09(codes069, []Number{Numbers()[0], Numbers()[9]}, &knownNumbers, &knownSegments)

		outputValue := DecodeOutput(entry[1], &knownNumbers)
		sum += outputValue

		// PrintEntry(entry[0], &knownNumbers)
		// fmt.Println(outputValue)
	}
	fmt.Println(sum)
}

func PrintEntry(codes []string, knownNumbers *[10]Number) {
	for _, code := range codes {
		if number := Detect(code, knownNumbers); number != nil {
			fmt.Print(number.Value, " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func PrintRuneMap(segments map[rune]rune) {
	for k, v := range segments {
		fmt.Println(string(k), ":", string(v))
	}
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

func DecodeOutput(codes []string, knownNumbers *[10]Number) int {
	outputValueString := ""

	for i := 0; i < 4; i++ {
		number := Detect(codes[i], knownNumbers)
		outputValueString += strconv.Itoa(number.Value)
	}

	outputValue, _ := strconv.Atoi(outputValueString)
	return outputValue
}
