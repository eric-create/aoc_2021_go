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

	for _, call := range calls {
		for _, board := range boards {
			if board.CheckCall(call) {
				if board.CheckWin() {
					fmt.Println(call, "x", board.BoardSum(), "=", call*board.BoardSum())
					return
				}
			}
		}
	}
	fmt.Println("Error")
}

type Board struct {
	Graph *[][]*nodes.Node
	Calls *[]*nodes.Node
}

func NewBoard(paragraph []string) *Board {
	intsSlice := ParagraphsToIntsSlice(paragraph)
	graph := nodes.IntsToField(intsSlice)

	return &Board{Graph: &graph, Calls: &[]*nodes.Node{}}
}

func (b *Board) CheckCall(call int) bool {
	for y := 0; y <= 4; y++ {
		for x := 0; x <= 4; x++ {
			if (*b.Graph)[y][x].SymbolToInt() == call {
				*b.Calls = append(*b.Calls, (*b.Graph)[y][x])
				return true
			}
		}
	}
	return false
}

func (b *Board) CallCount() int {
	return len(*b.Calls)
}

func (b *Board) CheckWin() bool {
	if b.CallCount() >= 5 {
		callMap := [2][5]int{}

		for _, node := range *b.Calls {
			callMap[0][node.Position.X]++
			callMap[1][node.Position.Y]++
			if callMap[0][node.Position.X] == 5 || callMap[1][node.Position.Y] == 5 {
				return true
			}
		}
	}
	return false
}

func (b *Board) BoardSum() int {
	sum := 0

	for y := 0; y <= 4; y++ {
		for x := 0; x <= 4; x++ {
			node := (*b.Graph)[y][x]

			if !slices.Contains(*b.Calls, node) {
				sum += node.SymbolToInt()
			}
		}
	}

	return sum
}

func Boards(paragraphs [][]string, calls []int) []*Board {
	boards := []*Board{}

	for _, paragraph := range paragraphs {
		boards = append(boards, NewBoard(paragraph))
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
