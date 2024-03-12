package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	numbers := Numbers(*lines)

	common := Common(numbers)
	notCommon := Not(common)

	commonDecimal := Decimal(common)
	notCommonDecimal := Decimal(notCommon)

	fmt.Println(common, commonDecimal, notCommon, notCommonDecimal)
	fmt.Println(commonDecimal * notCommonDecimal)
}

func Decimal(number []int) int {
	dec := 0

	for i := 0; i < len(number); i++ {
		dec += number[len(number)-1-i] * int(math.Pow(2, float64(i)))
	}

	return dec
}

func Common(numbers [][]int) []int {
	max := len(numbers)
	numberLen := len(numbers[0])
	common := []int{}

	for i := 0; i < numberLen; i++ {
		count := 0

		for _, number := range numbers {
			count += number[i]
		}

		if max-count < count {
			common = append(common, 1)
		} else if max-count > count {
			common = append(common, 0)
		} else {
			panic("oh no")
		}
	}

	return common
}

func Numbers(lines []string) [][]int {
	numbers := [][]int{}

	for _, line := range lines {
		numbers = append(numbers, Number(line))
	}

	return numbers
}

func Number(line string) []int {
	number := []int{}

	for _, r := range line {
		bit, _ := strconv.Atoi(string(r))
		number = append(number, bit)
	}

	return number
}

func Not(number []int) []int {
	notted := []int{}

	for _, digit := range number {
		if digit == 1 {
			notted = append(notted, 0)
		} else {
			notted = append(notted, 1)
		}
	}

	return notted
}
