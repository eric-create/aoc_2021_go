package main

import (
	"eric-create/aoc_2021/nodes"
	"eric-create/aoc_2021/utils"
	"eric-create/aoc_2021/vectors"
	"fmt"
	"slices"
	"sort"
)

const (
	LOWEST = "lowest"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	field := Field(*lines)
	lows := TagLows(&field)
	// nodes.PrintFieldWithTag(field, LOWEST)

	basins := GetBasins(lows)
	riskLevel := ThreeLargestBasins(basins)

	fmt.Println(riskLevel)
}

func ThreeLargestBasins(basins [][]*nodes.Node) int {
	sizes := []int{}

	for _, basin := range basins {
		sizes = append(sizes, len(basin))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	return sizes[0] * sizes[1] * sizes[2]
}

func GetBasins(lows []*nodes.Node) [][]*nodes.Node {
	basins := [][]*nodes.Node{}

	for _, low := range lows {
		basins = append(basins, GetBasin(low, []*nodes.Node{low}))
	}

	return basins
}

func GetBasin(node *nodes.Node, basin []*nodes.Node) []*nodes.Node {
	neighbors := node.GetNeighbors(vectors.ManhattanDirections())

	for _, neighbor := range neighbors {
		if neighbor.Symbol != "9" {
			if !slices.Contains(basin, neighbor) {
				basin = GetBasin(neighbor, append(basin, neighbor))
			}
		}
	}

	return basin
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
