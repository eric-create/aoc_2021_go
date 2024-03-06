package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	measurements, _ := utils.StringSliceToInt(*lines)

	windows := []int{}
	increaseCount := 0

	for i := 0; i <= len(*measurements)-3; i++ {
		window := utils.IntSliceSum((*measurements)[i : i+3])

		if i > 0 {
			if window > windows[i-1] {
				increaseCount++
			}
		}

		windows = append(windows, window)
	}

	fmt.Println(increaseCount)
}
