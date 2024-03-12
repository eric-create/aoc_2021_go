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

	oxygen := FilterCommon(numbers)
	co2 := FilterUncommon(numbers)

	fmt.Println(oxygen, co2)
	fmt.Println(Decimal(oxygen[0]) * Decimal(co2[0]))
}

func Decimal(number []int) int {
	dec := 0

	for i := 0; i < len(number); i++ {
		dec += number[len(number)-1-i] * int(math.Pow(2, float64(i)))
	}

	return dec
}

func FilterCommon(numbers [][]int) [][]int {
	numberLen := len(numbers[0])

	for i := 0; i < numberLen; i++ {
		if len(numbers) == 1 {
			return numbers
		}

		commonDigit := CommonDigit(numbers, i, 1)
		selection := [][]int{}

		for j := range numbers {

			if numbers[j][i] == commonDigit {
				selection = append(selection, numbers[j])
			}
		}

		numbers = selection
	}

	return numbers
}

func FilterUncommon(numbers [][]int) [][]int {
	numberLen := len(numbers[0])

	for i := 0; i < numberLen; i++ {
		if len(numbers) == 1 {
			return numbers
		}

		commonDigit := CommonDigit(numbers, i, 1)
		selection := [][]int{}

		for j := range numbers {

			if numbers[j][i] != commonDigit {
				selection = append(selection, numbers[j])
			}
		}

		numbers = selection
	}

	return numbers
}

func CommonDigit(numbers [][]int, digit, decider int) int {
	max := len(numbers)
	count := 0

	for _, number := range numbers {
		count += number[digit]
	}

	if max-count < count {
		return 1
	} else if max-count > count {
		return 0
	} else {
		return decider
	}
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
