package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	measurements, _ := utils.StringSliceToInt(*lines)

	increaseCount := 0

	for i, measurement := range (*measurements)[1:] {
		// Because we start at slice index 1, i points to index-1
		if measurement > (*measurements)[i] {
			increaseCount++
		}
	}

	fmt.Println(increaseCount)
}
