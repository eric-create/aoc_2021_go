package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	numbers := utils.ExtractInts((*lines)[0])
	pools := Populate(numbers, [10]int{})
	pools = CountDown(256, pools)
	fmt.Println(FishSum(pools))
}

func CountDown(days int, pools [10]int) [10]int {
	for day := 0; day < days; day++ {

		// Shift zeroes to temporary pool (index 10)
		pools[9] = pools[0]

		// Shift everybody else to lower time pool
		for i := 0; i < 8; i++ {
			pools[i] = pools[i+1]
		}

		pools[6] = pools[6] + pools[9]
		pools[8] = pools[9]
		pools[9] = 0
	}

	return pools
}

func FishSum(pools [10]int) int {
	sum := 0

	for _, pool := range pools {
		sum += pool
	}

	return sum
}

func Populate(numbers []int, pools [10]int) [10]int {
	for _, n := range numbers {
		pools[n]++
	}

	return pools
}
