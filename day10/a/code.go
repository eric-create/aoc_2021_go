package main

import (
	"eric-create/aoc_2021/utils"
	"fmt"
	"slices"
)

const (
	CORRECT    = "correct"
	CORRUPTED  = "corrupted"
	INCOMPLETE = "incomplete"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")

	score := 0

	for _, line := range *lines {
		_, points := CheckLine(line)
		score += points
	}

	fmt.Println(score)
}

func CheckLine(line string) (string, int) {
	stack := []rune{}

	for _, r := range line {
		if IsOpening(r) {
			// Character is an opening: Top stack
			stack = append(stack, r)

		} else {
			// Character is a closing: Pop stack

			// Character is excess (equals incomplete?!)
			if len(stack) == 0 {
				return INCOMPLETE, 0
			}

			popped := Pop(&stack)

			// Character is corrupted
			if !Close(r, popped) {
				return CORRUPTED, Points(r)
			}

		}
	}

	if len(stack) > 0 {
		return INCOMPLETE, 0
	}

	return CORRECT, 0
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

func Points(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return -1
	}
}
