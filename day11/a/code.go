package main

import (
	"eric-create/aoc_2021/nodes"
	"eric-create/aoc_2021/utils"
	"eric-create/aoc_2021/vectors"
	"fmt"
	"strconv"
)

const (
	FLASHED = "flashed"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	field := Field(*lines)
	flashes := Steps(&field, 100)
	// nodes.PrintField(field)
	fmt.Println(flashes)
}

func Steps(field *[][]*nodes.Node, stepsCount int) int {
	flashes := 0

	for i := 1; i <= stepsCount; i++ {
		stepFlashes := Step(field)
		fmt.Println("Step", i, ":", stepFlashes)
		flashes += stepFlashes
	}

	return flashes
}

func Step(field *[][]*nodes.Node) int {
	IncreaseField(field)
	FlashField(field)
	flashes := Reset(field)
	return flashes
}

func IncreaseField(field *[][]*nodes.Node) {
	for _, row := range *field {
		for _, node := range row {
			Increase(node)
		}
	}
}

func Increase(node *nodes.Node) {
	node.Symbol = strconv.Itoa(node.SymbolToInt() + 1)
}

func FlashField(field *[][]*nodes.Node) {
	for _, row := range *field {
		for _, node := range row {
			Flash(node)
		}
	}
}

func Flash(node *nodes.Node) {
	if node.SymbolToInt() > 9 && len(node.Tags) == 0 {
		node.Tags = append(node.Tags, FLASHED)
		neighbors := node.GetNeighbors(vectors.AllDirections())

		for _, neighbor := range neighbors {
			Increase(neighbor)
			Flash(neighbor)
		}
	}
}

func Reset(field *[][]*nodes.Node) int {
	flashes := 0

	for _, row := range *field {
		for _, node := range row {
			if node.SymbolToInt() > 9 {
				flashes++
				node.Symbol = "0"
				node.Tags = []string{}
			}
		}
	}

	return flashes
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
