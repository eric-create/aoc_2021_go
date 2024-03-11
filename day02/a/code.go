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
	finish := Move(&vectors.Vector{0, 0}, movements)
	fmt.Println(finish.X * finish.Y)
}

func Move(v *vectors.Vector, movements []*vectors.Vector) *vectors.Vector {
	for _, movement := range movements {
		*v = *v.Add(movement)
	}
	return v
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
