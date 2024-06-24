package main

import (
	"eric-create/aoc_2021/nodes"
	"eric-create/aoc_2021/utils"
	"eric-create/aoc_2021/vectors"
	"fmt"
	"slices"
)

func main() {
	textLines, _ := utils.ReadLines("input.txt")
	lines, xMax, yMax := GetLines((*textLines)[:len(*textLines)-1])
	// lines = Filter(lines)

	field := nodes.EmptyField(xMax, yMax)
	for _, line := range lines {
		line.Walk(&field)
	}
	fmt.Println(CountPoints(&field, xMax, yMax))
	// nodes.Print(field)

}

func GetLines(textLines []string) ([]*Line, int, int) {
	lines := []*Line{}

	var line *Line = nil
	xMax := 0
	yMax := 0

	for _, textLine := range textLines {
		line, xMax, yMax = GetLine(textLine, xMax, yMax)
		lines = append(lines, line)
	}

	return lines, xMax + 1, yMax + 1
}

func GetLine(textLine string, xMax, yMax int) (*Line, int, int) {
	numbers := utils.ExtractInts(textLine)

	xStart := numbers[0]
	yStart := numbers[1]
	xEnd := numbers[2]
	yEnd := numbers[3]
	xMax, yMax = GetMax(xStart, yStart, xEnd, yEnd, xMax, yMax)

	start := vectors.NewVector(xStart, yStart)
	end := vectors.NewVector(xEnd, yEnd)
	line := NewLine(start, end)

	return line, xMax, yMax
}

func GetMax(xStart, yStart, xEnd, yEnd, xMax, yMax int) (int, int) {
	if xStart > xMax {
		xMax = xStart
	}
	if xEnd > xMax {
		xMax = xEnd
	}

	if yStart > yMax {
		yMax = yStart
	}
	if yEnd > yMax {
		yMax = yEnd
	}

	return xMax, yMax
}

type Line struct {
	Start      *vectors.Vector
	End        *vectors.Vector
	Difference *vectors.Vector
	Direction  *vectors.Vector
	Positions  []*nodes.Node
}

func NewLine(start, end *vectors.Vector) *Line {
	difference := vectors.Difference(start, end)
	direction := vectors.Normalize(difference)
	return &Line{start, end, difference, direction, []*nodes.Node{}}
}

func (l *Line) Walk(field *[][]*nodes.Node) {
	cursor := l.Start
	finish := false

	for {

		// Create node if never been there.
		if (*field)[cursor.Y][cursor.X] == nil {
			(*field)[cursor.Y][cursor.X] = nodes.NewNode(cursor.X, cursor.Y, "o")
		}

		// Always tag.
		(*field)[cursor.Y][cursor.X].Tags = append((*field)[cursor.Y][cursor.X].Tags, "1")

		// Also tag the end position.
		if finish {
			break
		}
		cursor = cursor.Add(l.Direction)
		if *cursor == *l.End {
			finish = true
		}
	}
}

func Filter(lines []*Line) []*Line {
	filtered := []*Line{}

	for _, line := range lines {
		if slices.Contains(vectors.ManhattanDirections(), *line.Direction) {
			filtered = append(filtered, line)
		}
	}

	return filtered
}

func CountPoints(field *[][]*nodes.Node, xMax, yMax int) int {
	points := 0

	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if (*field)[y][x] != nil {

				node := (*field)[y][x]
				if len(node.Tags) > 1 {
					node.Symbol = "x"
					points++
				}
			}
		}
	}
	return points
}
