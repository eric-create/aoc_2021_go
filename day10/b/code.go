package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
	"slices"
	"sort"
)

const (
	CORRECT    = "correct"
	CORRUPTED  = "corrupted"
	INCOMPLETE = "incomplete"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")

	lineScores := []int{}

	for _, line := range *lines {
		stack := GetStack(line)
		if len(stack) > 0 {
			completion := Completion(stack)
			lineScore := LineScore(completion)
			lineScores = append(lineScores, lineScore)
		}

	}

	sort.Ints(lineScores)
	fmt.Println(lineScores)
	fmt.Println(lineScores[len(lineScores)/2])
}

func LineScore(completionStack []rune) int {
	lineScore := 0

	for _, r := range completionStack {
		lineScore = (lineScore * 5) + Points(r)
	}

	return lineScore
}

func Points(r rune) int {
	switch r {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	default:
		return -1
	}
}

func Completion(stack []rune) []rune {
	completion := []rune{}

	for i := len(stack) - 1; i >= 0; i-- {
		completion = append(completion, GetComplementary(stack[i]))
	}

	return completion
}

func GetComplementary(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		return ' '
	}
}

func GetStack(line string) []rune {
	stack := []rune{}

	for _, r := range line {
		if IsOpening(r) {
			// Character is an opening: Top stack
			stack = append(stack, r)

		} else {
			// Character is a closing: Pop stack

			// Character is corrupted
			if !Close(r, Pop(&stack)) {
				return []rune{}
			}
		}
	}

	return stack
}

func IsOpening(r rune) bool {
	return slices.Contains([]rune{'(', '[', '{', '<'}, r)
}

func Close(current, popped rune) bool {
	if current == ')' && popped == '(' {
		return true
	} else if current == ']' && popped == '[' {
		return true
	} else if current == '}' && popped == '{' {
		return true
	} else if current == '>' && popped == '<' {
		return true
	}

	return false
}

func Pop[T any](s *[]T) T {
	iMax := len(*s) - 1
	popped := (*s)[iMax]
	*s = (*s)[:iMax]
	return popped
}
