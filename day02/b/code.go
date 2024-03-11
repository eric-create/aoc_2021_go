package main

import (
	"eric-create/aoc_2021/utils"
	"eric-create/aoc_2021/vectors"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	movements := Vectors(*lines)
	Move(&vectors.Vector{X: 0, Y: 0}, movements)
}

func Move(v *vectors.Vector, movements []*vectors.Vector) {
	x, y, aim := 0, 0, 0

	for _, movement := range movements {

		// forward
		if movement.X != 0 {
			x += movement.X
			y += movement.X * aim

		} else { // up, down
			aim += movement.Y
		}
	}

	fmt.Println(x * y)
}

func Vectors(lines []string) []*vectors.Vector {
	vectors := []*vectors.Vector{}

	for _, line := range lines {
		vectors = append(vectors, Vector(line))
	}

	return vectors
}

func Vector(line string) *vectors.Vector {
	if strings.HasPrefix(line, "forward") {
		magnitude, _ := utils.ExtractInt(line)
		return vectors.NewVector(magnitude, 0)

	} else if strings.HasPrefix(line, "up") {
		magnitude, _ := utils.ExtractInt(line)
		return vectors.NewVector(0, -magnitude)

	} else if strings.HasPrefix(line, "down") {
		magnitude, _ := utils.ExtractInt(line)
		return vectors.NewVector(0, magnitude)
	}

	return nil
}
