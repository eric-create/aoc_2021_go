package main

import (
	"eric-create/aoc_2021/nodes"
	"eric-create/aoc_2021/utils"
	"fmt"
	"slices"
)

const (
	CALLED = "called"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	calls := utils.ExtractInts((*lines)[0])
	paragraphs := utils.SplitParagraphs((*lines)[2:])
	boards := Boards(paragraphs, calls)

	for _, board := range boards {
		PrintBoard(board)
	}
}

func PrintBoard(board [][]*nodes.Node) {
	for _, nodes := range board {
		for _, node := range nodes {
			if slices.Contains(node.Tags, CALLED) {
				s := node.Symbol
				if len(s) == 1 {
					s = s + " "
				}
				fmt.Print(s, " ")
			} else {
				fmt.Print(".. ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Boards(paragraphs [][]string, calls []int) [][][]*nodes.Node {
	boards := [][][]*nodes.Node{}

	for _, paragraph := range paragraphs {
		intsSlice := ParagraphsToIntsSlice(paragraph)
		board := nodes.IntsToField(intsSlice)
		MarkCalls(&board, calls)
		boards = append(boards, board)
	}

	return boards
}

func ParagraphsToIntsSlice(paragraph []string) [][]int {
	intsSlice := [][]int{}

	for i, line := range paragraph {
		intsSlice = append(intsSlice, []int{})
		ints := utils.ExtractInts(line)
		intsSlice[i] = append(intsSlice[i], ints...)
	}

	return intsSlice
}

func MarkCalls(board *[][]*nodes.Node, calls []int) {
	for _, nodes := range *board {
		for _, node := range nodes {
			if slices.Contains(calls, node.SymbolToInt()) {
				node.Tags = append(node.Tags, CALLED)
			}
		}
	}
}
