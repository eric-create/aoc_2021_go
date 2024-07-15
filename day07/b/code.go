package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
	"math"
)

func main() {
	lines, _ := utils.ReadLines("Input.txt")
	numbers := utils.ExtractInts((*lines)[0])

	// Mean cost
	mean := Mean(numbers)
	roundedMean := int(math.Round(mean))
	cost, costDeviations := Cost(numbers, roundedMean)

	// Mean alt cost
	roundedMeanAlt := roundedMean - 1
	altCost, _ := Cost(numbers, roundedMeanAlt)

	fmt.Println()
	// fmt.Println(numbers)
	fmt.Println("Mean       :", mean)
	fmt.Println("Cost       :", roundedMean, cost, costDeviations)
	fmt.Println("Cost (alt) :", roundedMeanAlt, altCost)
}

func Mean(numbers []int) float64 {
	count := len(numbers)
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return float64(sum) / float64(count)
}

func Cost(numbers []int, mean int) (int, []int) {
	deviations := []int{}
	cost := 0

	for _, n := range numbers {
		absDiff := utils.AbsDiff(mean, n)
		curCost := 0

		for d := 0; d <= absDiff; d++ {
			curCost += d
		}

		cost += curCost
		// deviations = append(deviations, curCost)
	}

	return cost, deviations
}
