package main

import (
	"eric-create/aoc_2021/nodes"
	"eric-create/aoc_2021/utils"
	"eric-create/aoc_2021/vectors"
	"fmt"
)

const (
	LOWEST = "lowest"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	field := Field(*lines)
	lows := TagLows(&field)
	// nodes.PrintFieldWithTag(field, LOWEST)

	riskLevel := 0

	for _, low := range lows {
		riskLevel += low.SymbolToInt() + 1
	}

	fmt.Println(riskLevel)
}

func TagLows(field *[][]*nodes.Node) []*nodes.Node {
	lows := []*nodes.Node{}

	for _, row := range *field {
		for _, node := range row {
			isLowest := true

			for _, neighbor := range node.GetNeighbors(vectors.ManhattanDirections()) {
				if neighbor.SymbolToInt() <= node.SymbolToInt() {
					isLowest = false
				}
			}

			if isLowest {
				node.Tags = append(node.Tags, LOWEST)
				lows = append(lows, node)
			}
		}
	}

	return lows
}

func Field(lines []string) [][]*nodes.Node {
	intField := [][]int{}

	for y, line := range lines {
		intField = append(intField, []int{})

		for _, char := range line {
			intField[y] = append(intField[y], utils.RuneToInt(char))
		}
	}

	nodeField := nodes.IntsToField(intField)
	return nodeField
}
