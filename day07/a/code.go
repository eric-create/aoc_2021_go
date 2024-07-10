package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
	"math"
	"sort"
)

func main() {
	lines, _ := utils.ReadLines("Input.txt")
	numbers := utils.ExtractInts((*lines)[0])
	sort.Ints(numbers)
	count := len(numbers)

	median_1 := Cost(numbers, numbers[count/2])
	median_2 := Cost(numbers, numbers[count/2+1])

	fmt.Println(count/2, median_1)
	fmt.Println(count/2+1, median_2)
}

func Cost(numbers []int, position int) int {
	cost := 0

	for _, n := range numbers {
		cost += int(math.Abs(float64(position - n)))
	}

	return cost
}
